package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSessionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_session_stats`: Statistics for the object session\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSessionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sess_create_ip": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Session Created",
						},
						"sess_create_ip6": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Session Created",
						},
						"sess_create_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Created",
						},
						"conn_tcp_est_w_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Established With SYN",
						},
						"conn_tcp_est_w_ack": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Established With ACK",
						},
						"conn_tcp_est": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Established",
						},
						"conn_tcp_close_w_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Closed With RST",
						},
						"conn_tcp_close_w_fin": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Closed With FIN",
						},
						"conn_tcp_close_w_idle": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Closed With Idle",
						},
						"conn_tcp_close_w_half_open": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Closed With Half Open",
						},
						"conn_tcp_close": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Closed",
						},
						"sess_create_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Created",
						},
						"conn_udp_est": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Established",
						},
						"conn_udp_close": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Closed",
						},
						"sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "Session Aged Out",
						},
						"conn_entry_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Entry Mismatch Dropped",
						},
						"conn_tcp_create_from_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Created From SYN",
						},
						"conn_tcp_create_from_ack": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Created From ACK",
						},
						"sess_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Session Memory",
						},
						"sess_create_udp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Helper Auth Sessions Created",
						},
						"sess_aged_udp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Helper Auth Sessions Aged Out",
						},
						"sess_snat_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Session Source NAT Failure",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSessionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSessionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSessionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSessionStatsStats := setObjectDdosSessionStatsStats(res)
		d.Set("stats", DdosSessionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSessionStatsStats(ret edpt.DataDdosSessionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sess_create_ip":             ret.DtDdosSessionStats.Stats.Sess_create_ip,
			"sess_create_ip6":            ret.DtDdosSessionStats.Stats.Sess_create_ip6,
			"sess_create_tcp":            ret.DtDdosSessionStats.Stats.Sess_create_tcp,
			"conn_tcp_est_w_syn":         ret.DtDdosSessionStats.Stats.Conn_tcp_est_w_syn,
			"conn_tcp_est_w_ack":         ret.DtDdosSessionStats.Stats.Conn_tcp_est_w_ack,
			"conn_tcp_est":               ret.DtDdosSessionStats.Stats.Conn_tcp_est,
			"conn_tcp_close_w_rst":       ret.DtDdosSessionStats.Stats.Conn_tcp_close_w_rst,
			"conn_tcp_close_w_fin":       ret.DtDdosSessionStats.Stats.Conn_tcp_close_w_fin,
			"conn_tcp_close_w_idle":      ret.DtDdosSessionStats.Stats.Conn_tcp_close_w_idle,
			"conn_tcp_close_w_half_open": ret.DtDdosSessionStats.Stats.Conn_tcp_close_w_half_open,
			"conn_tcp_close":             ret.DtDdosSessionStats.Stats.Conn_tcp_close,
			"sess_create_udp":            ret.DtDdosSessionStats.Stats.Sess_create_udp,
			"conn_udp_est":               ret.DtDdosSessionStats.Stats.Conn_udp_est,
			"conn_udp_close":             ret.DtDdosSessionStats.Stats.Conn_udp_close,
			"sess_aged_out":              ret.DtDdosSessionStats.Stats.Sess_aged_out,
			"conn_entry_mismatch":        ret.DtDdosSessionStats.Stats.Conn_entry_mismatch,
			"conn_tcp_create_from_syn":   ret.DtDdosSessionStats.Stats.Conn_tcp_create_from_syn,
			"conn_tcp_create_from_ack":   ret.DtDdosSessionStats.Stats.Conn_tcp_create_from_ack,
			"sess_oom":                   ret.DtDdosSessionStats.Stats.Sess_oom,
			"sess_create_udp_auth":       ret.DtDdosSessionStats.Stats.Sess_create_udp_auth,
			"sess_aged_udp_auth":         ret.DtDdosSessionStats.Stats.Sess_aged_udp_auth,
			"sess_snat_failed":           ret.DtDdosSessionStats.Stats.Sess_snat_failed,
		},
	}
}

func getObjectDdosSessionStatsStats(d []interface{}) edpt.DdosSessionStatsStats {

	count1 := len(d)
	var ret edpt.DdosSessionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sess_create_ip = in["sess_create_ip"].(int)
		ret.Sess_create_ip6 = in["sess_create_ip6"].(int)
		ret.Sess_create_tcp = in["sess_create_tcp"].(int)
		ret.Conn_tcp_est_w_syn = in["conn_tcp_est_w_syn"].(int)
		ret.Conn_tcp_est_w_ack = in["conn_tcp_est_w_ack"].(int)
		ret.Conn_tcp_est = in["conn_tcp_est"].(int)
		ret.Conn_tcp_close_w_rst = in["conn_tcp_close_w_rst"].(int)
		ret.Conn_tcp_close_w_fin = in["conn_tcp_close_w_fin"].(int)
		ret.Conn_tcp_close_w_idle = in["conn_tcp_close_w_idle"].(int)
		ret.Conn_tcp_close_w_half_open = in["conn_tcp_close_w_half_open"].(int)
		ret.Conn_tcp_close = in["conn_tcp_close"].(int)
		ret.Sess_create_udp = in["sess_create_udp"].(int)
		ret.Conn_udp_est = in["conn_udp_est"].(int)
		ret.Conn_udp_close = in["conn_udp_close"].(int)
		ret.Sess_aged_out = in["sess_aged_out"].(int)
		ret.Conn_entry_mismatch = in["conn_entry_mismatch"].(int)
		ret.Conn_tcp_create_from_syn = in["conn_tcp_create_from_syn"].(int)
		ret.Conn_tcp_create_from_ack = in["conn_tcp_create_from_ack"].(int)
		ret.Sess_oom = in["sess_oom"].(int)
		ret.Sess_create_udp_auth = in["sess_create_udp_auth"].(int)
		ret.Sess_aged_udp_auth = in["sess_aged_udp_auth"].(int)
		ret.Sess_snat_failed = in["sess_snat_failed"].(int)
	}
	return ret
}

func dataToEndpointDdosSessionStats(d *schema.ResourceData) edpt.DdosSessionStats {
	var ret edpt.DdosSessionStats

	ret.Stats = getObjectDdosSessionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
