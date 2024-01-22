package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsCaaRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_caa_record_stats`: Statistics for the object dns-caa-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsCaaRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"critical_flag": {
				Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
			},
			"property_tag": {
				Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
			},
			"rdata": {
				Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the CAA has been used",
						},
					},
				},
			},
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceDnsCaaRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsCaaRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsCaaRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsCaaRecordStatsStats := setObjectGslbZoneServiceDnsCaaRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsCaaRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsCaaRecordStatsStats(ret edpt.DataGslbZoneServiceDnsCaaRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsCaaRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsCaaRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsCaaRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsCaaRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsCaaRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsCaaRecordStats {
	var ret edpt.GslbZoneServiceDnsCaaRecordStats

	ret.CriticalFlag = d.Get("critical_flag").(int)

	ret.PropertyTag = d.Get("property_tag").(string)

	ret.Rdata = d.Get("rdata").(string)

	ret.Stats = getObjectGslbZoneServiceDnsCaaRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServicePort = d.Get("service_port").(string)

	ret.ServiceName = d.Get("service_name").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
