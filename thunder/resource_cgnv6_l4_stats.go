package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6L4Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_l4_stats`: Statistics for the object l4\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6L4StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Forward Route for Session",
						},
						"no_rev_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Reverse Route for Session",
						},
						"out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Session Memory",
						},
						"tcp_rst_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Sent",
						},
						"ipip_icmp_reply_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IPIP ICMP Echo Reply Sent",
						},
						"icmp_filtered_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Administratively Filtered Sent",
						},
						"icmp_host_unreachable_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Host Unreachable Sent",
						},
						"icmp_reply_no_session_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Reply No Session Drop",
						},
						"ipip_truncated": {
							Type: schema.TypeInt, Optional: true, Description: "IPIP Truncated Packet",
						},
						"ip_src_invalid_unicast": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Source Not Valid Unicast",
						},
						"ip_dst_invalid_unicast": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Destination Not Valid Unicast",
						},
						"ipv6_src_invalid_unicast": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Source Not Valid Unicast",
						},
						"ipv6_dst_invalid_unicast": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Destination Not Valid Unicast",
						},
						"rate_drop_reset_unkn": {
							Type: schema.TypeInt, Optional: true, Description: "Rate Drop reset",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6L4StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6L4StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6L4Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6L4StatsStats := setObjectCgnv6L4StatsStats(res)
		d.Set("stats", Cgnv6L4StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6L4StatsStats(ret edpt.DataCgnv6L4Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"no_fwd_route":               ret.DtCgnv6L4Stats.Stats.NoFwdRoute,
			"no_rev_route":               ret.DtCgnv6L4Stats.Stats.NoRevRoute,
			"out_of_session_memory":      ret.DtCgnv6L4Stats.Stats.OutOfSessionMemory,
			"tcp_rst_sent":               ret.DtCgnv6L4Stats.Stats.TcpRstSent,
			"ipip_icmp_reply_sent":       ret.DtCgnv6L4Stats.Stats.IpipIcmpReplySent,
			"icmp_filtered_sent":         ret.DtCgnv6L4Stats.Stats.IcmpFilteredSent,
			"icmp_host_unreachable_sent": ret.DtCgnv6L4Stats.Stats.IcmpHostUnreachableSent,
			"icmp_reply_no_session_drop": ret.DtCgnv6L4Stats.Stats.IcmpReplyNoSessionDrop,
			"ipip_truncated":             ret.DtCgnv6L4Stats.Stats.IpipTruncated,
			"ip_src_invalid_unicast":     ret.DtCgnv6L4Stats.Stats.IpSrcInvalidUnicast,
			"ip_dst_invalid_unicast":     ret.DtCgnv6L4Stats.Stats.IpDstInvalidUnicast,
			"ipv6_src_invalid_unicast":   ret.DtCgnv6L4Stats.Stats.Ipv6SrcInvalidUnicast,
			"ipv6_dst_invalid_unicast":   ret.DtCgnv6L4Stats.Stats.Ipv6DstInvalidUnicast,
			"rate_drop_reset_unkn":       ret.DtCgnv6L4Stats.Stats.Rate_drop_reset_unkn,
		},
	}
}

func getObjectCgnv6L4StatsStats(d []interface{}) edpt.Cgnv6L4StatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6L4StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NoFwdRoute = in["no_fwd_route"].(int)
		ret.NoRevRoute = in["no_rev_route"].(int)
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.TcpRstSent = in["tcp_rst_sent"].(int)
		ret.IpipIcmpReplySent = in["ipip_icmp_reply_sent"].(int)
		ret.IcmpFilteredSent = in["icmp_filtered_sent"].(int)
		ret.IcmpHostUnreachableSent = in["icmp_host_unreachable_sent"].(int)
		ret.IcmpReplyNoSessionDrop = in["icmp_reply_no_session_drop"].(int)
		ret.IpipTruncated = in["ipip_truncated"].(int)
		ret.IpSrcInvalidUnicast = in["ip_src_invalid_unicast"].(int)
		ret.IpDstInvalidUnicast = in["ip_dst_invalid_unicast"].(int)
		ret.Ipv6SrcInvalidUnicast = in["ipv6_src_invalid_unicast"].(int)
		ret.Ipv6DstInvalidUnicast = in["ipv6_dst_invalid_unicast"].(int)
		ret.Rate_drop_reset_unkn = in["rate_drop_reset_unkn"].(int)
	}
	return ret
}

func dataToEndpointCgnv6L4Stats(d *schema.ResourceData) edpt.Cgnv6L4Stats {
	var ret edpt.Cgnv6L4Stats

	ret.Stats = getObjectCgnv6L4StatsStats(d.Get("stats").([]interface{}))
	return ret
}
