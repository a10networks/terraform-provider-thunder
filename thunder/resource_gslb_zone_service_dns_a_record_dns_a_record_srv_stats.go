package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsARecordDnsARecordSrvStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_a_record_dns_a_record_srv_stats`: Statistics for the object dns-a-record-srv\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsARecordDnsARecordSrvStatsRead,

		Schema: map[string]*schema.Schema{
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
			"svrname": {
				Type: schema.TypeString, Required: true, Description: "Specify name",
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

func resourceGslbZoneServiceDnsARecordDnsARecordSrvStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordSrvStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrvStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsARecordDnsARecordSrvStatsStats := setObjectGslbZoneServiceDnsARecordDnsARecordSrvStatsStats(res)
		d.Set("stats", GslbZoneServiceDnsARecordDnsARecordSrvStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsARecordDnsARecordSrvStatsStats(ret edpt.DataGslbZoneServiceDnsARecordDnsARecordSrvStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsARecordDnsARecordSrvStats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsARecordDnsARecordSrvStatsStats(d []interface{}) edpt.GslbZoneServiceDnsARecordDnsARecordSrvStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordSrvStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsARecordDnsARecordSrvStats(d *schema.ResourceData) edpt.GslbZoneServiceDnsARecordDnsARecordSrvStats {
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordSrvStats

	ret.Stats = getObjectGslbZoneServiceDnsARecordDnsARecordSrvStatsStats(d.Get("stats").([]interface{}))

	ret.Svrname = d.Get("svrname").(string)

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
