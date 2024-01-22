package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4UdpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_udp_stats`: Statistics for the object l4-udp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4UdpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Created",
						},
						"inudp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Received",
						},
						"instateless": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Stateless Packets",
						},
						"udp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Dropped",
						},
						"udp_drop_dst": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst Packets Dropped",
						},
						"udp_drop_src": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Packets Dropped",
						},
						"udp_drop_black_user_cfg_src": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Blacklist User Packets Dropped",
						},
						"udp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcDst Packets Dropped",
						},
						"udp_drop_black_user_cfg_src_dst": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcDst Blacklist User Packets Dropped",
						},
						"udp_port_zero_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port 0 Packets Dropped",
						},
						"udp_wellknown_src_port_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown Dropped",
						},
						"udp_retry_start": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Init",
						},
						"udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Passed",
						},
						"udp_retry_fail": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Failed",
						},
						"udp_retry_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Timeout",
						},
						"udp_payload_too_big_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large Dropped",
						},
						"udp_payload_too_small_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small Dropped",
						},
						"ntp_monlist_req_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Request Dropped",
						},
						"ntp_monlist_resp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Response Dropped",
						},
						"udp_conn_prate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Pkt Rate Dropped",
						},
						"dst_udp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Match",
						},
						"dst_udp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter No Match",
						},
						"dst_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Blacklist",
						},
						"dst_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Drop",
						},
						"dst_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Default Pass",
						},
						"dst_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action WL",
						},
						"udp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth Passed",
						},
						"src_udp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Match",
						},
						"src_udp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
						},
						"src_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
						},
						"src_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
						},
						"src_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
						},
						"src_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action WL",
						},
						"src_dst_udp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Match",
						},
						"src_dst_udp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter No Match",
						},
						"src_dst_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Blacklist",
						},
						"src_dst_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Drop",
						},
						"src_dst_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Default Pass",
						},
						"src_dst_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action WL",
						},
						"udp_wellknown_src_port": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown",
						},
						"udp_wellknown_src_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown Blacklisted",
						},
						"udp_retry_pass_wl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Pass WL",
						},
						"udp_retry_fail_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Fail Blacklisted",
						},
						"udp_payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large",
						},
						"udp_payload_too_big_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large Blacklisted",
						},
						"udp_payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small",
						},
						"udp_payload_too_small_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small Blacklisted",
						},
						"ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Request",
						},
						"ntp_monlist_req_bl": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Request Blacklisted",
						},
						"ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Response",
						},
						"ntp_monlist_resp_bl": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Response Blacklsited",
						},
						"udp_conn_prate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Pkt Rate Exceeded",
						},
						"udp_conn_prate_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Pkt Rate Blacklisted",
						},
						"udp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Exceeded",
						},
						"udp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Blacklist Packets Dropped",
						},
						"udp_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Frag Received",
						},
						"udp_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Frag Dropped",
						},
						"udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth Dropped",
						},
						"udp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Received",
						},
						"udp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Dropped",
						},
						"udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry-Gap Drop",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4UdpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4UdpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4UdpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4UdpStatsStats := setObjectDdosL4UdpStatsStats(res)
		d.Set("stats", DdosL4UdpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4UdpStatsStats(ret edpt.DataDdosL4UdpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"udp_sess_create":                        ret.DtDdosL4UdpStats.Stats.Udp_sess_create,
			"inudp":                                  ret.DtDdosL4UdpStats.Stats.Inudp,
			"instateless":                            ret.DtDdosL4UdpStats.Stats.Instateless,
			"udp_total_drop":                         ret.DtDdosL4UdpStats.Stats.Udp_total_drop,
			"udp_drop_dst":                           ret.DtDdosL4UdpStats.Stats.Udp_drop_dst,
			"udp_drop_src":                           ret.DtDdosL4UdpStats.Stats.Udp_drop_src,
			"udp_drop_black_user_cfg_src":            ret.DtDdosL4UdpStats.Stats.Udp_drop_black_user_cfg_src,
			"udp_src_dst_drop":                       ret.DtDdosL4UdpStats.Stats.Udp_src_dst_drop,
			"udp_drop_black_user_cfg_src_dst":        ret.DtDdosL4UdpStats.Stats.Udp_drop_black_user_cfg_src_dst,
			"udp_port_zero_drop":                     ret.DtDdosL4UdpStats.Stats.Udp_port_zero_drop,
			"udp_wellknown_src_port_drop":            ret.DtDdosL4UdpStats.Stats.Udp_wellknown_src_port_drop,
			"udp_retry_start":                        ret.DtDdosL4UdpStats.Stats.Udp_retry_start,
			"udp_retry_pass":                         ret.DtDdosL4UdpStats.Stats.Udp_retry_pass,
			"udp_retry_fail":                         ret.DtDdosL4UdpStats.Stats.Udp_retry_fail,
			"udp_retry_timeout":                      ret.DtDdosL4UdpStats.Stats.Udp_retry_timeout,
			"udp_payload_too_big_drop":               ret.DtDdosL4UdpStats.Stats.Udp_payload_too_big_drop,
			"udp_payload_too_small_drop":             ret.DtDdosL4UdpStats.Stats.Udp_payload_too_small_drop,
			"ntp_monlist_req_drop":                   ret.DtDdosL4UdpStats.Stats.Ntp_monlist_req_drop,
			"ntp_monlist_resp_drop":                  ret.DtDdosL4UdpStats.Stats.Ntp_monlist_resp_drop,
			"udp_conn_prate_drop":                    ret.DtDdosL4UdpStats.Stats.Udp_conn_prate_drop,
			"dst_udp_filter_match":                   ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_match,
			"dst_udp_filter_not_match":               ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_not_match,
			"dst_udp_filter_action_blacklist":        ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_action_blacklist,
			"dst_udp_filter_action_drop":             ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_action_drop,
			"dst_udp_filter_action_default_pass":     ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_action_default_pass,
			"dst_udp_filter_action_whitelist":        ret.DtDdosL4UdpStats.Stats.Dst_udp_filter_action_whitelist,
			"udp_auth_pass":                          ret.DtDdosL4UdpStats.Stats.Udp_auth_pass,
			"src_udp_filter_match":                   ret.DtDdosL4UdpStats.Stats.Src_udp_filter_match,
			"src_udp_filter_not_match":               ret.DtDdosL4UdpStats.Stats.Src_udp_filter_not_match,
			"src_udp_filter_action_blacklist":        ret.DtDdosL4UdpStats.Stats.Src_udp_filter_action_blacklist,
			"src_udp_filter_action_drop":             ret.DtDdosL4UdpStats.Stats.Src_udp_filter_action_drop,
			"src_udp_filter_action_default_pass":     ret.DtDdosL4UdpStats.Stats.Src_udp_filter_action_default_pass,
			"src_udp_filter_action_whitelist":        ret.DtDdosL4UdpStats.Stats.Src_udp_filter_action_whitelist,
			"src_dst_udp_filter_match":               ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_match,
			"src_dst_udp_filter_not_match":           ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_not_match,
			"src_dst_udp_filter_action_blacklist":    ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_action_blacklist,
			"src_dst_udp_filter_action_drop":         ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_action_drop,
			"src_dst_udp_filter_action_default_pass": ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_action_default_pass,
			"src_dst_udp_filter_action_whitelist":    ret.DtDdosL4UdpStats.Stats.Src_dst_udp_filter_action_whitelist,
			"udp_wellknown_src_port":                 ret.DtDdosL4UdpStats.Stats.Udp_wellknown_src_port,
			"udp_wellknown_src_port_bl":              ret.DtDdosL4UdpStats.Stats.Udp_wellknown_src_port_bl,
			"udp_retry_pass_wl":                      ret.DtDdosL4UdpStats.Stats.Udp_retry_pass_wl,
			"udp_retry_fail_bl":                      ret.DtDdosL4UdpStats.Stats.Udp_retry_fail_bl,
			"udp_payload_too_big":                    ret.DtDdosL4UdpStats.Stats.Udp_payload_too_big,
			"udp_payload_too_big_bl":                 ret.DtDdosL4UdpStats.Stats.Udp_payload_too_big_bl,
			"udp_payload_too_small":                  ret.DtDdosL4UdpStats.Stats.Udp_payload_too_small,
			"udp_payload_too_small_bl":               ret.DtDdosL4UdpStats.Stats.Udp_payload_too_small_bl,
			"ntp_monlist_req":                        ret.DtDdosL4UdpStats.Stats.Ntp_monlist_req,
			"ntp_monlist_req_bl":                     ret.DtDdosL4UdpStats.Stats.Ntp_monlist_req_bl,
			"ntp_monlist_resp":                       ret.DtDdosL4UdpStats.Stats.Ntp_monlist_resp,
			"ntp_monlist_resp_bl":                    ret.DtDdosL4UdpStats.Stats.Ntp_monlist_resp_bl,
			"udp_conn_prate_exceed":                  ret.DtDdosL4UdpStats.Stats.Udp_conn_prate_exceed,
			"udp_conn_prate_bl":                      ret.DtDdosL4UdpStats.Stats.Udp_conn_prate_bl,
			"udp_any_exceed":                         ret.DtDdosL4UdpStats.Stats.Udp_any_exceed,
			"udp_drop_bl":                            ret.DtDdosL4UdpStats.Stats.Udp_drop_bl,
			"udp_frag_rcvd":                          ret.DtDdosL4UdpStats.Stats.Udp_frag_rcvd,
			"udp_frag_drop":                          ret.DtDdosL4UdpStats.Stats.Udp_frag_drop,
			"udp_auth_drop":                          ret.DtDdosL4UdpStats.Stats.Udp_auth_drop,
			"udp_total_bytes_rcv":                    ret.DtDdosL4UdpStats.Stats.Udp_total_bytes_rcv,
			"udp_total_bytes_drop":                   ret.DtDdosL4UdpStats.Stats.Udp_total_bytes_drop,
			"udp_retry_gap_drop":                     ret.DtDdosL4UdpStats.Stats.Udp_retry_gap_drop,
		},
	}
}

func getObjectDdosL4UdpStatsStats(d []interface{}) edpt.DdosL4UdpStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4UdpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Udp_sess_create = in["udp_sess_create"].(int)
		ret.Inudp = in["inudp"].(int)
		ret.Instateless = in["instateless"].(int)
		ret.Udp_total_drop = in["udp_total_drop"].(int)
		ret.Udp_drop_dst = in["udp_drop_dst"].(int)
		ret.Udp_drop_src = in["udp_drop_src"].(int)
		ret.Udp_drop_black_user_cfg_src = in["udp_drop_black_user_cfg_src"].(int)
		ret.Udp_src_dst_drop = in["udp_src_dst_drop"].(int)
		ret.Udp_drop_black_user_cfg_src_dst = in["udp_drop_black_user_cfg_src_dst"].(int)
		ret.Udp_port_zero_drop = in["udp_port_zero_drop"].(int)
		ret.Udp_wellknown_src_port_drop = in["udp_wellknown_src_port_drop"].(int)
		ret.Udp_retry_start = in["udp_retry_start"].(int)
		ret.Udp_retry_pass = in["udp_retry_pass"].(int)
		ret.Udp_retry_fail = in["udp_retry_fail"].(int)
		ret.Udp_retry_timeout = in["udp_retry_timeout"].(int)
		ret.Udp_payload_too_big_drop = in["udp_payload_too_big_drop"].(int)
		ret.Udp_payload_too_small_drop = in["udp_payload_too_small_drop"].(int)
		ret.Ntp_monlist_req_drop = in["ntp_monlist_req_drop"].(int)
		ret.Ntp_monlist_resp_drop = in["ntp_monlist_resp_drop"].(int)
		ret.Udp_conn_prate_drop = in["udp_conn_prate_drop"].(int)
		ret.Dst_udp_filter_match = in["dst_udp_filter_match"].(int)
		ret.Dst_udp_filter_not_match = in["dst_udp_filter_not_match"].(int)
		ret.Dst_udp_filter_action_blacklist = in["dst_udp_filter_action_blacklist"].(int)
		ret.Dst_udp_filter_action_drop = in["dst_udp_filter_action_drop"].(int)
		ret.Dst_udp_filter_action_default_pass = in["dst_udp_filter_action_default_pass"].(int)
		ret.Dst_udp_filter_action_whitelist = in["dst_udp_filter_action_whitelist"].(int)
		ret.Udp_auth_pass = in["udp_auth_pass"].(int)
		ret.Src_udp_filter_match = in["src_udp_filter_match"].(int)
		ret.Src_udp_filter_not_match = in["src_udp_filter_not_match"].(int)
		ret.Src_udp_filter_action_blacklist = in["src_udp_filter_action_blacklist"].(int)
		ret.Src_udp_filter_action_drop = in["src_udp_filter_action_drop"].(int)
		ret.Src_udp_filter_action_default_pass = in["src_udp_filter_action_default_pass"].(int)
		ret.Src_udp_filter_action_whitelist = in["src_udp_filter_action_whitelist"].(int)
		ret.Src_dst_udp_filter_match = in["src_dst_udp_filter_match"].(int)
		ret.Src_dst_udp_filter_not_match = in["src_dst_udp_filter_not_match"].(int)
		ret.Src_dst_udp_filter_action_blacklist = in["src_dst_udp_filter_action_blacklist"].(int)
		ret.Src_dst_udp_filter_action_drop = in["src_dst_udp_filter_action_drop"].(int)
		ret.Src_dst_udp_filter_action_default_pass = in["src_dst_udp_filter_action_default_pass"].(int)
		ret.Src_dst_udp_filter_action_whitelist = in["src_dst_udp_filter_action_whitelist"].(int)
		ret.Udp_wellknown_src_port = in["udp_wellknown_src_port"].(int)
		ret.Udp_wellknown_src_port_bl = in["udp_wellknown_src_port_bl"].(int)
		ret.Udp_retry_pass_wl = in["udp_retry_pass_wl"].(int)
		ret.Udp_retry_fail_bl = in["udp_retry_fail_bl"].(int)
		ret.Udp_payload_too_big = in["udp_payload_too_big"].(int)
		ret.Udp_payload_too_big_bl = in["udp_payload_too_big_bl"].(int)
		ret.Udp_payload_too_small = in["udp_payload_too_small"].(int)
		ret.Udp_payload_too_small_bl = in["udp_payload_too_small_bl"].(int)
		ret.Ntp_monlist_req = in["ntp_monlist_req"].(int)
		ret.Ntp_monlist_req_bl = in["ntp_monlist_req_bl"].(int)
		ret.Ntp_monlist_resp = in["ntp_monlist_resp"].(int)
		ret.Ntp_monlist_resp_bl = in["ntp_monlist_resp_bl"].(int)
		ret.Udp_conn_prate_exceed = in["udp_conn_prate_exceed"].(int)
		ret.Udp_conn_prate_bl = in["udp_conn_prate_bl"].(int)
		ret.Udp_any_exceed = in["udp_any_exceed"].(int)
		ret.Udp_drop_bl = in["udp_drop_bl"].(int)
		ret.Udp_frag_rcvd = in["udp_frag_rcvd"].(int)
		ret.Udp_frag_drop = in["udp_frag_drop"].(int)
		ret.Udp_auth_drop = in["udp_auth_drop"].(int)
		ret.Udp_total_bytes_rcv = in["udp_total_bytes_rcv"].(int)
		ret.Udp_total_bytes_drop = in["udp_total_bytes_drop"].(int)
		ret.Udp_retry_gap_drop = in["udp_retry_gap_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosL4UdpStats(d *schema.ResourceData) edpt.DdosL4UdpStats {
	var ret edpt.DdosL4UdpStats

	ret.Stats = getObjectDdosL4UdpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
