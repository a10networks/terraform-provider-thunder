package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGenericProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_generic_proxy_oper`: Operational Status for the object generic-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbGenericProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"generic_proxy_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_http_conn_generic_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_selection_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_route_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"source_nat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"user_session": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"acr_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"acr_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aca_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aca_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dpr_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dpr_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dpa_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dpa_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cea_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cea_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cer_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cer_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwa_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwa_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwr_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwr_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"str_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"str_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sta_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sta_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asr_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asr_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asa_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asa_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mismatch_fwd_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mismatch_rev_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unkwn_cmd_code": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_session_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_fwd_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_rev_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcmsg_fwd_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcmsg_fwd_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcmsg_rev_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcmsg_rev_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcmsg_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry_client_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry_client_request_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reply_unknown_session_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ccr_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ccr_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cca_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cca_in": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ccr_i": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ccr_u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ccr_t": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cca_t": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"terminate_on_cca_t": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"forward_unknown_session_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_latest_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_select_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_conn_when_vport_down": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_avp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reselect_fwd_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reselect_fwd_tuple_other_cpu": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reselect_rev_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_closed_by_client": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_closed_by_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reply_invalid_avp_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reply_unable_to_deliver": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reply_error_info_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbGenericProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbGenericProxyOperOper := setObjectSlbGenericProxyOperOper(res)
		d.Set("oper", SlbGenericProxyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbGenericProxyOperOper(ret edpt.DataSlbGenericProxyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"generic_proxy_cpu_list": setSliceSlbGenericProxyOperOperGenericProxyCpuList(ret.DtSlbGenericProxyOper.Oper.GenericProxyCpuList),
			"cpu_count":              ret.DtSlbGenericProxyOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbGenericProxyOperOperGenericProxyCpuList(d []edpt.SlbGenericProxyOperOperGenericProxyCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy_conns"] = item.Curr_proxy_conns
		in["total_proxy_conns"] = item.Total_proxy_conns
		in["total_http_conn_generic_proxy"] = item.Total_http_conn_generic_proxy
		in["client_fail"] = item.Client_fail
		in["server_fail"] = item.Server_fail
		in["server_selection_fail"] = item.Server_selection_fail
		in["no_route_fail"] = item.No_route_fail
		in["source_nat_fail"] = item.Source_nat_fail
		in["user_session"] = item.User_session
		in["acr_out"] = item.Acr_out
		in["acr_in"] = item.Acr_in
		in["aca_out"] = item.Aca_out
		in["aca_in"] = item.Aca_in
		in["dpr_out"] = item.Dpr_out
		in["dpr_in"] = item.Dpr_in
		in["dpa_out"] = item.Dpa_out
		in["dpa_in"] = item.Dpa_in
		in["cea_out"] = item.Cea_out
		in["cea_in"] = item.Cea_in
		in["cer_out"] = item.Cer_out
		in["cer_in"] = item.Cer_in
		in["dwa_out"] = item.Dwa_out
		in["dwa_in"] = item.Dwa_in
		in["dwr_out"] = item.Dwr_out
		in["dwr_in"] = item.Dwr_in
		in["str_out"] = item.Str_out
		in["str_in"] = item.Str_in
		in["sta_out"] = item.Sta_out
		in["sta_in"] = item.Sta_in
		in["asr_out"] = item.Asr_out
		in["asr_in"] = item.Asr_in
		in["asa_out"] = item.Asa_out
		in["asa_in"] = item.Asa_in
		in["other_out"] = item.Other_out
		in["other_in"] = item.Other_in
		in["mismatch_fwd_id"] = item.Mismatch_fwd_id
		in["mismatch_rev_id"] = item.Mismatch_rev_id
		in["unkwn_cmd_code"] = item.Unkwn_cmd_code
		in["no_session_id"] = item.No_session_id
		in["no_fwd_tuple"] = item.No_fwd_tuple
		in["no_rev_tuple"] = item.No_rev_tuple
		in["dcmsg_fwd_in"] = item.Dcmsg_fwd_in
		in["dcmsg_fwd_out"] = item.Dcmsg_fwd_out
		in["dcmsg_rev_in"] = item.Dcmsg_rev_in
		in["dcmsg_rev_out"] = item.Dcmsg_rev_out
		in["dcmsg_error"] = item.Dcmsg_error
		in["retry_client_request"] = item.Retry_client_request
		in["retry_client_request_fail"] = item.Retry_client_request_fail
		in["reply_unknown_session_id"] = item.Reply_unknown_session_id
		in["ccr_out"] = item.Ccr_out
		in["ccr_in"] = item.Ccr_in
		in["cca_out"] = item.Cca_out
		in["cca_in"] = item.Cca_in
		in["ccr_i"] = item.Ccr_i
		in["ccr_u"] = item.Ccr_u
		in["ccr_t"] = item.Ccr_t
		in["cca_t"] = item.Cca_t
		in["terminate_on_cca_t"] = item.Terminate_on_cca_t
		in["forward_unknown_session_id"] = item.Forward_unknown_session_id
		in["update_latest_server"] = item.Update_latest_server
		in["client_select_fail"] = item.Client_select_fail
		in["close_conn_when_vport_down"] = item.Close_conn_when_vport_down
		in["invalid_avp"] = item.Invalid_avp
		in["reselect_fwd_tuple"] = item.Reselect_fwd_tuple
		in["reselect_fwd_tuple_other_cpu"] = item.Reselect_fwd_tuple_other_cpu
		in["reselect_rev_tuple"] = item.Reselect_rev_tuple
		in["conn_closed_by_client"] = item.Conn_closed_by_client
		in["conn_closed_by_server"] = item.Conn_closed_by_server
		in["reply_invalid_avp_value"] = item.Reply_invalid_avp_value
		in["reply_unable_to_deliver"] = item.Reply_unable_to_deliver
		in["reply_error_info_fail"] = item.Reply_error_info_fail
		result = append(result, in)
	}
	return result
}

func getObjectSlbGenericProxyOperOper(d []interface{}) edpt.SlbGenericProxyOperOper {

	count1 := len(d)
	var ret edpt.SlbGenericProxyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GenericProxyCpuList = getSliceSlbGenericProxyOperOperGenericProxyCpuList(in["generic_proxy_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbGenericProxyOperOperGenericProxyCpuList(d []interface{}) []edpt.SlbGenericProxyOperOperGenericProxyCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbGenericProxyOperOperGenericProxyCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbGenericProxyOperOperGenericProxyCpuList
		oi.Curr_proxy_conns = in["curr_proxy_conns"].(int)
		oi.Total_proxy_conns = in["total_proxy_conns"].(int)
		oi.Total_http_conn_generic_proxy = in["total_http_conn_generic_proxy"].(int)
		oi.Client_fail = in["client_fail"].(int)
		oi.Server_fail = in["server_fail"].(int)
		oi.Server_selection_fail = in["server_selection_fail"].(int)
		oi.No_route_fail = in["no_route_fail"].(int)
		oi.Source_nat_fail = in["source_nat_fail"].(int)
		oi.User_session = in["user_session"].(string)
		oi.Acr_out = in["acr_out"].(int)
		oi.Acr_in = in["acr_in"].(int)
		oi.Aca_out = in["aca_out"].(int)
		oi.Aca_in = in["aca_in"].(int)
		oi.Dpr_out = in["dpr_out"].(int)
		oi.Dpr_in = in["dpr_in"].(int)
		oi.Dpa_out = in["dpa_out"].(int)
		oi.Dpa_in = in["dpa_in"].(int)
		oi.Cea_out = in["cea_out"].(int)
		oi.Cea_in = in["cea_in"].(int)
		oi.Cer_out = in["cer_out"].(int)
		oi.Cer_in = in["cer_in"].(int)
		oi.Dwa_out = in["dwa_out"].(int)
		oi.Dwa_in = in["dwa_in"].(int)
		oi.Dwr_out = in["dwr_out"].(int)
		oi.Dwr_in = in["dwr_in"].(int)
		oi.Str_out = in["str_out"].(int)
		oi.Str_in = in["str_in"].(int)
		oi.Sta_out = in["sta_out"].(int)
		oi.Sta_in = in["sta_in"].(int)
		oi.Asr_out = in["asr_out"].(int)
		oi.Asr_in = in["asr_in"].(int)
		oi.Asa_out = in["asa_out"].(int)
		oi.Asa_in = in["asa_in"].(int)
		oi.Other_out = in["other_out"].(int)
		oi.Other_in = in["other_in"].(int)
		oi.Mismatch_fwd_id = in["mismatch_fwd_id"].(int)
		oi.Mismatch_rev_id = in["mismatch_rev_id"].(int)
		oi.Unkwn_cmd_code = in["unkwn_cmd_code"].(int)
		oi.No_session_id = in["no_session_id"].(int)
		oi.No_fwd_tuple = in["no_fwd_tuple"].(int)
		oi.No_rev_tuple = in["no_rev_tuple"].(int)
		oi.Dcmsg_fwd_in = in["dcmsg_fwd_in"].(int)
		oi.Dcmsg_fwd_out = in["dcmsg_fwd_out"].(int)
		oi.Dcmsg_rev_in = in["dcmsg_rev_in"].(int)
		oi.Dcmsg_rev_out = in["dcmsg_rev_out"].(int)
		oi.Dcmsg_error = in["dcmsg_error"].(int)
		oi.Retry_client_request = in["retry_client_request"].(int)
		oi.Retry_client_request_fail = in["retry_client_request_fail"].(int)
		oi.Reply_unknown_session_id = in["reply_unknown_session_id"].(int)
		oi.Ccr_out = in["ccr_out"].(int)
		oi.Ccr_in = in["ccr_in"].(int)
		oi.Cca_out = in["cca_out"].(int)
		oi.Cca_in = in["cca_in"].(int)
		oi.Ccr_i = in["ccr_i"].(int)
		oi.Ccr_u = in["ccr_u"].(int)
		oi.Ccr_t = in["ccr_t"].(int)
		oi.Cca_t = in["cca_t"].(int)
		oi.Terminate_on_cca_t = in["terminate_on_cca_t"].(int)
		oi.Forward_unknown_session_id = in["forward_unknown_session_id"].(int)
		oi.Update_latest_server = in["update_latest_server"].(int)
		oi.Client_select_fail = in["client_select_fail"].(int)
		oi.Close_conn_when_vport_down = in["close_conn_when_vport_down"].(int)
		oi.Invalid_avp = in["invalid_avp"].(int)
		oi.Reselect_fwd_tuple = in["reselect_fwd_tuple"].(int)
		oi.Reselect_fwd_tuple_other_cpu = in["reselect_fwd_tuple_other_cpu"].(int)
		oi.Reselect_rev_tuple = in["reselect_rev_tuple"].(int)
		oi.Conn_closed_by_client = in["conn_closed_by_client"].(int)
		oi.Conn_closed_by_server = in["conn_closed_by_server"].(int)
		oi.Reply_invalid_avp_value = in["reply_invalid_avp_value"].(int)
		oi.Reply_unable_to_deliver = in["reply_unable_to_deliver"].(int)
		oi.Reply_error_info_fail = in["reply_error_info_fail"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbGenericProxyOper(d *schema.ResourceData) edpt.SlbGenericProxyOper {
	var ret edpt.SlbGenericProxyOper

	ret.Oper = getObjectSlbGenericProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
