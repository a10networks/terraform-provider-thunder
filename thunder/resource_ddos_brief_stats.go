package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosBriefStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_brief_stats`: Statistics for the object brief\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosBriefStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Received",
						},
						"ip_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Sent",
						},
						"ipv6_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Received",
						},
						"ipv6_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Sent",
						},
						"out_no_route": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4/v6 Out No Route",
						},
						"not_for_ddos": {
							Type: schema.TypeInt, Optional: true, Description: "Not For DDOS",
						},
						"instateless": {
							Type: schema.TypeInt, Optional: true, Description: "Stateless Packets Received",
						},
						"intcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Received",
						},
						"inudp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Received",
						},
						"inicmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Received",
						},
						"inother": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Received",
						},
						"v4_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Session Created",
						},
						"v6_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Session Created",
						},
						"tcp_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Created",
						},
						"udp_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Created",
						},
						"sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "Session Aged Out",
						},
						"tcp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Dropped",
						},
						"tcp_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst Packets Dropped",
						},
						"tcp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Packets Dropped",
						},
						"tcp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SrcDst Packets Dropped",
						},
						"udp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Dropped",
						},
						"udp_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst Packets Dropped",
						},
						"udp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Packets Dropped",
						},
						"udp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcDst Packets Dropped",
						},
						"icmp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Dropped",
						},
						"icmp_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Packets Dropped",
						},
						"icmp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Packets Dropped",
						},
						"icmp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP SrcDst Packets Dropped",
						},
						"other_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Dropped",
						},
						"other_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst Packets Dropped",
						},
						"other_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Packets Dropped",
						},
						"other_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER SrcDst Packets Dropped",
						},
						"frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
						},
						"frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
						},
						"dst_port_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Undefined Dropped",
						},
						"dst_port_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Exceed Dropped",
						},
						"dst_ipproto_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Blacklist Packets Dropped",
						},
						"dst_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Blacklist Packets Dropped",
						},
						"dst_sport_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Blacklist Packets Dropped",
						},
						"dst_sport_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Exceed Dropped",
						},
						"dst_ipproto_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Total Packets Received",
						},
						"dst_ipproto_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Total Packets Dropped",
						},
						"dst_ipproto_exceed_drop_any": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Exceed Dropped",
						},
						"src_ip_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Src IP Bypass",
						},
						"dst_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Received",
						},
						"dst_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Received",
						},
						"dst_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Received",
						},
						"dst_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Received",
						},
						"dst_ip_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP Bypass",
						},
						"dst_blackhole_inject": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Injected",
						},
						"dst_blackhole_withdraw": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Withdrawn",
						},
						"tcp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Received",
						},
						"tcp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Dropped",
						},
						"udp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Received",
						},
						"udp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Dropped",
						},
						"icmp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Received",
						},
						"icmp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Dropped",
						},
						"other_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Received",
						},
						"other_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Dropped",
						},
						"udp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Exceeded",
						},
						"tcp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Exceeded",
						},
						"icmp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Exceeded",
						},
						"other_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Exceeded",
						},
						"tcp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Blacklist Packets Dropped",
						},
						"udp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Blacklist Packets Dropped",
						},
						"icmp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Blacklisted Packets Dropped",
						},
						"other_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Blacklisted Packets Dropped",
						},
						"glid_action_encap_send_immed": {
							Type: schema.TypeInt, Optional: true, Description: "Glid Action Tunnel-encap",
						},
						"glid_action_encap_send_delay": {
							Type: schema.TypeInt, Optional: true, Description: "Glid Action Tunnel-encap with Scrub",
						},
						"dst_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Default Hardware Packets Dropped",
						},
						"dst_hw_drop_rule_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Default Hardware Drop Rules Inserted",
						},
						"dst_hw_drop_rule_removed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Default Hardware Drop Rules Removed",
						},
						"src_hw_drop_rule_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Src Default Hardware Drop Rules Inserted",
						},
						"src_hw_drop_rule_removed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Default Hardware Drop Rules Removed",
						},
					},
				},
			},
		},
	}
}

func resourceDdosBriefStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosBriefStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosBriefStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosBriefStatsStats := setObjectDdosBriefStatsStats(res)
		d.Set("stats", DdosBriefStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosBriefStatsStats(ret edpt.DataDdosBriefStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_rcv":                       ret.DtDdosBriefStats.Stats.Ip_rcv,
			"ip_sent":                      ret.DtDdosBriefStats.Stats.Ip_sent,
			"ipv6_rcv":                     ret.DtDdosBriefStats.Stats.Ipv6_rcv,
			"ipv6_sent":                    ret.DtDdosBriefStats.Stats.Ipv6_sent,
			"out_no_route":                 ret.DtDdosBriefStats.Stats.Out_no_route,
			"not_for_ddos":                 ret.DtDdosBriefStats.Stats.Not_for_ddos,
			"instateless":                  ret.DtDdosBriefStats.Stats.Instateless,
			"intcp":                        ret.DtDdosBriefStats.Stats.Intcp,
			"inudp":                        ret.DtDdosBriefStats.Stats.Inudp,
			"inicmp":                       ret.DtDdosBriefStats.Stats.Inicmp,
			"inother":                      ret.DtDdosBriefStats.Stats.Inother,
			"v4_sess_create":               ret.DtDdosBriefStats.Stats.V4_sess_create,
			"v6_sess_create":               ret.DtDdosBriefStats.Stats.V6_sess_create,
			"tcp_sess_create":              ret.DtDdosBriefStats.Stats.Tcp_sess_create,
			"udp_sess_create":              ret.DtDdosBriefStats.Stats.Udp_sess_create,
			"sess_aged_out":                ret.DtDdosBriefStats.Stats.Sess_aged_out,
			"tcp_total_drop":               ret.DtDdosBriefStats.Stats.Tcp_total_drop,
			"tcp_dst_drop":                 ret.DtDdosBriefStats.Stats.Tcp_dst_drop,
			"tcp_src_drop":                 ret.DtDdosBriefStats.Stats.Tcp_src_drop,
			"tcp_src_dst_drop":             ret.DtDdosBriefStats.Stats.Tcp_src_dst_drop,
			"udp_total_drop":               ret.DtDdosBriefStats.Stats.Udp_total_drop,
			"udp_dst_drop":                 ret.DtDdosBriefStats.Stats.Udp_dst_drop,
			"udp_src_drop":                 ret.DtDdosBriefStats.Stats.Udp_src_drop,
			"udp_src_dst_drop":             ret.DtDdosBriefStats.Stats.Udp_src_dst_drop,
			"icmp_total_drop":              ret.DtDdosBriefStats.Stats.Icmp_total_drop,
			"icmp_dst_drop":                ret.DtDdosBriefStats.Stats.Icmp_dst_drop,
			"icmp_src_drop":                ret.DtDdosBriefStats.Stats.Icmp_src_drop,
			"icmp_src_dst_drop":            ret.DtDdosBriefStats.Stats.Icmp_src_dst_drop,
			"other_total_drop":             ret.DtDdosBriefStats.Stats.Other_total_drop,
			"other_dst_drop":               ret.DtDdosBriefStats.Stats.Other_dst_drop,
			"other_src_drop":               ret.DtDdosBriefStats.Stats.Other_src_drop,
			"other_src_dst_drop":           ret.DtDdosBriefStats.Stats.Other_src_dst_drop,
			"frag_rcvd":                    ret.DtDdosBriefStats.Stats.Frag_rcvd,
			"frag_drop":                    ret.DtDdosBriefStats.Stats.Frag_drop,
			"dst_port_undef_drop":          ret.DtDdosBriefStats.Stats.Dst_port_undef_drop,
			"dst_port_exceed_drop_any":     ret.DtDdosBriefStats.Stats.Dst_port_exceed_drop_any,
			"dst_ipproto_bl":               ret.DtDdosBriefStats.Stats.Dst_ipproto_bl,
			"dst_port_bl":                  ret.DtDdosBriefStats.Stats.Dst_port_bl,
			"dst_sport_bl":                 ret.DtDdosBriefStats.Stats.Dst_sport_bl,
			"dst_sport_exceed_drop_any":    ret.DtDdosBriefStats.Stats.Dst_sport_exceed_drop_any,
			"dst_ipproto_rcvd":             ret.DtDdosBriefStats.Stats.Dst_ipproto_rcvd,
			"dst_ipproto_drop":             ret.DtDdosBriefStats.Stats.Dst_ipproto_drop,
			"dst_ipproto_exceed_drop_any":  ret.DtDdosBriefStats.Stats.Dst_ipproto_exceed_drop_any,
			"src_ip_bypass":                ret.DtDdosBriefStats.Stats.Src_ip_bypass,
			"dst_ingress_bytes":            ret.DtDdosBriefStats.Stats.Dst_ingress_bytes,
			"dst_egress_bytes":             ret.DtDdosBriefStats.Stats.Dst_egress_bytes,
			"dst_ingress_packets":          ret.DtDdosBriefStats.Stats.Dst_ingress_packets,
			"dst_egress_packets":           ret.DtDdosBriefStats.Stats.Dst_egress_packets,
			"dst_ip_bypass":                ret.DtDdosBriefStats.Stats.Dst_ip_bypass,
			"dst_blackhole_inject":         ret.DtDdosBriefStats.Stats.Dst_blackhole_inject,
			"dst_blackhole_withdraw":       ret.DtDdosBriefStats.Stats.Dst_blackhole_withdraw,
			"tcp_total_bytes_rcv":          ret.DtDdosBriefStats.Stats.Tcp_total_bytes_rcv,
			"tcp_total_bytes_drop":         ret.DtDdosBriefStats.Stats.Tcp_total_bytes_drop,
			"udp_total_bytes_rcv":          ret.DtDdosBriefStats.Stats.Udp_total_bytes_rcv,
			"udp_total_bytes_drop":         ret.DtDdosBriefStats.Stats.Udp_total_bytes_drop,
			"icmp_total_bytes_rcv":         ret.DtDdosBriefStats.Stats.Icmp_total_bytes_rcv,
			"icmp_total_bytes_drop":        ret.DtDdosBriefStats.Stats.Icmp_total_bytes_drop,
			"other_total_bytes_rcv":        ret.DtDdosBriefStats.Stats.Other_total_bytes_rcv,
			"other_total_bytes_drop":       ret.DtDdosBriefStats.Stats.Other_total_bytes_drop,
			"udp_any_exceed":               ret.DtDdosBriefStats.Stats.Udp_any_exceed,
			"tcp_any_exceed":               ret.DtDdosBriefStats.Stats.Tcp_any_exceed,
			"icmp_any_exceed":              ret.DtDdosBriefStats.Stats.Icmp_any_exceed,
			"other_any_exceed":             ret.DtDdosBriefStats.Stats.Other_any_exceed,
			"tcp_drop_bl":                  ret.DtDdosBriefStats.Stats.Tcp_drop_bl,
			"udp_drop_bl":                  ret.DtDdosBriefStats.Stats.Udp_drop_bl,
			"icmp_drop_bl":                 ret.DtDdosBriefStats.Stats.Icmp_drop_bl,
			"other_drop_bl":                ret.DtDdosBriefStats.Stats.Other_drop_bl,
			"glid_action_encap_send_immed": ret.DtDdosBriefStats.Stats.Glid_action_encap_send_immed,
			"glid_action_encap_send_delay": ret.DtDdosBriefStats.Stats.Glid_action_encap_send_delay,
			"dst_hw_drop":                  ret.DtDdosBriefStats.Stats.Dst_hw_drop,
			"dst_hw_drop_rule_inserted":    ret.DtDdosBriefStats.Stats.Dst_hw_drop_rule_inserted,
			"dst_hw_drop_rule_removed":     ret.DtDdosBriefStats.Stats.Dst_hw_drop_rule_removed,
			"src_hw_drop_rule_inserted":    ret.DtDdosBriefStats.Stats.Src_hw_drop_rule_inserted,
			"src_hw_drop_rule_removed":     ret.DtDdosBriefStats.Stats.Src_hw_drop_rule_removed,
		},
	}
}

func getObjectDdosBriefStatsStats(d []interface{}) edpt.DdosBriefStatsStats {

	count1 := len(d)
	var ret edpt.DdosBriefStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip_rcv = in["ip_rcv"].(int)
		ret.Ip_sent = in["ip_sent"].(int)
		ret.Ipv6_rcv = in["ipv6_rcv"].(int)
		ret.Ipv6_sent = in["ipv6_sent"].(int)
		ret.Out_no_route = in["out_no_route"].(int)
		ret.Not_for_ddos = in["not_for_ddos"].(int)
		ret.Instateless = in["instateless"].(int)
		ret.Intcp = in["intcp"].(int)
		ret.Inudp = in["inudp"].(int)
		ret.Inicmp = in["inicmp"].(int)
		ret.Inother = in["inother"].(int)
		ret.V4_sess_create = in["v4_sess_create"].(int)
		ret.V6_sess_create = in["v6_sess_create"].(int)
		ret.Tcp_sess_create = in["tcp_sess_create"].(int)
		ret.Udp_sess_create = in["udp_sess_create"].(int)
		ret.Sess_aged_out = in["sess_aged_out"].(int)
		ret.Tcp_total_drop = in["tcp_total_drop"].(int)
		ret.Tcp_dst_drop = in["tcp_dst_drop"].(int)
		ret.Tcp_src_drop = in["tcp_src_drop"].(int)
		ret.Tcp_src_dst_drop = in["tcp_src_dst_drop"].(int)
		ret.Udp_total_drop = in["udp_total_drop"].(int)
		ret.Udp_dst_drop = in["udp_dst_drop"].(int)
		ret.Udp_src_drop = in["udp_src_drop"].(int)
		ret.Udp_src_dst_drop = in["udp_src_dst_drop"].(int)
		ret.Icmp_total_drop = in["icmp_total_drop"].(int)
		ret.Icmp_dst_drop = in["icmp_dst_drop"].(int)
		ret.Icmp_src_drop = in["icmp_src_drop"].(int)
		ret.Icmp_src_dst_drop = in["icmp_src_dst_drop"].(int)
		ret.Other_total_drop = in["other_total_drop"].(int)
		ret.Other_dst_drop = in["other_dst_drop"].(int)
		ret.Other_src_drop = in["other_src_drop"].(int)
		ret.Other_src_dst_drop = in["other_src_dst_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Dst_port_undef_drop = in["dst_port_undef_drop"].(int)
		ret.Dst_port_exceed_drop_any = in["dst_port_exceed_drop_any"].(int)
		ret.Dst_ipproto_bl = in["dst_ipproto_bl"].(int)
		ret.Dst_port_bl = in["dst_port_bl"].(int)
		ret.Dst_sport_bl = in["dst_sport_bl"].(int)
		ret.Dst_sport_exceed_drop_any = in["dst_sport_exceed_drop_any"].(int)
		ret.Dst_ipproto_rcvd = in["dst_ipproto_rcvd"].(int)
		ret.Dst_ipproto_drop = in["dst_ipproto_drop"].(int)
		ret.Dst_ipproto_exceed_drop_any = in["dst_ipproto_exceed_drop_any"].(int)
		ret.Src_ip_bypass = in["src_ip_bypass"].(int)
		ret.Dst_ingress_bytes = in["dst_ingress_bytes"].(int)
		ret.Dst_egress_bytes = in["dst_egress_bytes"].(int)
		ret.Dst_ingress_packets = in["dst_ingress_packets"].(int)
		ret.Dst_egress_packets = in["dst_egress_packets"].(int)
		ret.Dst_ip_bypass = in["dst_ip_bypass"].(int)
		ret.Dst_blackhole_inject = in["dst_blackhole_inject"].(int)
		ret.Dst_blackhole_withdraw = in["dst_blackhole_withdraw"].(int)
		ret.Tcp_total_bytes_rcv = in["tcp_total_bytes_rcv"].(int)
		ret.Tcp_total_bytes_drop = in["tcp_total_bytes_drop"].(int)
		ret.Udp_total_bytes_rcv = in["udp_total_bytes_rcv"].(int)
		ret.Udp_total_bytes_drop = in["udp_total_bytes_drop"].(int)
		ret.Icmp_total_bytes_rcv = in["icmp_total_bytes_rcv"].(int)
		ret.Icmp_total_bytes_drop = in["icmp_total_bytes_drop"].(int)
		ret.Other_total_bytes_rcv = in["other_total_bytes_rcv"].(int)
		ret.Other_total_bytes_drop = in["other_total_bytes_drop"].(int)
		ret.Udp_any_exceed = in["udp_any_exceed"].(int)
		ret.Tcp_any_exceed = in["tcp_any_exceed"].(int)
		ret.Icmp_any_exceed = in["icmp_any_exceed"].(int)
		ret.Other_any_exceed = in["other_any_exceed"].(int)
		ret.Tcp_drop_bl = in["tcp_drop_bl"].(int)
		ret.Udp_drop_bl = in["udp_drop_bl"].(int)
		ret.Icmp_drop_bl = in["icmp_drop_bl"].(int)
		ret.Other_drop_bl = in["other_drop_bl"].(int)
		ret.Glid_action_encap_send_immed = in["glid_action_encap_send_immed"].(int)
		ret.Glid_action_encap_send_delay = in["glid_action_encap_send_delay"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Dst_hw_drop_rule_inserted = in["dst_hw_drop_rule_inserted"].(int)
		ret.Dst_hw_drop_rule_removed = in["dst_hw_drop_rule_removed"].(int)
		ret.Src_hw_drop_rule_inserted = in["src_hw_drop_rule_inserted"].(int)
		ret.Src_hw_drop_rule_removed = in["src_hw_drop_rule_removed"].(int)
	}
	return ret
}

func dataToEndpointDdosBriefStats(d *schema.ResourceData) edpt.DdosBriefStats {
	var ret edpt.DdosBriefStats

	ret.Stats = getObjectDdosBriefStatsStats(d.Get("stats").([]interface{}))
	return ret
}
