package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6IcmpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_icmp_stats`: Statistics for the object icmp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6IcmpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_unknown_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Unknown Type",
						},
						"icmp_no_port_info": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Port Info Not Included",
						},
						"icmp_no_session_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP No Matching Session Drop",
						},
						"icmpv6_unknown_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 Unknown Type",
						},
						"icmpv6_no_port_info": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 Port Info Not Included",
						},
						"icmpv6_no_session_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 No Matching Session Drop",
						},
						"icmp_to_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMP Conversion",
						},
						"icmp_to_icmpv6": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMPv6 Conversion",
						},
						"icmpv6_to_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMP Conversion",
						},
						"icmpv6_to_icmpv6": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMPv6 Conversion",
						},
						"icmp_bad_type": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Embedded ICMP Type",
						},
						"icmpv6_bad_type": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Embedded ICMPv6 Type",
						},
						"known_drop64": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Forward Known ICMPv6 Drop",
						},
						"unknown_drop64": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Forward Unknown ICMPv6 Drop",
						},
						"midpoint_hop64": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Forward Unknown Source Drop",
						},
						"known_drop46": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Reverse Known ICMP Drop",
						},
						"unknown_drop46": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Reverse Known ICMPv6 Drop",
						},
						"no_prefix_for_ipv446": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Reverse No Prefix Match for IPv4",
						},
						"icmp_to_icmp_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMP Conversion Error",
						},
						"icmp_to_icmpv6_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP to ICMPv6 Conversion Error",
						},
						"icmpv6_to_icmp_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMP Conversion Error",
						},
						"icmpv6_to_icmpv6_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 to ICMPv6 Conversion Error",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6IcmpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6IcmpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6IcmpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6IcmpStatsStats := setObjectCgnv6IcmpStatsStats(res)
		d.Set("stats", Cgnv6IcmpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6IcmpStatsStats(ret edpt.DataCgnv6IcmpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"icmp_unknown_type":      ret.DtCgnv6IcmpStats.Stats.IcmpUnknownType,
			"icmp_no_port_info":      ret.DtCgnv6IcmpStats.Stats.IcmpNoPortInfo,
			"icmp_no_session_drop":   ret.DtCgnv6IcmpStats.Stats.IcmpNoSessionDrop,
			"icmpv6_unknown_type":    ret.DtCgnv6IcmpStats.Stats.Icmpv6UnknownType,
			"icmpv6_no_port_info":    ret.DtCgnv6IcmpStats.Stats.Icmpv6NoPortInfo,
			"icmpv6_no_session_drop": ret.DtCgnv6IcmpStats.Stats.Icmpv6NoSessionDrop,
			"icmp_to_icmp":           ret.DtCgnv6IcmpStats.Stats.IcmpToIcmp,
			"icmp_to_icmpv6":         ret.DtCgnv6IcmpStats.Stats.IcmpToIcmpv6,
			"icmpv6_to_icmp":         ret.DtCgnv6IcmpStats.Stats.Icmpv6ToIcmp,
			"icmpv6_to_icmpv6":       ret.DtCgnv6IcmpStats.Stats.Icmpv6ToIcmpv6,
			"icmp_bad_type":          ret.DtCgnv6IcmpStats.Stats.IcmpBadType,
			"icmpv6_bad_type":        ret.DtCgnv6IcmpStats.Stats.Icmpv6BadType,
			"known_drop64":           ret.DtCgnv6IcmpStats.Stats.KnownDrop64,
			"unknown_drop64":         ret.DtCgnv6IcmpStats.Stats.UnknownDrop64,
			"midpoint_hop64":         ret.DtCgnv6IcmpStats.Stats.MidpointHop64,
			"known_drop46":           ret.DtCgnv6IcmpStats.Stats.KnownDrop46,
			"unknown_drop46":         ret.DtCgnv6IcmpStats.Stats.UnknownDrop46,
			"no_prefix_for_ipv446":   ret.DtCgnv6IcmpStats.Stats.NoPrefixForIpv446,
			"icmp_to_icmp_err":       ret.DtCgnv6IcmpStats.Stats.IcmpToIcmpErr,
			"icmp_to_icmpv6_err":     ret.DtCgnv6IcmpStats.Stats.IcmpToIcmpv6Err,
			"icmpv6_to_icmp_err":     ret.DtCgnv6IcmpStats.Stats.Icmpv6ToIcmpErr,
			"icmpv6_to_icmpv6_err":   ret.DtCgnv6IcmpStats.Stats.Icmpv6ToIcmpv6Err,
		},
	}
}

func getObjectCgnv6IcmpStatsStats(d []interface{}) edpt.Cgnv6IcmpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6IcmpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpUnknownType = in["icmp_unknown_type"].(int)
		ret.IcmpNoPortInfo = in["icmp_no_port_info"].(int)
		ret.IcmpNoSessionDrop = in["icmp_no_session_drop"].(int)
		ret.Icmpv6UnknownType = in["icmpv6_unknown_type"].(int)
		ret.Icmpv6NoPortInfo = in["icmpv6_no_port_info"].(int)
		ret.Icmpv6NoSessionDrop = in["icmpv6_no_session_drop"].(int)
		ret.IcmpToIcmp = in["icmp_to_icmp"].(int)
		ret.IcmpToIcmpv6 = in["icmp_to_icmpv6"].(int)
		ret.Icmpv6ToIcmp = in["icmpv6_to_icmp"].(int)
		ret.Icmpv6ToIcmpv6 = in["icmpv6_to_icmpv6"].(int)
		ret.IcmpBadType = in["icmp_bad_type"].(int)
		ret.Icmpv6BadType = in["icmpv6_bad_type"].(int)
		ret.KnownDrop64 = in["known_drop64"].(int)
		ret.UnknownDrop64 = in["unknown_drop64"].(int)
		ret.MidpointHop64 = in["midpoint_hop64"].(int)
		ret.KnownDrop46 = in["known_drop46"].(int)
		ret.UnknownDrop46 = in["unknown_drop46"].(int)
		ret.NoPrefixForIpv446 = in["no_prefix_for_ipv446"].(int)
		ret.IcmpToIcmpErr = in["icmp_to_icmp_err"].(int)
		ret.IcmpToIcmpv6Err = in["icmp_to_icmpv6_err"].(int)
		ret.Icmpv6ToIcmpErr = in["icmpv6_to_icmp_err"].(int)
		ret.Icmpv6ToIcmpv6Err = in["icmpv6_to_icmpv6_err"].(int)
	}
	return ret
}

func dataToEndpointCgnv6IcmpStats(d *schema.ResourceData) edpt.Cgnv6IcmpStats {
	var ret edpt.Cgnv6IcmpStats

	ret.Stats = getObjectCgnv6IcmpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
