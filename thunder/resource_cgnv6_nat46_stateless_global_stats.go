package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat46_stateless_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat46StatelessGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outbound_ipv4_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv4 packets received",
						},
						"outbound_ipv4_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv4 packets dropped",
						},
						"outbound_ipv4_fragment_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv4 fragment packets received",
						},
						"outbound_ipv6_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 destination unreachable",
						},
						"outbound_ipv6_fragmented": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 packets fragmented",
						},
						"inbound_ipv6_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv6 packets received",
						},
						"inbound_ipv6_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv6 packets dropped",
						},
						"inbound_ipv6_fragment_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv6 fragment packets received",
						},
						"inbound_ipv4_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 destination unreachable",
						},
						"inbound_ipv4_fragmented": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 packets fragmented",
						},
						"packet_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Packet too big",
						},
						"fragment_error": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment processing errors",
						},
						"icmpv6_to_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMP",
						},
						"icmpv6_to_icmp_error": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMP errors",
						},
						"icmp_to_icmpv6": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMPv6",
						},
						"icmp_to_icmpv6_error": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMPv6 errors",
						},
						"ha_standby": {
							Type: schema.TypeInt, Optional: true, Description: "HA is standby",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other errors",
						},
						"conn_count": {
							Type: schema.TypeInt, Optional: true, Description: "conn count",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nat46StatelessGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat46StatelessGlobalStatsStats := setObjectCgnv6Nat46StatelessGlobalStatsStats(res)
		d.Set("stats", Cgnv6Nat46StatelessGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat46StatelessGlobalStatsStats(ret edpt.DataCgnv6Nat46StatelessGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"outbound_ipv4_received":          ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Outbound_ipv4_received,
			"outbound_ipv4_drop":              ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Outbound_ipv4_drop,
			"outbound_ipv4_fragment_received": ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Outbound_ipv4_fragment_received,
			"outbound_ipv6_unreachable":       ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Outbound_ipv6_unreachable,
			"outbound_ipv6_fragmented":        ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Outbound_ipv6_fragmented,
			"inbound_ipv6_received":           ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Inbound_ipv6_received,
			"inbound_ipv6_drop":               ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Inbound_ipv6_drop,
			"inbound_ipv6_fragment_received":  ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Inbound_ipv6_fragment_received,
			"inbound_ipv4_unreachable":        ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Inbound_ipv4_unreachable,
			"inbound_ipv4_fragmented":         ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Inbound_ipv4_fragmented,
			"packet_too_big":                  ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Packet_too_big,
			"fragment_error":                  ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Fragment_error,
			"icmpv6_to_icmp":                  ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Icmpv6_to_icmp,
			"icmpv6_to_icmp_error":            ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Icmpv6_to_icmp_error,
			"icmp_to_icmpv6":                  ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Icmp_to_icmpv6,
			"icmp_to_icmpv6_error":            ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Icmp_to_icmpv6_error,
			"ha_standby":                      ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Ha_standby,
			"other_error":                     ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Other_error,
			"conn_count":                      ret.DtCgnv6Nat46StatelessGlobalStats.Stats.Conn_count,
		},
	}
}

func getObjectCgnv6Nat46StatelessGlobalStatsStats(d []interface{}) edpt.Cgnv6Nat46StatelessGlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Nat46StatelessGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outbound_ipv4_received = in["outbound_ipv4_received"].(int)
		ret.Outbound_ipv4_drop = in["outbound_ipv4_drop"].(int)
		ret.Outbound_ipv4_fragment_received = in["outbound_ipv4_fragment_received"].(int)
		ret.Outbound_ipv6_unreachable = in["outbound_ipv6_unreachable"].(int)
		ret.Outbound_ipv6_fragmented = in["outbound_ipv6_fragmented"].(int)
		ret.Inbound_ipv6_received = in["inbound_ipv6_received"].(int)
		ret.Inbound_ipv6_drop = in["inbound_ipv6_drop"].(int)
		ret.Inbound_ipv6_fragment_received = in["inbound_ipv6_fragment_received"].(int)
		ret.Inbound_ipv4_unreachable = in["inbound_ipv4_unreachable"].(int)
		ret.Inbound_ipv4_fragmented = in["inbound_ipv4_fragmented"].(int)
		ret.Packet_too_big = in["packet_too_big"].(int)
		ret.Fragment_error = in["fragment_error"].(int)
		ret.Icmpv6_to_icmp = in["icmpv6_to_icmp"].(int)
		ret.Icmpv6_to_icmp_error = in["icmpv6_to_icmp_error"].(int)
		ret.Icmp_to_icmpv6 = in["icmp_to_icmpv6"].(int)
		ret.Icmp_to_icmpv6_error = in["icmp_to_icmpv6_error"].(int)
		ret.Ha_standby = in["ha_standby"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.Conn_count = in["conn_count"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Nat46StatelessGlobalStats(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessGlobalStats {
	var ret edpt.Cgnv6Nat46StatelessGlobalStats

	ret.Stats = getObjectCgnv6Nat46StatelessGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
