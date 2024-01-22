package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsNsRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_ns_record_stats`: Statistics for the object dns-ns-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsNsRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"ns_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
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

func resourceGslbZoneServiceDnsNsRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNsRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNsRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsNsRecordStatsStats := setObjectGslbZoneServiceDnsNsRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsNsRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsNsRecordStatsStats(ret edpt.DataGslbZoneServiceDnsNsRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsNsRecordStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsNsRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsNsRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsNsRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsNsRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsNsRecordStats {
	var ret edpt.GslbZoneServiceDnsNsRecordStats

	ret.NsName = d.Get("ns_name").(string)

	ret.Stats = getObjectGslbZoneServiceDnsNsRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
