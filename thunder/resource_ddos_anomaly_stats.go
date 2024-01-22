package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosAnomalyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_anomaly_stats`: Statistics for the object anomaly\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosAnomalyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"land_attack": {
							Type: schema.TypeInt, Optional: true, Description: "IP Land Attack",
						},
						"empty_frag": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Empty Fragment",
						},
						"micro_frag": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Micro Fragment",
						},
						"ipv4_opt": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Option",
						},
						"ip_frag": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment",
						},
						"bad_ip_hdr_len": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Invalid Header Length",
						},
						"bad_ip_flags": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Invalid Flags",
						},
						"bad_ip_ttl": {
							Type: schema.TypeInt, Optional: true, Description: "IP Invalid TTL",
						},
						"no_ip_payload": {
							Type: schema.TypeInt, Optional: true, Description: "IP Payload None",
						},
						"oversize_ip_pl": {
							Type: schema.TypeInt, Optional: true, Description: "IP Payload Too Large",
						},
						"bad_ip_pl_len": {
							Type: schema.TypeInt, Optional: true, Description: "IP Invalid Payload Length",
						},
						"bad_ip_frag_off": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Invalid Fragment Offset",
						},
						"bad_ip_csum": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Invalid Checksum",
						},
						"icmp_pod": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Ping Of Death",
						},
						"tcp_bad_urg_off": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Invalid Urgent Offset",
						},
						"tcp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Short Header",
						},
						"tcp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Invalid IPv4 Length",
						},
						"tcp_null_flags": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Null Flags",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Null Scan",
						},
						"tcp_syn_fin": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN&FIN",
						},
						"tcp_xmas_flags": {
							Type: schema.TypeInt, Optional: true, Description: "TCP XMAS Flags",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Description: "TCP XMAS Scan",
						},
						"tcp_syn_frag": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Fragment",
						},
						"tcp_frag_header": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Fragment Header",
						},
						"tcp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Invalid Checksum",
						},
						"udp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Short Header",
						},
						"udp_short_leng": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Invalid Length",
						},
						"udp_kb_frag": {
							Type: schema.TypeInt, Optional: true, Description: "UDP KB Fragment",
						},
						"udp_port_lb": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port LB",
						},
						"udp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Invalid Checksum",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "IP Runt Header",
						},
						"runt_tcpudp_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "TCPUDP Runt Header",
						},
						"tun_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "IP Tunnel Mismatch",
						},
						"tcp_opt_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Option Error",
						},
					},
				},
			},
		},
	}
}

func resourceDdosAnomalyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosAnomalyStatsStats := setObjectDdosAnomalyStatsStats(res)
		d.Set("stats", DdosAnomalyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosAnomalyStatsStats(ret edpt.DataDdosAnomalyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"land_attack":      ret.DtDdosAnomalyStats.Stats.Land_attack,
			"empty_frag":       ret.DtDdosAnomalyStats.Stats.Empty_frag,
			"micro_frag":       ret.DtDdosAnomalyStats.Stats.Micro_frag,
			"ipv4_opt":         ret.DtDdosAnomalyStats.Stats.Ipv4_opt,
			"ip_frag":          ret.DtDdosAnomalyStats.Stats.Ip_frag,
			"bad_ip_hdr_len":   ret.DtDdosAnomalyStats.Stats.Bad_ip_hdr_len,
			"bad_ip_flags":     ret.DtDdosAnomalyStats.Stats.Bad_ip_flags,
			"bad_ip_ttl":       ret.DtDdosAnomalyStats.Stats.Bad_ip_ttl,
			"no_ip_payload":    ret.DtDdosAnomalyStats.Stats.No_ip_payload,
			"oversize_ip_pl":   ret.DtDdosAnomalyStats.Stats.Oversize_ip_pl,
			"bad_ip_pl_len":    ret.DtDdosAnomalyStats.Stats.Bad_ip_pl_len,
			"bad_ip_frag_off":  ret.DtDdosAnomalyStats.Stats.Bad_ip_frag_off,
			"bad_ip_csum":      ret.DtDdosAnomalyStats.Stats.Bad_ip_csum,
			"icmp_pod":         ret.DtDdosAnomalyStats.Stats.Icmp_pod,
			"tcp_bad_urg_off":  ret.DtDdosAnomalyStats.Stats.Tcp_bad_urg_off,
			"tcp_short_hdr":    ret.DtDdosAnomalyStats.Stats.Tcp_short_hdr,
			"tcp_bad_ip_len":   ret.DtDdosAnomalyStats.Stats.Tcp_bad_ip_len,
			"tcp_null_flags":   ret.DtDdosAnomalyStats.Stats.Tcp_null_flags,
			"tcp_null_scan":    ret.DtDdosAnomalyStats.Stats.Tcp_null_scan,
			"tcp_syn_fin":      ret.DtDdosAnomalyStats.Stats.Tcp_syn_fin,
			"tcp_xmas_flags":   ret.DtDdosAnomalyStats.Stats.Tcp_xmas_flags,
			"tcp_xmas_scan":    ret.DtDdosAnomalyStats.Stats.Tcp_xmas_scan,
			"tcp_syn_frag":     ret.DtDdosAnomalyStats.Stats.Tcp_syn_frag,
			"tcp_frag_header":  ret.DtDdosAnomalyStats.Stats.Tcp_frag_header,
			"tcp_bad_csum":     ret.DtDdosAnomalyStats.Stats.Tcp_bad_csum,
			"udp_short_hdr":    ret.DtDdosAnomalyStats.Stats.Udp_short_hdr,
			"udp_short_leng":   ret.DtDdosAnomalyStats.Stats.Udp_short_leng,
			"udp_kb_frag":      ret.DtDdosAnomalyStats.Stats.Udp_kb_frag,
			"udp_port_lb":      ret.DtDdosAnomalyStats.Stats.Udp_port_lb,
			"udp_bad_csum":     ret.DtDdosAnomalyStats.Stats.Udp_bad_csum,
			"runt_ip_hdr":      ret.DtDdosAnomalyStats.Stats.Runt_ip_hdr,
			"runt_tcpudp_hdr":  ret.DtDdosAnomalyStats.Stats.Runt_tcpudp_hdr,
			"tun_mismatch":     ret.DtDdosAnomalyStats.Stats.Tun_mismatch,
			"tcp_opt_overflow": ret.DtDdosAnomalyStats.Stats.Tcp_opt_overflow,
		},
	}
}

func getObjectDdosAnomalyStatsStats(d []interface{}) edpt.DdosAnomalyStatsStats {

	count1 := len(d)
	var ret edpt.DdosAnomalyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Land_attack = in["land_attack"].(int)
		ret.Empty_frag = in["empty_frag"].(int)
		ret.Micro_frag = in["micro_frag"].(int)
		ret.Ipv4_opt = in["ipv4_opt"].(int)
		ret.Ip_frag = in["ip_frag"].(int)
		ret.Bad_ip_hdr_len = in["bad_ip_hdr_len"].(int)
		ret.Bad_ip_flags = in["bad_ip_flags"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Oversize_ip_pl = in["oversize_ip_pl"].(int)
		ret.Bad_ip_pl_len = in["bad_ip_pl_len"].(int)
		ret.Bad_ip_frag_off = in["bad_ip_frag_off"].(int)
		ret.Bad_ip_csum = in["bad_ip_csum"].(int)
		ret.Icmp_pod = in["icmp_pod"].(int)
		ret.Tcp_bad_urg_off = in["tcp_bad_urg_off"].(int)
		ret.Tcp_short_hdr = in["tcp_short_hdr"].(int)
		ret.Tcp_bad_ip_len = in["tcp_bad_ip_len"].(int)
		ret.Tcp_null_flags = in["tcp_null_flags"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas_flags = in["tcp_xmas_flags"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frag = in["tcp_syn_frag"].(int)
		ret.Tcp_frag_header = in["tcp_frag_header"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_short_hdr = in["udp_short_hdr"].(int)
		ret.Udp_short_leng = in["udp_short_leng"].(int)
		ret.Udp_kb_frag = in["udp_kb_frag"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcpudp_hdr = in["runt_tcpudp_hdr"].(int)
		ret.Tun_mismatch = in["tun_mismatch"].(int)
		ret.Tcp_opt_overflow = in["tcp_opt_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosAnomalyStats(d *schema.ResourceData) edpt.DdosAnomalyStats {
	var ret edpt.DdosAnomalyStats

	ret.Stats = getObjectDdosAnomalyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
