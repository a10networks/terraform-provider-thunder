package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsNaptrRecordStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_naptr_record_stats`: Statistics for the object dns-naptr-record\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsNaptrRecordStatsRead,

		Schema: map[string]*schema.Schema{
			"flag": {
				Type: schema.TypeString, Required: true, Description: "Specify the flag (e.g., a, s). Default is empty flag",
			},
			"naptr_target": {
				Type: schema.TypeString, Required: true, Description: "Specify the replacement or regular expression",
			},
			"service_proto": {
				Type: schema.TypeString, Required: true, Description: "Specify Service and Protocol",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"naptr_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the NAPTR has been used",
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

func resourceGslbZoneServiceDnsNaptrRecordStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsNaptrRecordStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsNaptrRecordStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsNaptrRecordStatsStats := setObjectGslbZoneServiceDnsNaptrRecordStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsNaptrRecordStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsNaptrRecordStatsStats(ret edpt.DataGslbZoneServiceDnsNaptrRecordStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"naptr_hits": ret.DtGslbZoneServiceDnsNaptrRecordStats.Stats.NaptrHits,
		},
	}
}

func getObjectGslbZoneServiceDnsNaptrRecordStatsStats(d []interface{}) edpt.GslbZoneServiceDnsNaptrRecordStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsNaptrRecordStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NaptrHits = in["naptr_hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsNaptrRecordStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsNaptrRecordStats {
	var ret edpt.GslbZoneServiceDnsNaptrRecordStats

	ret.Flag = d.Get("flag").(string)

	ret.NaptrTarget = d.Get("naptr_target").(string)

	ret.ServiceProto = d.Get("service_proto").(string)

	ret.Stats = getObjectGslbZoneServiceDnsNaptrRecordStatsStats(d.Get("stats").([]interface{}))

	ret.ServicePort = d.Get("service_port").(string)

	ret.ServiceName = d.Get("service_name").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
