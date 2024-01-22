package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGenericProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_generic_proxy_stats`: Statistics for the object generic-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbGenericProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num": {
							Type: schema.TypeInt, Optional: true, Description: "Number",
						},
						"curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "Total",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server selection failed",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Description: "Number of no routes",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of snat failures",
						},
						"client_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of client failures",
						},
						"server_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server failures",
						},
						"no_sess": {
							Type: schema.TypeInt, Optional: true, Description: "Number of no sessions",
						},
						"user_session": {
							Type: schema.TypeInt, Optional: true, Description: "Number of user sessions",
						},
						"acr_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ACRs out",
						},
						"acr_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ACRs in",
						},
						"aca_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ACAs out",
						},
						"aca_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ACAs in",
						},
						"cea_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CEAs out",
						},
						"cea_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CEAs in",
						},
						"cer_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CERs out",
						},
						"cer_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CERs in",
						},
						"dwr_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DWRs out",
						},
						"dwr_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DWRs in",
						},
						"dwa_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DWAs out",
						},
						"dwa_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DWAs in",
						},
						"str_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of STRs out",
						},
						"str_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of STRs in",
						},
						"sta_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of STAs out",
						},
						"sta_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of STAs in",
						},
						"asr_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ASRs out",
						},
						"asr_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ASRs in",
						},
						"asa_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ASAs out",
						},
						"asa_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ASAs in",
						},
						"other_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of other messages out",
						},
						"other_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of other messages in",
						},
						"total_http_req_enter_gen": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of HTTP requests enter generic proxy",
						},
						"mismatch_fwd_id": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter mismatch fwd session id",
						},
						"mismatch_rev_id": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter mismatch rev session id",
						},
						"unkwn_cmd_code": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter unkown cmd code",
						},
						"no_session_id": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter no session id avp",
						},
						"no_fwd_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter no fwd tuple matched",
						},
						"no_rev_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter no rev tuple matched",
						},
						"dcmsg_fwd_in": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter cross cpu fwd in",
						},
						"dcmsg_fwd_out": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter cross cpu fwd out",
						},
						"dcmsg_rev_in": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter cross cpu rev in",
						},
						"dcmsg_rev_out": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter cross cpu rev out",
						},
						"dcmsg_error": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter cross cpu error",
						},
						"retry_client_request": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter retry client request",
						},
						"retry_client_request_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter retry client request fail",
						},
						"reply_unknown_session_id": {
							Type: schema.TypeInt, Optional: true, Description: "Reply with unknown session ID error info",
						},
						"ccr_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCRs out",
						},
						"ccr_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCRs in",
						},
						"cca_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCAs out",
						},
						"cca_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCAs in",
						},
						"ccr_i": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCRs initial",
						},
						"ccr_u": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCRs update",
						},
						"ccr_t": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCRs terminate",
						},
						"cca_t": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CCAs terminate",
						},
						"terminate_on_cca_t": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter terminate on cca_t",
						},
						"forward_unknown_session_id": {
							Type: schema.TypeInt, Optional: true, Description: "Forward server side message with unknown session id",
						},
						"update_latest_server": {
							Type: schema.TypeInt, Optional: true, Description: "Update to the latest server that used a session id",
						},
						"client_select_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fail to select client",
						},
						"close_conn_when_vport_down": {
							Type: schema.TypeInt, Optional: true, Description: "Close client conn when virtual port is down",
						},
						"invalid_avp": {
							Type: schema.TypeInt, Optional: true, Description: "AVP value contains illegal chars",
						},
						"reselect_fwd_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Original client tuple does not exist so reselect another one",
						},
						"reselect_fwd_tuple_other_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "Original client tuple does not exist so reselect another one on other CPUs",
						},
						"reselect_rev_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Original server tuple does not exist so reselect another one",
						},
						"conn_closed_by_client": {
							Type: schema.TypeInt, Optional: true, Description: "Client initiates TCP close/reset",
						},
						"conn_closed_by_server": {
							Type: schema.TypeInt, Optional: true, Description: "Server initiates TCP close/reset",
						},
						"reply_invalid_avp_value": {
							Type: schema.TypeInt, Optional: true, Description: "Reply with invalid AVP error info",
						},
						"reply_unable_to_deliver": {
							Type: schema.TypeInt, Optional: true, Description: "Reply with unable to deliver error info",
						},
						"reply_error_info_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fail to reply error info to peer",
						},
						"dpr_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DPRs out",
						},
						"dpr_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DPRs in",
						},
						"dpa_out": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DPAs out",
						},
						"dpa_in": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DPAs in",
						},
					},
				},
			},
		},
	}
}

func resourceSlbGenericProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbGenericProxyStatsStats := setObjectSlbGenericProxyStatsStats(res)
		d.Set("stats", SlbGenericProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbGenericProxyStatsStats(ret edpt.DataSlbGenericProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num":                          ret.DtSlbGenericProxyStats.Stats.Num,
			"curr":                         ret.DtSlbGenericProxyStats.Stats.Curr,
			"total":                        ret.DtSlbGenericProxyStats.Stats.Total,
			"svrsel_fail":                  ret.DtSlbGenericProxyStats.Stats.Svrsel_fail,
			"no_route":                     ret.DtSlbGenericProxyStats.Stats.No_route,
			"snat_fail":                    ret.DtSlbGenericProxyStats.Stats.Snat_fail,
			"client_fail":                  ret.DtSlbGenericProxyStats.Stats.Client_fail,
			"server_fail":                  ret.DtSlbGenericProxyStats.Stats.Server_fail,
			"no_sess":                      ret.DtSlbGenericProxyStats.Stats.No_sess,
			"user_session":                 ret.DtSlbGenericProxyStats.Stats.User_session,
			"acr_out":                      ret.DtSlbGenericProxyStats.Stats.Acr_out,
			"acr_in":                       ret.DtSlbGenericProxyStats.Stats.Acr_in,
			"aca_out":                      ret.DtSlbGenericProxyStats.Stats.Aca_out,
			"aca_in":                       ret.DtSlbGenericProxyStats.Stats.Aca_in,
			"cea_out":                      ret.DtSlbGenericProxyStats.Stats.Cea_out,
			"cea_in":                       ret.DtSlbGenericProxyStats.Stats.Cea_in,
			"cer_out":                      ret.DtSlbGenericProxyStats.Stats.Cer_out,
			"cer_in":                       ret.DtSlbGenericProxyStats.Stats.Cer_in,
			"dwr_out":                      ret.DtSlbGenericProxyStats.Stats.Dwr_out,
			"dwr_in":                       ret.DtSlbGenericProxyStats.Stats.Dwr_in,
			"dwa_out":                      ret.DtSlbGenericProxyStats.Stats.Dwa_out,
			"dwa_in":                       ret.DtSlbGenericProxyStats.Stats.Dwa_in,
			"str_out":                      ret.DtSlbGenericProxyStats.Stats.Str_out,
			"str_in":                       ret.DtSlbGenericProxyStats.Stats.Str_in,
			"sta_out":                      ret.DtSlbGenericProxyStats.Stats.Sta_out,
			"sta_in":                       ret.DtSlbGenericProxyStats.Stats.Sta_in,
			"asr_out":                      ret.DtSlbGenericProxyStats.Stats.Asr_out,
			"asr_in":                       ret.DtSlbGenericProxyStats.Stats.Asr_in,
			"asa_out":                      ret.DtSlbGenericProxyStats.Stats.Asa_out,
			"asa_in":                       ret.DtSlbGenericProxyStats.Stats.Asa_in,
			"other_out":                    ret.DtSlbGenericProxyStats.Stats.Other_out,
			"other_in":                     ret.DtSlbGenericProxyStats.Stats.Other_in,
			"total_http_req_enter_gen":     ret.DtSlbGenericProxyStats.Stats.Total_http_req_enter_gen,
			"mismatch_fwd_id":              ret.DtSlbGenericProxyStats.Stats.Mismatch_fwd_id,
			"mismatch_rev_id":              ret.DtSlbGenericProxyStats.Stats.Mismatch_rev_id,
			"unkwn_cmd_code":               ret.DtSlbGenericProxyStats.Stats.Unkwn_cmd_code,
			"no_session_id":                ret.DtSlbGenericProxyStats.Stats.No_session_id,
			"no_fwd_tuple":                 ret.DtSlbGenericProxyStats.Stats.No_fwd_tuple,
			"no_rev_tuple":                 ret.DtSlbGenericProxyStats.Stats.No_rev_tuple,
			"dcmsg_fwd_in":                 ret.DtSlbGenericProxyStats.Stats.Dcmsg_fwd_in,
			"dcmsg_fwd_out":                ret.DtSlbGenericProxyStats.Stats.Dcmsg_fwd_out,
			"dcmsg_rev_in":                 ret.DtSlbGenericProxyStats.Stats.Dcmsg_rev_in,
			"dcmsg_rev_out":                ret.DtSlbGenericProxyStats.Stats.Dcmsg_rev_out,
			"dcmsg_error":                  ret.DtSlbGenericProxyStats.Stats.Dcmsg_error,
			"retry_client_request":         ret.DtSlbGenericProxyStats.Stats.Retry_client_request,
			"retry_client_request_fail":    ret.DtSlbGenericProxyStats.Stats.Retry_client_request_fail,
			"reply_unknown_session_id":     ret.DtSlbGenericProxyStats.Stats.Reply_unknown_session_id,
			"ccr_out":                      ret.DtSlbGenericProxyStats.Stats.Ccr_out,
			"ccr_in":                       ret.DtSlbGenericProxyStats.Stats.Ccr_in,
			"cca_out":                      ret.DtSlbGenericProxyStats.Stats.Cca_out,
			"cca_in":                       ret.DtSlbGenericProxyStats.Stats.Cca_in,
			"ccr_i":                        ret.DtSlbGenericProxyStats.Stats.Ccr_i,
			"ccr_u":                        ret.DtSlbGenericProxyStats.Stats.Ccr_u,
			"ccr_t":                        ret.DtSlbGenericProxyStats.Stats.Ccr_t,
			"cca_t":                        ret.DtSlbGenericProxyStats.Stats.Cca_t,
			"terminate_on_cca_t":           ret.DtSlbGenericProxyStats.Stats.Terminate_on_cca_t,
			"forward_unknown_session_id":   ret.DtSlbGenericProxyStats.Stats.Forward_unknown_session_id,
			"update_latest_server":         ret.DtSlbGenericProxyStats.Stats.Update_latest_server,
			"client_select_fail":           ret.DtSlbGenericProxyStats.Stats.Client_select_fail,
			"close_conn_when_vport_down":   ret.DtSlbGenericProxyStats.Stats.Close_conn_when_vport_down,
			"invalid_avp":                  ret.DtSlbGenericProxyStats.Stats.Invalid_avp,
			"reselect_fwd_tuple":           ret.DtSlbGenericProxyStats.Stats.Reselect_fwd_tuple,
			"reselect_fwd_tuple_other_cpu": ret.DtSlbGenericProxyStats.Stats.Reselect_fwd_tuple_other_cpu,
			"reselect_rev_tuple":           ret.DtSlbGenericProxyStats.Stats.Reselect_rev_tuple,
			"conn_closed_by_client":        ret.DtSlbGenericProxyStats.Stats.Conn_closed_by_client,
			"conn_closed_by_server":        ret.DtSlbGenericProxyStats.Stats.Conn_closed_by_server,
			"reply_invalid_avp_value":      ret.DtSlbGenericProxyStats.Stats.Reply_invalid_avp_value,
			"reply_unable_to_deliver":      ret.DtSlbGenericProxyStats.Stats.Reply_unable_to_deliver,
			"reply_error_info_fail":        ret.DtSlbGenericProxyStats.Stats.Reply_error_info_fail,
			"dpr_out":                      ret.DtSlbGenericProxyStats.Stats.Dpr_out,
			"dpr_in":                       ret.DtSlbGenericProxyStats.Stats.Dpr_in,
			"dpa_out":                      ret.DtSlbGenericProxyStats.Stats.Dpa_out,
			"dpa_in":                       ret.DtSlbGenericProxyStats.Stats.Dpa_in,
		},
	}
}

func getObjectSlbGenericProxyStatsStats(d []interface{}) edpt.SlbGenericProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbGenericProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_fail = in["client_fail"].(int)
		ret.Server_fail = in["server_fail"].(int)
		ret.No_sess = in["no_sess"].(int)
		ret.User_session = in["user_session"].(int)
		ret.Acr_out = in["acr_out"].(int)
		ret.Acr_in = in["acr_in"].(int)
		ret.Aca_out = in["aca_out"].(int)
		ret.Aca_in = in["aca_in"].(int)
		ret.Cea_out = in["cea_out"].(int)
		ret.Cea_in = in["cea_in"].(int)
		ret.Cer_out = in["cer_out"].(int)
		ret.Cer_in = in["cer_in"].(int)
		ret.Dwr_out = in["dwr_out"].(int)
		ret.Dwr_in = in["dwr_in"].(int)
		ret.Dwa_out = in["dwa_out"].(int)
		ret.Dwa_in = in["dwa_in"].(int)
		ret.Str_out = in["str_out"].(int)
		ret.Str_in = in["str_in"].(int)
		ret.Sta_out = in["sta_out"].(int)
		ret.Sta_in = in["sta_in"].(int)
		ret.Asr_out = in["asr_out"].(int)
		ret.Asr_in = in["asr_in"].(int)
		ret.Asa_out = in["asa_out"].(int)
		ret.Asa_in = in["asa_in"].(int)
		ret.Other_out = in["other_out"].(int)
		ret.Other_in = in["other_in"].(int)
		ret.Total_http_req_enter_gen = in["total_http_req_enter_gen"].(int)
		ret.Mismatch_fwd_id = in["mismatch_fwd_id"].(int)
		ret.Mismatch_rev_id = in["mismatch_rev_id"].(int)
		ret.Unkwn_cmd_code = in["unkwn_cmd_code"].(int)
		ret.No_session_id = in["no_session_id"].(int)
		ret.No_fwd_tuple = in["no_fwd_tuple"].(int)
		ret.No_rev_tuple = in["no_rev_tuple"].(int)
		ret.Dcmsg_fwd_in = in["dcmsg_fwd_in"].(int)
		ret.Dcmsg_fwd_out = in["dcmsg_fwd_out"].(int)
		ret.Dcmsg_rev_in = in["dcmsg_rev_in"].(int)
		ret.Dcmsg_rev_out = in["dcmsg_rev_out"].(int)
		ret.Dcmsg_error = in["dcmsg_error"].(int)
		ret.Retry_client_request = in["retry_client_request"].(int)
		ret.Retry_client_request_fail = in["retry_client_request_fail"].(int)
		ret.Reply_unknown_session_id = in["reply_unknown_session_id"].(int)
		ret.Ccr_out = in["ccr_out"].(int)
		ret.Ccr_in = in["ccr_in"].(int)
		ret.Cca_out = in["cca_out"].(int)
		ret.Cca_in = in["cca_in"].(int)
		ret.Ccr_i = in["ccr_i"].(int)
		ret.Ccr_u = in["ccr_u"].(int)
		ret.Ccr_t = in["ccr_t"].(int)
		ret.Cca_t = in["cca_t"].(int)
		ret.Terminate_on_cca_t = in["terminate_on_cca_t"].(int)
		ret.Forward_unknown_session_id = in["forward_unknown_session_id"].(int)
		ret.Update_latest_server = in["update_latest_server"].(int)
		ret.Client_select_fail = in["client_select_fail"].(int)
		ret.Close_conn_when_vport_down = in["close_conn_when_vport_down"].(int)
		ret.Invalid_avp = in["invalid_avp"].(int)
		ret.Reselect_fwd_tuple = in["reselect_fwd_tuple"].(int)
		ret.Reselect_fwd_tuple_other_cpu = in["reselect_fwd_tuple_other_cpu"].(int)
		ret.Reselect_rev_tuple = in["reselect_rev_tuple"].(int)
		ret.Conn_closed_by_client = in["conn_closed_by_client"].(int)
		ret.Conn_closed_by_server = in["conn_closed_by_server"].(int)
		ret.Reply_invalid_avp_value = in["reply_invalid_avp_value"].(int)
		ret.Reply_unable_to_deliver = in["reply_unable_to_deliver"].(int)
		ret.Reply_error_info_fail = in["reply_error_info_fail"].(int)
		ret.Dpr_out = in["dpr_out"].(int)
		ret.Dpr_in = in["dpr_in"].(int)
		ret.Dpa_out = in["dpa_out"].(int)
		ret.Dpa_in = in["dpa_in"].(int)
	}
	return ret
}

func dataToEndpointSlbGenericProxyStats(d *schema.ResourceData) edpt.SlbGenericProxyStats {
	var ret edpt.SlbGenericProxyStats

	ret.Stats = getObjectSlbGenericProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
