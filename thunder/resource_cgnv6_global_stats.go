package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6GlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6GlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_total_ports_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP ports allocated",
						},
						"udp_total_ports_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP ports allocated",
						},
						"icmp_total_ports_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total ICMP ports allocated",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6GlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6GlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6GlobalStatsStats := setObjectCgnv6GlobalStatsStats(res)
		d.Set("stats", Cgnv6GlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6GlobalStatsStats(ret edpt.DataCgnv6GlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_total_ports_allocated":  ret.DtCgnv6GlobalStats.Stats.TcpTotalPortsAllocated,
			"udp_total_ports_allocated":  ret.DtCgnv6GlobalStats.Stats.UdpTotalPortsAllocated,
			"icmp_total_ports_allocated": ret.DtCgnv6GlobalStats.Stats.IcmpTotalPortsAllocated,
		},
	}
}

func getObjectCgnv6GlobalStatsStats(d []interface{}) edpt.Cgnv6GlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6GlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpTotalPortsAllocated = in["tcp_total_ports_allocated"].(int)
		ret.UdpTotalPortsAllocated = in["udp_total_ports_allocated"].(int)
		ret.IcmpTotalPortsAllocated = in["icmp_total_ports_allocated"].(int)
	}
	return ret
}

func dataToEndpointCgnv6GlobalStats(d *schema.ResourceData) edpt.Cgnv6GlobalStats {
	var ret edpt.Cgnv6GlobalStats

	ret.Stats = getObjectCgnv6GlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
