package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAnomalyDropStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_anomaly_drop_stats`: Statistics for the object anomaly-drop\n\n__PLACEHOLDER__",
		ReadContext: resourceIpAnomalyDropStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"land": {
							Type: schema.TypeInt, Optional: true, Description: "Land Attack Drop",
						},
						"emp_frg": {
							Type: schema.TypeInt, Optional: true, Description: "Empty Fragment Drop",
						},
						"emp_mic_frg": {
							Type: schema.TypeInt, Optional: true, Description: "Micro Fragment Drop",
						},
						"opt": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Options Drop",
						},
						"frg": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Fragment Drop",
						},
						"bad_ip_hdrlen": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Header Len Drop",
						},
						"bad_ip_flg": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Flags Drop",
						},
						"bad_ip_ttl": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP TTL Drop",
						},
						"no_ip_payload": {
							Type: schema.TypeInt, Optional: true, Description: "No IP Payload drop",
						},
						"over_ip_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Oversize IP Payload Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Payload Len Drop",
						},
						"bad_ip_frg_offset": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Fragment Offset Drop",
						},
						"csum": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Checksum Drop",
						},
						"pod": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Ping of Death Drop",
						},
						"bad_tcp_urg_offset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Bad Urgent Offset Drop",
						},
						"tcp_sht_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Short Header Drop",
						},
						"tcp_bad_iplen": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Bad IP Length Drop",
						},
						"tcp_null_frg": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Null Flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Null Scan Drop",
						},
						"tcp_syn_fin": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Syn and Fin Drop",
						},
						"tcp_xmas": {
							Type: schema.TypeInt, Optional: true, Description: "TCP XMAS Flags Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Description: "TCP XMAS Scan Drop",
						},
						"tcp_syn_frg": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Syn Fragment Drop",
						},
						"tcp_frg_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Fragmented Header Drop",
						},
						"tcp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Bad Checksum Drop",
						},
						"udp_srt_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Short Header Drop",
						},
						"udp_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Bad Length Drop",
						},
						"udp_kerb_frg": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Kerberos Fragment Drop",
						},
						"udp_port_lb": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Loopback Drop",
						},
						"udp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Bad Checksum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Runt IP Header Drop",
						},
						"runt_tcp_udp_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Runt TCP/UDP Header Drop",
						},
						"ipip_tnl_msmtch": {
							Type: schema.TypeInt, Optional: true, Description: "IP-over-IP Tunnel Mismatch Drop",
						},
						"tcp_opt_err": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Option Error Drop",
						},
						"ipip_tnl_err": {
							Type: schema.TypeInt, Optional: true, Description: "IP-over-IP Tunnel Error Drop",
						},
						"vxlan_err": {
							Type: schema.TypeInt, Optional: true, Description: "VXLAN Tunnel Error Drop",
						},
						"nvgre_err": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Error Drop",
						},
						"gre_pptp_err": {
							Type: schema.TypeInt, Optional: true, Description: "GRE PPTP Error Drop",
						},
						"ipv6_eh_hbh": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Hop by Hop Header Drop",
						},
						"ipv6_eh_dest": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Destination Header Drop",
						},
						"ipv6_eh_routing": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Routing Header Drop",
						},
						"ipv6_eh_frag": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Fragmentation Header Drop",
						},
						"ipv6_eh_ah": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Authentication Header Drop",
						},
						"ipv6_eh_esp": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 ESP Header Drop",
						},
						"ipv6_eh_mobility": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Mobility Header Drop",
						},
						"ipv6_eh_none": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 No Next Header Drop",
						},
						"ipv6_eh_other": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Unknown Extension Header Drop",
						},
						"ipv6_eh_malformed": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Malformed Extension Header Drop",
						},
					},
				},
			},
		},
	}
}

func resourceIpAnomalyDropStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAnomalyDropStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAnomalyDropStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpAnomalyDropStatsStats := setObjectIpAnomalyDropStatsStats(res)
		d.Set("stats", IpAnomalyDropStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpAnomalyDropStatsStats(ret edpt.DataIpAnomalyDropStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"land":               ret.DtIpAnomalyDropStats.Stats.Land,
			"emp_frg":            ret.DtIpAnomalyDropStats.Stats.Emp_frg,
			"emp_mic_frg":        ret.DtIpAnomalyDropStats.Stats.Emp_mic_frg,
			"opt":                ret.DtIpAnomalyDropStats.Stats.Opt,
			"frg":                ret.DtIpAnomalyDropStats.Stats.Frg,
			"bad_ip_hdrlen":      ret.DtIpAnomalyDropStats.Stats.Bad_ip_hdrlen,
			"bad_ip_flg":         ret.DtIpAnomalyDropStats.Stats.Bad_ip_flg,
			"bad_ip_ttl":         ret.DtIpAnomalyDropStats.Stats.Bad_ip_ttl,
			"no_ip_payload":      ret.DtIpAnomalyDropStats.Stats.No_ip_payload,
			"over_ip_payload":    ret.DtIpAnomalyDropStats.Stats.Over_ip_payload,
			"bad_ip_payload_len": ret.DtIpAnomalyDropStats.Stats.Bad_ip_payload_len,
			"bad_ip_frg_offset":  ret.DtIpAnomalyDropStats.Stats.Bad_ip_frg_offset,
			"csum":               ret.DtIpAnomalyDropStats.Stats.Csum,
			"pod":                ret.DtIpAnomalyDropStats.Stats.Pod,
			"bad_tcp_urg_offset": ret.DtIpAnomalyDropStats.Stats.Bad_tcp_urg_offset,
			"tcp_sht_hdr":        ret.DtIpAnomalyDropStats.Stats.Tcp_sht_hdr,
			"tcp_bad_iplen":      ret.DtIpAnomalyDropStats.Stats.Tcp_bad_iplen,
			"tcp_null_frg":       ret.DtIpAnomalyDropStats.Stats.Tcp_null_frg,
			"tcp_null_scan":      ret.DtIpAnomalyDropStats.Stats.Tcp_null_scan,
			"tcp_syn_fin":        ret.DtIpAnomalyDropStats.Stats.Tcp_syn_fin,
			"tcp_xmas":           ret.DtIpAnomalyDropStats.Stats.Tcp_xmas,
			"tcp_xmas_scan":      ret.DtIpAnomalyDropStats.Stats.Tcp_xmas_scan,
			"tcp_syn_frg":        ret.DtIpAnomalyDropStats.Stats.Tcp_syn_frg,
			"tcp_frg_hdr":        ret.DtIpAnomalyDropStats.Stats.Tcp_frg_hdr,
			"tcp_bad_csum":       ret.DtIpAnomalyDropStats.Stats.Tcp_bad_csum,
			"udp_srt_hdr":        ret.DtIpAnomalyDropStats.Stats.Udp_srt_hdr,
			"udp_bad_len":        ret.DtIpAnomalyDropStats.Stats.Udp_bad_len,
			"udp_kerb_frg":       ret.DtIpAnomalyDropStats.Stats.Udp_kerb_frg,
			"udp_port_lb":        ret.DtIpAnomalyDropStats.Stats.Udp_port_lb,
			"udp_bad_csum":       ret.DtIpAnomalyDropStats.Stats.Udp_bad_csum,
			"runt_ip_hdr":        ret.DtIpAnomalyDropStats.Stats.Runt_ip_hdr,
			"runt_tcp_udp_hdr":   ret.DtIpAnomalyDropStats.Stats.Runt_tcp_udp_hdr,
			"ipip_tnl_msmtch":    ret.DtIpAnomalyDropStats.Stats.Ipip_tnl_msmtch,
			"tcp_opt_err":        ret.DtIpAnomalyDropStats.Stats.Tcp_opt_err,
			"ipip_tnl_err":       ret.DtIpAnomalyDropStats.Stats.Ipip_tnl_err,
			"vxlan_err":          ret.DtIpAnomalyDropStats.Stats.Vxlan_err,
			"nvgre_err":          ret.DtIpAnomalyDropStats.Stats.Nvgre_err,
			"gre_pptp_err":       ret.DtIpAnomalyDropStats.Stats.Gre_pptp_err,
			"ipv6_eh_hbh":        ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_hbh,
			"ipv6_eh_dest":       ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_dest,
			"ipv6_eh_routing":    ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_routing,
			"ipv6_eh_frag":       ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_frag,
			"ipv6_eh_ah":         ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_ah,
			"ipv6_eh_esp":        ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_esp,
			"ipv6_eh_mobility":   ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_mobility,
			"ipv6_eh_none":       ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_none,
			"ipv6_eh_other":      ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_other,
			"ipv6_eh_malformed":  ret.DtIpAnomalyDropStats.Stats.Ipv6_eh_malformed,
		},
	}
}

func getObjectIpAnomalyDropStatsStats(d []interface{}) edpt.IpAnomalyDropStatsStats {

	count1 := len(d)
	var ret edpt.IpAnomalyDropStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Land = in["land"].(int)
		ret.Emp_frg = in["emp_frg"].(int)
		ret.Emp_mic_frg = in["emp_mic_frg"].(int)
		ret.Opt = in["opt"].(int)
		ret.Frg = in["frg"].(int)
		ret.Bad_ip_hdrlen = in["bad_ip_hdrlen"].(int)
		ret.Bad_ip_flg = in["bad_ip_flg"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Over_ip_payload = in["over_ip_payload"].(int)
		ret.Bad_ip_payload_len = in["bad_ip_payload_len"].(int)
		ret.Bad_ip_frg_offset = in["bad_ip_frg_offset"].(int)
		ret.Csum = in["csum"].(int)
		ret.Pod = in["pod"].(int)
		ret.Bad_tcp_urg_offset = in["bad_tcp_urg_offset"].(int)
		ret.Tcp_sht_hdr = in["tcp_sht_hdr"].(int)
		ret.Tcp_bad_iplen = in["tcp_bad_iplen"].(int)
		ret.Tcp_null_frg = in["tcp_null_frg"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas = in["tcp_xmas"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frg = in["tcp_syn_frg"].(int)
		ret.Tcp_frg_hdr = in["tcp_frg_hdr"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_srt_hdr = in["udp_srt_hdr"].(int)
		ret.Udp_bad_len = in["udp_bad_len"].(int)
		ret.Udp_kerb_frg = in["udp_kerb_frg"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcp_udp_hdr = in["runt_tcp_udp_hdr"].(int)
		ret.Ipip_tnl_msmtch = in["ipip_tnl_msmtch"].(int)
		ret.Tcp_opt_err = in["tcp_opt_err"].(int)
		ret.Ipip_tnl_err = in["ipip_tnl_err"].(int)
		ret.Vxlan_err = in["vxlan_err"].(int)
		ret.Nvgre_err = in["nvgre_err"].(int)
		ret.Gre_pptp_err = in["gre_pptp_err"].(int)
		ret.Ipv6_eh_hbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6_eh_dest = in["ipv6_eh_dest"].(int)
		ret.Ipv6_eh_routing = in["ipv6_eh_routing"].(int)
		ret.Ipv6_eh_frag = in["ipv6_eh_frag"].(int)
		ret.Ipv6_eh_ah = in["ipv6_eh_ah"].(int)
		ret.Ipv6_eh_esp = in["ipv6_eh_esp"].(int)
		ret.Ipv6_eh_mobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6_eh_none = in["ipv6_eh_none"].(int)
		ret.Ipv6_eh_other = in["ipv6_eh_other"].(int)
		ret.Ipv6_eh_malformed = in["ipv6_eh_malformed"].(int)
	}
	return ret
}

func dataToEndpointIpAnomalyDropStats(d *schema.ResourceData) edpt.IpAnomalyDropStats {
	var ret edpt.IpAnomalyDropStats

	ret.Stats = getObjectIpAnomalyDropStatsStats(d.Get("stats").([]interface{}))
	return ret
}
