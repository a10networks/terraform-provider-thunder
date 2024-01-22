package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SixrdDomainStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_sixrd_domain_stats`: Statistics for the object domain\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6SixrdDomainStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "6rd Domain name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outbound_tcp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound TCP packets received",
						},
						"outbound_udp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound UDP packets received",
						},
						"outbound_icmp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound ICMP packets received",
						},
						"outbound_other_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound other packets received",
						},
						"outbound_packets_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound packets dropped",
						},
						"outbound_ipv6_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound IPv6 destination unreachable",
						},
						"outbound_fragment_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Fragmented IPv6",
						},
						"inbound_tcp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound TCP packets received",
						},
						"inbound_udp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound UDP packets received",
						},
						"inbound_icmp_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound ICMP packets received",
						},
						"inbound_other_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound other packets received",
						},
						"inbound_packets_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound packets dropped",
						},
						"inbound_ipv4_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound IPv4 destination unreachable",
						},
						"inbound_fragment_ipv4": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Fragmented IPv4",
						},
						"inbound_tunnel_fragment_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Fragmented IPv6 in tunnel",
						},
						"vport_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Traffic match SLB virtual port",
						},
						"unknown_delegated_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown 6rd delegated prefix",
						},
						"packet_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Packet too big",
						},
						"not_local_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Not local IP",
						},
						"fragment_error": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment processing errors",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other errors",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6SixrdDomainStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdDomainStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdDomainStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6SixrdDomainStatsStats := setObjectCgnv6SixrdDomainStatsStats(res)
		d.Set("stats", Cgnv6SixrdDomainStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6SixrdDomainStatsStats(ret edpt.DataCgnv6SixrdDomainStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"outbound_tcp_packets_received":   ret.DtCgnv6SixrdDomainStats.Stats.OutboundTcpPacketsReceived,
			"outbound_udp_packets_received":   ret.DtCgnv6SixrdDomainStats.Stats.OutboundUdpPacketsReceived,
			"outbound_icmp_packets_received":  ret.DtCgnv6SixrdDomainStats.Stats.OutboundIcmpPacketsReceived,
			"outbound_other_packets_received": ret.DtCgnv6SixrdDomainStats.Stats.OutboundOtherPacketsReceived,
			"outbound_packets_drop":           ret.DtCgnv6SixrdDomainStats.Stats.OutboundPacketsDrop,
			"outbound_ipv6_dest_unreachable":  ret.DtCgnv6SixrdDomainStats.Stats.OutboundIpv6DestUnreachable,
			"outbound_fragment_ipv6":          ret.DtCgnv6SixrdDomainStats.Stats.OutboundFragmentIpv6,
			"inbound_tcp_packets_received":    ret.DtCgnv6SixrdDomainStats.Stats.InboundTcpPacketsReceived,
			"inbound_udp_packets_received":    ret.DtCgnv6SixrdDomainStats.Stats.InboundUdpPacketsReceived,
			"inbound_icmp_packets_received":   ret.DtCgnv6SixrdDomainStats.Stats.InboundIcmpPacketsReceived,
			"inbound_other_packets_received":  ret.DtCgnv6SixrdDomainStats.Stats.InboundOtherPacketsReceived,
			"inbound_packets_drop":            ret.DtCgnv6SixrdDomainStats.Stats.InboundPacketsDrop,
			"inbound_ipv4_dest_unreachable":   ret.DtCgnv6SixrdDomainStats.Stats.InboundIpv4DestUnreachable,
			"inbound_fragment_ipv4":           ret.DtCgnv6SixrdDomainStats.Stats.InboundFragmentIpv4,
			"inbound_tunnel_fragment_ipv6":    ret.DtCgnv6SixrdDomainStats.Stats.InboundTunnelFragmentIpv6,
			"vport_matched":                   ret.DtCgnv6SixrdDomainStats.Stats.VportMatched,
			"unknown_delegated_prefix":        ret.DtCgnv6SixrdDomainStats.Stats.UnknownDelegatedPrefix,
			"packet_too_big":                  ret.DtCgnv6SixrdDomainStats.Stats.PacketTooBig,
			"not_local_ip":                    ret.DtCgnv6SixrdDomainStats.Stats.NotLocalIp,
			"fragment_error":                  ret.DtCgnv6SixrdDomainStats.Stats.FragmentError,
			"other_error":                     ret.DtCgnv6SixrdDomainStats.Stats.OtherError,
		},
	}
}

func getObjectCgnv6SixrdDomainStatsStats(d []interface{}) edpt.Cgnv6SixrdDomainStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6SixrdDomainStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutboundTcpPacketsReceived = in["outbound_tcp_packets_received"].(int)
		ret.OutboundUdpPacketsReceived = in["outbound_udp_packets_received"].(int)
		ret.OutboundIcmpPacketsReceived = in["outbound_icmp_packets_received"].(int)
		ret.OutboundOtherPacketsReceived = in["outbound_other_packets_received"].(int)
		ret.OutboundPacketsDrop = in["outbound_packets_drop"].(int)
		ret.OutboundIpv6DestUnreachable = in["outbound_ipv6_dest_unreachable"].(int)
		ret.OutboundFragmentIpv6 = in["outbound_fragment_ipv6"].(int)
		ret.InboundTcpPacketsReceived = in["inbound_tcp_packets_received"].(int)
		ret.InboundUdpPacketsReceived = in["inbound_udp_packets_received"].(int)
		ret.InboundIcmpPacketsReceived = in["inbound_icmp_packets_received"].(int)
		ret.InboundOtherPacketsReceived = in["inbound_other_packets_received"].(int)
		ret.InboundPacketsDrop = in["inbound_packets_drop"].(int)
		ret.InboundIpv4DestUnreachable = in["inbound_ipv4_dest_unreachable"].(int)
		ret.InboundFragmentIpv4 = in["inbound_fragment_ipv4"].(int)
		ret.InboundTunnelFragmentIpv6 = in["inbound_tunnel_fragment_ipv6"].(int)
		ret.VportMatched = in["vport_matched"].(int)
		ret.UnknownDelegatedPrefix = in["unknown_delegated_prefix"].(int)
		ret.PacketTooBig = in["packet_too_big"].(int)
		ret.NotLocalIp = in["not_local_ip"].(int)
		ret.FragmentError = in["fragment_error"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointCgnv6SixrdDomainStats(d *schema.ResourceData) edpt.Cgnv6SixrdDomainStats {
	var ret edpt.Cgnv6SixrdDomainStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectCgnv6SixrdDomainStatsStats(d.Get("stats").([]interface{}))
	return ret
}
