package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsCnameRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_cname_record_stats`: Statistics for the object dns-cname-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsCnameRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"alias_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the alias name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cname_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the CNAME has been used",
						},
					},
				},
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceDnsCnameRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCnameRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCnameRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsCnameRecordStatsStats := setObjectGslbZoneServiceDnsCnameRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsCnameRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsCnameRecordStatsStats(ret edpt.DataGslbZoneServiceDnsCnameRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cname_hits": ret.DtGslbZoneServiceDnsCnameRecordStats.Stats.CnameHits,
		},
	}
}

func getObjectGslbZoneServiceDnsCnameRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsCnameRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsCnameRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CnameHits = in["cname_hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsCnameRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsCnameRecordStats {
	var ret edpt.GslbZoneServiceDnsCnameRecordStats

	ret.AliasName = d.Get("alias_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsCnameRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
