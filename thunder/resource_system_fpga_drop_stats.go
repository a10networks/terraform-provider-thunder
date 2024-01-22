package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemFpgaDropStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_fpga_drop_stats`: Statistics for the object fpga-drop\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemFpgaDropStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mrx_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total MRX Drop",
						},
						"hrx_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total HRX Drop",
						},
						"siz_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Size Drop",
						},
						"fcs_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total FCS Drop",
						},
						"land_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total LAND Attack Drop",
						},
						"empty_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Empty frag Drop",
						},
						"mic_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Micro frag Drop",
						},
						"ipv4_opt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total IPv4 opt Drop",
						},
						"ipv4_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Total IP frag Drop",
						},
						"bad_ip_hdr_len": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP hdr len Drop",
						},
						"bad_ip_flags_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP Flags Drop",
						},
						"bad_ip_ttl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP TTL Drop",
						},
						"no_ip_payload_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total No IP Payload Drop",
						},
						"oversize_ip_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Total Oversize IP PL Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP PL len Drop",
						},
						"bad_ip_frag_offset": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP frag off Drop",
						},
						"bad_ip_chksum_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bad IP csum Drop",
						},
						"icmp_pod_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total ICMP POD Drop",
						},
						"tcp_bad_urg_offet": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP bad urg off Drop",
						},
						"tcp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP short hdr Drop",
						},
						"tcp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP Bad IP Len Drop",
						},
						"tcp_null_flags": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP null flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP null scan Drop",
						},
						"tcp_fin_sin": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP SYN&FIN Drop",
						},
						"tcp_xmas_flags": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP XMAS FLAGS Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP XMAS scan Drop",
						},
						"tcp_syn_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP SYN frag Drop",
						},
						"tcp_frag_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP frag header Drop",
						},
						"tcp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP bad csum Drop",
						},
						"udp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP short hdr Drop",
						},
						"udp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP bad leng Drop",
						},
						"udp_kb_frags": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP KB frag Drop",
						},
						"udp_port_lb": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP port LB Drop",
						},
						"udp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP bad csum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Total Runt IP hdr Drop",
						},
						"runt_tcpudp_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Total Runt TCPUDP hdr Drop",
						},
						"tun_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Total Tun mismatch Drop",
						},
						"qdr_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total QDR Drop",
						},
					},
				},
			},
		},
	}
}

func resourceSystemFpgaDropStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaDropStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaDropStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemFpgaDropStatsStats := setObjectSystemFpgaDropStatsStats(res)
		d.Set("stats", SystemFpgaDropStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemFpgaDropStatsStats(ret edpt.DataSystemFpgaDropStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mrx_drop":            ret.DtSystemFpgaDropStats.Stats.MrxDrop,
			"hrx_drop":            ret.DtSystemFpgaDropStats.Stats.HrxDrop,
			"siz_drop":            ret.DtSystemFpgaDropStats.Stats.SizDrop,
			"fcs_drop":            ret.DtSystemFpgaDropStats.Stats.FcsDrop,
			"land_drop":           ret.DtSystemFpgaDropStats.Stats.LandDrop,
			"empty_frag_drop":     ret.DtSystemFpgaDropStats.Stats.EmptyFragDrop,
			"mic_frag_drop":       ret.DtSystemFpgaDropStats.Stats.MicFragDrop,
			"ipv4_opt_drop":       ret.DtSystemFpgaDropStats.Stats.Ipv4OptDrop,
			"ipv4_frag":           ret.DtSystemFpgaDropStats.Stats.Ipv4Frag,
			"bad_ip_hdr_len":      ret.DtSystemFpgaDropStats.Stats.BadIpHdrLen,
			"bad_ip_flags_drop":   ret.DtSystemFpgaDropStats.Stats.BadIpFlagsDrop,
			"bad_ip_ttl_drop":     ret.DtSystemFpgaDropStats.Stats.BadIpTtlDrop,
			"no_ip_payload_drop":  ret.DtSystemFpgaDropStats.Stats.NoIpPayloadDrop,
			"oversize_ip_payload": ret.DtSystemFpgaDropStats.Stats.OversizeIpPayload,
			"bad_ip_payload_len":  ret.DtSystemFpgaDropStats.Stats.BadIpPayloadLen,
			"bad_ip_frag_offset":  ret.DtSystemFpgaDropStats.Stats.BadIpFragOffset,
			"bad_ip_chksum_drop":  ret.DtSystemFpgaDropStats.Stats.BadIpChksumDrop,
			"icmp_pod_drop":       ret.DtSystemFpgaDropStats.Stats.IcmpPodDrop,
			"tcp_bad_urg_offet":   ret.DtSystemFpgaDropStats.Stats.TcpBadUrgOffet,
			"tcp_short_hdr":       ret.DtSystemFpgaDropStats.Stats.TcpShortHdr,
			"tcp_bad_ip_len":      ret.DtSystemFpgaDropStats.Stats.TcpBadIpLen,
			"tcp_null_flags":      ret.DtSystemFpgaDropStats.Stats.TcpNullFlags,
			"tcp_null_scan":       ret.DtSystemFpgaDropStats.Stats.TcpNullScan,
			"tcp_fin_sin":         ret.DtSystemFpgaDropStats.Stats.TcpFinSin,
			"tcp_xmas_flags":      ret.DtSystemFpgaDropStats.Stats.TcpXmasFlags,
			"tcp_xmas_scan":       ret.DtSystemFpgaDropStats.Stats.TcpXmasScan,
			"tcp_syn_frag":        ret.DtSystemFpgaDropStats.Stats.TcpSynFrag,
			"tcp_frag_hdr":        ret.DtSystemFpgaDropStats.Stats.TcpFragHdr,
			"tcp_bad_chksum":      ret.DtSystemFpgaDropStats.Stats.TcpBadChksum,
			"udp_short_hdr":       ret.DtSystemFpgaDropStats.Stats.UdpShortHdr,
			"udp_bad_ip_len":      ret.DtSystemFpgaDropStats.Stats.UdpBadIpLen,
			"udp_kb_frags":        ret.DtSystemFpgaDropStats.Stats.UdpKbFrags,
			"udp_port_lb":         ret.DtSystemFpgaDropStats.Stats.UdpPortLb,
			"udp_bad_chksum":      ret.DtSystemFpgaDropStats.Stats.UdpBadChksum,
			"runt_ip_hdr":         ret.DtSystemFpgaDropStats.Stats.RuntIpHdr,
			"runt_tcpudp_hdr":     ret.DtSystemFpgaDropStats.Stats.RuntTcpudpHdr,
			"tun_mismatch":        ret.DtSystemFpgaDropStats.Stats.TunMismatch,
			"qdr_drop":            ret.DtSystemFpgaDropStats.Stats.QdrDrop,
		},
	}
}

func getObjectSystemFpgaDropStatsStats(d []interface{}) edpt.SystemFpgaDropStatsStats {

	count1 := len(d)
	var ret edpt.SystemFpgaDropStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MrxDrop = in["mrx_drop"].(int)
		ret.HrxDrop = in["hrx_drop"].(int)
		ret.SizDrop = in["siz_drop"].(int)
		ret.FcsDrop = in["fcs_drop"].(int)
		ret.LandDrop = in["land_drop"].(int)
		ret.EmptyFragDrop = in["empty_frag_drop"].(int)
		ret.MicFragDrop = in["mic_frag_drop"].(int)
		ret.Ipv4OptDrop = in["ipv4_opt_drop"].(int)
		ret.Ipv4Frag = in["ipv4_frag"].(int)
		ret.BadIpHdrLen = in["bad_ip_hdr_len"].(int)
		ret.BadIpFlagsDrop = in["bad_ip_flags_drop"].(int)
		ret.BadIpTtlDrop = in["bad_ip_ttl_drop"].(int)
		ret.NoIpPayloadDrop = in["no_ip_payload_drop"].(int)
		ret.OversizeIpPayload = in["oversize_ip_payload"].(int)
		ret.BadIpPayloadLen = in["bad_ip_payload_len"].(int)
		ret.BadIpFragOffset = in["bad_ip_frag_offset"].(int)
		ret.BadIpChksumDrop = in["bad_ip_chksum_drop"].(int)
		ret.IcmpPodDrop = in["icmp_pod_drop"].(int)
		ret.TcpBadUrgOffet = in["tcp_bad_urg_offet"].(int)
		ret.TcpShortHdr = in["tcp_short_hdr"].(int)
		ret.TcpBadIpLen = in["tcp_bad_ip_len"].(int)
		ret.TcpNullFlags = in["tcp_null_flags"].(int)
		ret.TcpNullScan = in["tcp_null_scan"].(int)
		ret.TcpFinSin = in["tcp_fin_sin"].(int)
		ret.TcpXmasFlags = in["tcp_xmas_flags"].(int)
		ret.TcpXmasScan = in["tcp_xmas_scan"].(int)
		ret.TcpSynFrag = in["tcp_syn_frag"].(int)
		ret.TcpFragHdr = in["tcp_frag_hdr"].(int)
		ret.TcpBadChksum = in["tcp_bad_chksum"].(int)
		ret.UdpShortHdr = in["udp_short_hdr"].(int)
		ret.UdpBadIpLen = in["udp_bad_ip_len"].(int)
		ret.UdpKbFrags = in["udp_kb_frags"].(int)
		ret.UdpPortLb = in["udp_port_lb"].(int)
		ret.UdpBadChksum = in["udp_bad_chksum"].(int)
		ret.RuntIpHdr = in["runt_ip_hdr"].(int)
		ret.RuntTcpudpHdr = in["runt_tcpudp_hdr"].(int)
		ret.TunMismatch = in["tun_mismatch"].(int)
		ret.QdrDrop = in["qdr_drop"].(int)
	}
	return ret
}

func dataToEndpointSystemFpgaDropStats(d *schema.ResourceData) edpt.SystemFpgaDropStats {
	var ret edpt.SystemFpgaDropStats

	ret.Stats = getObjectSystemFpgaDropStatsStats(d.Get("stats").([]interface{}))
	return ret
}
