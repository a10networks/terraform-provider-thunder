package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceDnsARecordDnsARecordIpv6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats`: Statistics for the object dns-a-record-ipv6\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceDnsARecordDnsARecordIpv6StatsRead,

		Schema: map[string]*schema.Schema{
			"dns_a_record_ipv6": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address",
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

func resourceGslbZoneServiceDnsARecordDnsARecordIpv6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceDnsARecordDnsARecordIpv6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats := setObjectGslbZoneServiceDnsARecordDnsARecordIpv6StatsStats(res)
		d.Set("stats", GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbZoneServiceDnsARecordDnsARecordIpv6StatsStats(ret edpt.DataGslbZoneServiceDnsARecordDnsARecordIpv6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtGslbZoneServiceDnsARecordDnsARecordIpv6Stats.Stats.Hits,
		},
	}
}

func getObjectGslbZoneServiceDnsARecordDnsARecordIpv6StatsStats(d []interface{}) edpt.GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceDnsARecordDnsARecordIpv6Stats(d *schema.ResourceData) edpt.GslbZoneServiceDnsARecordDnsARecordIpv6Stats {
	var ret edpt.GslbZoneServiceDnsARecordDnsARecordIpv6Stats

	ret.DnsARecordIpv6 = d.Get("dns_a_record_ipv6").(string)

	ret.Stats = getObjectGslbZoneServiceDnsARecordDnsARecordIpv6StatsStats(d.Get("stats").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
