package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSipOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_sip_oper`: Operational Status for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSipOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"msg_proxy_current": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_cached": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_send_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_incomplete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_connection": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_parse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_process": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_exceed_tmp_buff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_send_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_incomplete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_parse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_process": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_selec_connt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_exceed_tmp_buff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_create_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_conn_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_construct_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_reserve_pconn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_start_server_conn_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_conn_already_exists": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_insert_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_parse_msg_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_process_msg_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_no_vport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_select_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_alloc_mem": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_unexpected_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l7_cpu_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_to_l7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_from_l7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_to_l4_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_from_l4_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l7_to_l4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mag_back": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_dcmsg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_deprecated_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_hold_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_split_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_pipline_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_created": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_in_rml": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_invalid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_callid_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_callid_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_mem_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"table_mem_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"table_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cmsg_no_uri_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cmsg_no_uri_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sg_no_uri_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smsg_no_uri_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_read_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_parse_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_unknown_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_unknown_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_parse_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_name_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"body_too_big": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_get_counter": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_no_call_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"identify_dir_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_sip_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"deprecated_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_insert_callid_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_insert_uri_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_insert_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn_by_callid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn_by_uri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn_by_rev_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"x_forward_for_select_client": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"call_id_select_client": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"uri_select_client": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_select_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"acl_denied": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"assemble_frag_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"wrong_ip_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"size_too_large": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_split_fragment": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_keepalive_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_keepalive_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_keepalive_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_keepalive_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ax_health_check_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_request_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"concatenate_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"save_uri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"save_uri_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"save_call_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"save_call_id_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_translation": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_translation_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_trans_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_trans_start_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_trans_body": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_register": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_invite": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_cancel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_bye": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_options": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_prack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_subscribe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_notify": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_publish": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_info": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_refer": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_message": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_update": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_send_sip_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_send_sip_session_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_fail_get_msg_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_recv_sip_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_insert_sip_session_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_update_sip_session_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_invalid_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_fail_alloc_sip_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_fail_alloc_call_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_fail_clone_sip_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"save_smp_call_id_rtp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_smp_call_id_rtp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smp_call_id_rtp_session_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smp_call_id_rtp_session_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"process_error_when_message_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSipOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSipOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSipOperOper := setObjectSlbSipOperOper(res)
		d.Set("oper", SlbSipOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSipOperOper(ret edpt.DataSlbSipOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sip_cpu_list": setSliceSlbSipOperOperSipCpuList(ret.DtSlbSipOper.Oper.SipCpuList),
			"cpu_count":    ret.DtSlbSipOper.Oper.CpuCount,
			"filter_type":  ret.DtSlbSipOper.Oper.Filter_type,
		},
	}
}

func setSliceSlbSipOperOperSipCpuList(d []edpt.SlbSipOperOperSipCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["msg_proxy_current"] = item.Msg_proxy_current
		in["msg_proxy_total"] = item.Msg_proxy_total
		in["msg_proxy_mem_allocd"] = item.Msg_proxy_mem_allocd
		in["msg_proxy_mem_cached"] = item.Msg_proxy_mem_cached
		in["msg_proxy_mem_freed"] = item.Msg_proxy_mem_freed
		in["msg_proxy_client_recv"] = item.Msg_proxy_client_recv
		in["msg_proxy_client_send_success"] = item.Msg_proxy_client_send_success
		in["msg_proxy_client_incomplete"] = item.Msg_proxy_client_incomplete
		in["msg_proxy_client_drop"] = item.Msg_proxy_client_drop
		in["msg_proxy_client_connection"] = item.Msg_proxy_client_connection
		in["msg_proxy_client_fail"] = item.Msg_proxy_client_fail
		in["msg_proxy_client_fail_parse"] = item.Msg_proxy_client_fail_parse
		in["msg_proxy_client_fail_process"] = item.Msg_proxy_client_fail_process
		in["msg_proxy_client_fail_snat"] = item.Msg_proxy_client_fail_snat
		in["msg_proxy_client_exceed_tmp_buff"] = item.Msg_proxy_client_exceed_tmp_buff
		in["msg_proxy_client_fail_send_pkt"] = item.Msg_proxy_client_fail_send_pkt
		in["msg_proxy_client_fail_start_server_conn"] = item.Msg_proxy_client_fail_start_server_conn
		in["msg_proxy_server_recv"] = item.Msg_proxy_server_recv
		in["msg_proxy_server_send_success"] = item.Msg_proxy_server_send_success
		in["msg_proxy_server_incomplete"] = item.Msg_proxy_server_incomplete
		in["msg_proxy_server_drop"] = item.Msg_proxy_server_drop
		in["msg_proxy_server_fail"] = item.Msg_proxy_server_fail
		in["msg_proxy_server_fail_parse"] = item.Msg_proxy_server_fail_parse
		in["msg_proxy_server_fail_process"] = item.Msg_proxy_server_fail_process
		in["msg_proxy_server_fail_selec_connt"] = item.Msg_proxy_server_fail_selec_connt
		in["msg_proxy_server_fail_snat"] = item.Msg_proxy_server_fail_snat
		in["msg_proxy_server_exceed_tmp_buff"] = item.Msg_proxy_server_exceed_tmp_buff
		in["msg_proxy_server_fail_send_pkt"] = item.Msg_proxy_server_fail_send_pkt
		in["msg_proxy_create_server_conn"] = item.Msg_proxy_create_server_conn
		in["msg_proxy_start_server_conn"] = item.Msg_proxy_start_server_conn
		in["msg_proxy_fail_start_server_conn"] = item.Msg_proxy_fail_start_server_conn
		in["msg_proxy_server_conn_fail_snat"] = item.Msg_proxy_server_conn_fail_snat
		in["msg_proxy_fail_construct_server_conn"] = item.Msg_proxy_fail_construct_server_conn
		in["msg_proxy_fail_reserve_pconn"] = item.Msg_proxy_fail_reserve_pconn
		in["msg_proxy_start_server_conn_failed"] = item.Msg_proxy_start_server_conn_failed
		in["msg_proxy_server_conn_already_exists"] = item.Msg_proxy_server_conn_already_exists
		in["msg_proxy_fail_insert_server_conn"] = item.Msg_proxy_fail_insert_server_conn
		in["msg_proxy_parse_msg_fail"] = item.Msg_proxy_parse_msg_fail
		in["msg_proxy_process_msg_fail"] = item.Msg_proxy_process_msg_fail
		in["msg_proxy_no_vport"] = item.Msg_proxy_no_vport
		in["msg_proxy_fail_select_server"] = item.Msg_proxy_fail_select_server
		in["msg_proxy_fail_alloc_mem"] = item.Msg_proxy_fail_alloc_mem
		in["msg_proxy_unexpected_err"] = item.Msg_proxy_unexpected_err
		in["msg_proxy_l7_cpu_failed"] = item.Msg_proxy_l7_cpu_failed
		in["msg_proxy_l4_to_l7"] = item.Msg_proxy_l4_to_l7
		in["msg_proxy_l4_from_l7"] = item.Msg_proxy_l4_from_l7
		in["msg_proxy_to_l4_send_pkt"] = item.Msg_proxy_to_l4_send_pkt
		in["msg_proxy_l4_from_l4_send"] = item.Msg_proxy_l4_from_l4_send
		in["msg_proxy_l7_to_l4"] = item.Msg_proxy_l7_to_l4
		in["msg_proxy_mag_back"] = item.Msg_proxy_mag_back
		in["msg_proxy_fail_dcmsg"] = item.Msg_proxy_fail_dcmsg
		in["msg_proxy_deprecated_conn"] = item.Msg_proxy_deprecated_conn
		in["msg_proxy_hold_msg"] = item.Msg_proxy_hold_msg
		in["msg_proxy_split_pkt"] = item.Msg_proxy_split_pkt
		in["msg_proxy_pipline_msg"] = item.Msg_proxy_pipline_msg
		in["msg_proxy_client_reset"] = item.Msg_proxy_client_reset
		in["msg_proxy_server_reset"] = item.Msg_proxy_server_reset
		in["session_created"] = item.Session_created
		in["session_freed"] = item.Session_freed
		in["session_in_rml"] = item.Session_in_rml
		in["session_invalid"] = item.Session_invalid
		in["conn_allocd"] = item.Conn_allocd
		in["conn_freed"] = item.Conn_freed
		in["session_callid_allocd"] = item.Session_callid_allocd
		in["session_callid_freed"] = item.Session_callid_freed
		in["line_mem_allocd"] = item.Line_mem_allocd
		in["line_mem_freed"] = item.Line_mem_freed
		in["table_mem_allocd"] = item.Table_mem_allocd
		in["table_mem_freed"] = item.Table_mem_freed
		in["cmsg_no_uri_header"] = item.Cmsg_no_uri_header
		in["cmsg_no_uri_session"] = item.Cmsg_no_uri_session
		in["sg_no_uri_header"] = item.Sg_no_uri_header
		in["smsg_no_uri_session"] = item.Smsg_no_uri_session
		in["line_too_long"] = item.Line_too_long
		in["fail_read_start_line"] = item.Fail_read_start_line
		in["fail_parse_start_line"] = item.Fail_parse_start_line
		in["invalid_start_line"] = item.Invalid_start_line
		in["request_unknown_version"] = item.Request_unknown_version
		in["response_unknown_version"] = item.Response_unknown_version
		in["request_unknown"] = item.Request_unknown
		in["fail_parse_headers"] = item.Fail_parse_headers
		in["too_many_headers"] = item.Too_many_headers
		in["invalid_header"] = item.Invalid_header
		in["header_name_too_long"] = item.Header_name_too_long
		in["body_too_big"] = item.Body_too_big
		in["fail_get_counter"] = item.Fail_get_counter
		in["msg_no_call_id"] = item.Msg_no_call_id
		in["identify_dir_failed"] = item.Identify_dir_failed
		in["no_sip_request"] = item.No_sip_request
		in["deprecated_msg"] = item.Deprecated_msg
		in["fail_insert_callid_session"] = item.Fail_insert_callid_session
		in["fail_insert_uri_session"] = item.Fail_insert_uri_session
		in["fail_insert_header"] = item.Fail_insert_header
		in["select_server_conn"] = item.Select_server_conn
		in["select_server_conn_by_callid"] = item.Select_server_conn_by_callid
		in["select_server_conn_by_uri"] = item.Select_server_conn_by_uri
		in["select_server_conn_by_rev_tuple"] = item.Select_server_conn_by_rev_tuple
		in["select_server_conn_failed"] = item.Select_server_conn_failed
		in["select_client_conn"] = item.Select_client_conn
		in["x_forward_for_select_client"] = item.X_forward_for_select_client
		in["call_id_select_client"] = item.Call_id_select_client
		in["uri_select_client"] = item.Uri_select_client
		in["client_select_failed"] = item.Client_select_failed
		in["acl_denied"] = item.Acl_denied
		in["assemble_frag_failed"] = item.Assemble_frag_failed
		in["wrong_ip_version"] = item.Wrong_ip_version
		in["size_too_large"] = item.Size_too_large
		in["fail_split_fragment"] = item.Fail_split_fragment
		in["client_keepalive_received"] = item.Client_keepalive_received
		in["server_keepalive_received"] = item.Server_keepalive_received
		in["client_keepalive_send"] = item.Client_keepalive_send
		in["server_keepalive_send"] = item.Server_keepalive_send
		in["ax_health_check_received"] = item.Ax_health_check_received
		in["client_request"] = item.Client_request
		in["client_request_ok"] = item.Client_request_ok
		in["concatenate_msg"] = item.Concatenate_msg
		in["save_uri"] = item.Save_uri
		in["save_uri_ok"] = item.Save_uri_ok
		in["save_call_id"] = item.Save_call_id
		in["save_call_id_ok"] = item.Save_call_id_ok
		in["msg_translation"] = item.Msg_translation
		in["msg_translation_fail"] = item.Msg_translation_fail
		in["msg_trans_start_line"] = item.Msg_trans_start_line
		in["msg_trans_start_headers"] = item.Msg_trans_start_headers
		in["msg_trans_body"] = item.Msg_trans_body
		in["request_register"] = item.Request_register
		in["request_invite"] = item.Request_invite
		in["request_ack"] = item.Request_ack
		in["request_cancel"] = item.Request_cancel
		in["request_bye"] = item.Request_bye
		in["request_options"] = item.Request_options
		in["request_prack"] = item.Request_prack
		in["request_subscribe"] = item.Request_subscribe
		in["request_notify"] = item.Request_notify
		in["request_publish"] = item.Request_publish
		in["request_info"] = item.Request_info
		in["request_refer"] = item.Request_refer
		in["request_message"] = item.Request_message
		in["request_update"] = item.Request_update
		in["response_unknown"] = item.Response_unknown
		in["response_1xx"] = item.Response_1xx
		in["response_2xx"] = item.Response_2xx
		in["response_3xx"] = item.Response_3xx
		in["response_4xx"] = item.Response_4xx
		in["response_5xx"] = item.Response_5xx
		in["response_6xx"] = item.Response_6xx
		in["ha_send_sip_session"] = item.Ha_send_sip_session
		in["ha_send_sip_session_ok"] = item.Ha_send_sip_session_ok
		in["ha_fail_get_msg_header"] = item.Ha_fail_get_msg_header
		in["ha_recv_sip_session"] = item.Ha_recv_sip_session
		in["ha_insert_sip_session_ok"] = item.Ha_insert_sip_session_ok
		in["ha_update_sip_session_ok"] = item.Ha_update_sip_session_ok
		in["ha_invalid_pkt"] = item.Ha_invalid_pkt
		in["ha_fail_alloc_sip_session"] = item.Ha_fail_alloc_sip_session
		in["ha_fail_alloc_call_id"] = item.Ha_fail_alloc_call_id
		in["ha_fail_clone_sip_session"] = item.Ha_fail_clone_sip_session
		in["save_smp_call_id_rtp"] = item.Save_smp_call_id_rtp
		in["update_smp_call_id_rtp"] = item.Update_smp_call_id_rtp
		in["smp_call_id_rtp_session_match"] = item.Smp_call_id_rtp_session_match
		in["smp_call_id_rtp_session_not_match"] = item.Smp_call_id_rtp_session_not_match
		in["process_error_when_message_switch"] = item.Process_error_when_message_switch
		result = append(result, in)
	}
	return result
}

func getObjectSlbSipOperOper(d []interface{}) edpt.SlbSipOperOper {

	count1 := len(d)
	var ret edpt.SlbSipOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipCpuList = getSliceSlbSipOperOperSipCpuList(in["sip_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.Filter_type = in["filter_type"].(string)
	}
	return ret
}

func getSliceSlbSipOperOperSipCpuList(d []interface{}) []edpt.SlbSipOperOperSipCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbSipOperOperSipCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSipOperOperSipCpuList
		oi.Msg_proxy_current = in["msg_proxy_current"].(int)
		oi.Msg_proxy_total = in["msg_proxy_total"].(int)
		oi.Msg_proxy_mem_allocd = in["msg_proxy_mem_allocd"].(int)
		oi.Msg_proxy_mem_cached = in["msg_proxy_mem_cached"].(int)
		oi.Msg_proxy_mem_freed = in["msg_proxy_mem_freed"].(int)
		oi.Msg_proxy_client_recv = in["msg_proxy_client_recv"].(int)
		oi.Msg_proxy_client_send_success = in["msg_proxy_client_send_success"].(int)
		oi.Msg_proxy_client_incomplete = in["msg_proxy_client_incomplete"].(int)
		oi.Msg_proxy_client_drop = in["msg_proxy_client_drop"].(int)
		oi.Msg_proxy_client_connection = in["msg_proxy_client_connection"].(int)
		oi.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		oi.Msg_proxy_client_fail_parse = in["msg_proxy_client_fail_parse"].(int)
		oi.Msg_proxy_client_fail_process = in["msg_proxy_client_fail_process"].(int)
		oi.Msg_proxy_client_fail_snat = in["msg_proxy_client_fail_snat"].(int)
		oi.Msg_proxy_client_exceed_tmp_buff = in["msg_proxy_client_exceed_tmp_buff"].(int)
		oi.Msg_proxy_client_fail_send_pkt = in["msg_proxy_client_fail_send_pkt"].(int)
		oi.Msg_proxy_client_fail_start_server_conn = in["msg_proxy_client_fail_start_server_conn"].(int)
		oi.Msg_proxy_server_recv = in["msg_proxy_server_recv"].(int)
		oi.Msg_proxy_server_send_success = in["msg_proxy_server_send_success"].(int)
		oi.Msg_proxy_server_incomplete = in["msg_proxy_server_incomplete"].(int)
		oi.Msg_proxy_server_drop = in["msg_proxy_server_drop"].(int)
		oi.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		oi.Msg_proxy_server_fail_parse = in["msg_proxy_server_fail_parse"].(int)
		oi.Msg_proxy_server_fail_process = in["msg_proxy_server_fail_process"].(int)
		oi.Msg_proxy_server_fail_selec_connt = in["msg_proxy_server_fail_selec_connt"].(int)
		oi.Msg_proxy_server_fail_snat = in["msg_proxy_server_fail_snat"].(int)
		oi.Msg_proxy_server_exceed_tmp_buff = in["msg_proxy_server_exceed_tmp_buff"].(int)
		oi.Msg_proxy_server_fail_send_pkt = in["msg_proxy_server_fail_send_pkt"].(int)
		oi.Msg_proxy_create_server_conn = in["msg_proxy_create_server_conn"].(int)
		oi.Msg_proxy_start_server_conn = in["msg_proxy_start_server_conn"].(int)
		oi.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		oi.Msg_proxy_server_conn_fail_snat = in["msg_proxy_server_conn_fail_snat"].(int)
		oi.Msg_proxy_fail_construct_server_conn = in["msg_proxy_fail_construct_server_conn"].(int)
		oi.Msg_proxy_fail_reserve_pconn = in["msg_proxy_fail_reserve_pconn"].(int)
		oi.Msg_proxy_start_server_conn_failed = in["msg_proxy_start_server_conn_failed"].(int)
		oi.Msg_proxy_server_conn_already_exists = in["msg_proxy_server_conn_already_exists"].(int)
		oi.Msg_proxy_fail_insert_server_conn = in["msg_proxy_fail_insert_server_conn"].(int)
		oi.Msg_proxy_parse_msg_fail = in["msg_proxy_parse_msg_fail"].(int)
		oi.Msg_proxy_process_msg_fail = in["msg_proxy_process_msg_fail"].(int)
		oi.Msg_proxy_no_vport = in["msg_proxy_no_vport"].(int)
		oi.Msg_proxy_fail_select_server = in["msg_proxy_fail_select_server"].(int)
		oi.Msg_proxy_fail_alloc_mem = in["msg_proxy_fail_alloc_mem"].(int)
		oi.Msg_proxy_unexpected_err = in["msg_proxy_unexpected_err"].(int)
		oi.Msg_proxy_l7_cpu_failed = in["msg_proxy_l7_cpu_failed"].(int)
		oi.Msg_proxy_l4_to_l7 = in["msg_proxy_l4_to_l7"].(int)
		oi.Msg_proxy_l4_from_l7 = in["msg_proxy_l4_from_l7"].(int)
		oi.Msg_proxy_to_l4_send_pkt = in["msg_proxy_to_l4_send_pkt"].(int)
		oi.Msg_proxy_l4_from_l4_send = in["msg_proxy_l4_from_l4_send"].(int)
		oi.Msg_proxy_l7_to_l4 = in["msg_proxy_l7_to_l4"].(int)
		oi.Msg_proxy_mag_back = in["msg_proxy_mag_back"].(int)
		oi.Msg_proxy_fail_dcmsg = in["msg_proxy_fail_dcmsg"].(int)
		oi.Msg_proxy_deprecated_conn = in["msg_proxy_deprecated_conn"].(int)
		oi.Msg_proxy_hold_msg = in["msg_proxy_hold_msg"].(int)
		oi.Msg_proxy_split_pkt = in["msg_proxy_split_pkt"].(int)
		oi.Msg_proxy_pipline_msg = in["msg_proxy_pipline_msg"].(int)
		oi.Msg_proxy_client_reset = in["msg_proxy_client_reset"].(int)
		oi.Msg_proxy_server_reset = in["msg_proxy_server_reset"].(int)
		oi.Session_created = in["session_created"].(int)
		oi.Session_freed = in["session_freed"].(int)
		oi.Session_in_rml = in["session_in_rml"].(int)
		oi.Session_invalid = in["session_invalid"].(int)
		oi.Conn_allocd = in["conn_allocd"].(int)
		oi.Conn_freed = in["conn_freed"].(int)
		oi.Session_callid_allocd = in["session_callid_allocd"].(int)
		oi.Session_callid_freed = in["session_callid_freed"].(int)
		oi.Line_mem_allocd = in["line_mem_allocd"].(int)
		oi.Line_mem_freed = in["line_mem_freed"].(int)
		oi.Table_mem_allocd = in["table_mem_allocd"].(int)
		oi.Table_mem_freed = in["table_mem_freed"].(int)
		oi.Cmsg_no_uri_header = in["cmsg_no_uri_header"].(int)
		oi.Cmsg_no_uri_session = in["cmsg_no_uri_session"].(int)
		oi.Sg_no_uri_header = in["sg_no_uri_header"].(int)
		oi.Smsg_no_uri_session = in["smsg_no_uri_session"].(int)
		oi.Line_too_long = in["line_too_long"].(int)
		oi.Fail_read_start_line = in["fail_read_start_line"].(int)
		oi.Fail_parse_start_line = in["fail_parse_start_line"].(int)
		oi.Invalid_start_line = in["invalid_start_line"].(int)
		oi.Request_unknown_version = in["request_unknown_version"].(int)
		oi.Response_unknown_version = in["response_unknown_version"].(int)
		oi.Request_unknown = in["request_unknown"].(int)
		oi.Fail_parse_headers = in["fail_parse_headers"].(int)
		oi.Too_many_headers = in["too_many_headers"].(int)
		oi.Invalid_header = in["invalid_header"].(int)
		oi.Header_name_too_long = in["header_name_too_long"].(int)
		oi.Body_too_big = in["body_too_big"].(int)
		oi.Fail_get_counter = in["fail_get_counter"].(int)
		oi.Msg_no_call_id = in["msg_no_call_id"].(int)
		oi.Identify_dir_failed = in["identify_dir_failed"].(int)
		oi.No_sip_request = in["no_sip_request"].(int)
		oi.Deprecated_msg = in["deprecated_msg"].(int)
		oi.Fail_insert_callid_session = in["fail_insert_callid_session"].(int)
		oi.Fail_insert_uri_session = in["fail_insert_uri_session"].(int)
		oi.Fail_insert_header = in["fail_insert_header"].(int)
		oi.Select_server_conn = in["select_server_conn"].(int)
		oi.Select_server_conn_by_callid = in["select_server_conn_by_callid"].(int)
		oi.Select_server_conn_by_uri = in["select_server_conn_by_uri"].(int)
		oi.Select_server_conn_by_rev_tuple = in["select_server_conn_by_rev_tuple"].(int)
		oi.Select_server_conn_failed = in["select_server_conn_failed"].(int)
		oi.Select_client_conn = in["select_client_conn"].(int)
		oi.X_forward_for_select_client = in["x_forward_for_select_client"].(int)
		oi.Call_id_select_client = in["call_id_select_client"].(int)
		oi.Uri_select_client = in["uri_select_client"].(int)
		oi.Client_select_failed = in["client_select_failed"].(int)
		oi.Acl_denied = in["acl_denied"].(int)
		oi.Assemble_frag_failed = in["assemble_frag_failed"].(int)
		oi.Wrong_ip_version = in["wrong_ip_version"].(int)
		oi.Size_too_large = in["size_too_large"].(int)
		oi.Fail_split_fragment = in["fail_split_fragment"].(int)
		oi.Client_keepalive_received = in["client_keepalive_received"].(int)
		oi.Server_keepalive_received = in["server_keepalive_received"].(int)
		oi.Client_keepalive_send = in["client_keepalive_send"].(int)
		oi.Server_keepalive_send = in["server_keepalive_send"].(int)
		oi.Ax_health_check_received = in["ax_health_check_received"].(int)
		oi.Client_request = in["client_request"].(int)
		oi.Client_request_ok = in["client_request_ok"].(int)
		oi.Concatenate_msg = in["concatenate_msg"].(int)
		oi.Save_uri = in["save_uri"].(int)
		oi.Save_uri_ok = in["save_uri_ok"].(int)
		oi.Save_call_id = in["save_call_id"].(int)
		oi.Save_call_id_ok = in["save_call_id_ok"].(int)
		oi.Msg_translation = in["msg_translation"].(int)
		oi.Msg_translation_fail = in["msg_translation_fail"].(int)
		oi.Msg_trans_start_line = in["msg_trans_start_line"].(int)
		oi.Msg_trans_start_headers = in["msg_trans_start_headers"].(int)
		oi.Msg_trans_body = in["msg_trans_body"].(int)
		oi.Request_register = in["request_register"].(int)
		oi.Request_invite = in["request_invite"].(int)
		oi.Request_ack = in["request_ack"].(int)
		oi.Request_cancel = in["request_cancel"].(int)
		oi.Request_bye = in["request_bye"].(int)
		oi.Request_options = in["request_options"].(int)
		oi.Request_prack = in["request_prack"].(int)
		oi.Request_subscribe = in["request_subscribe"].(int)
		oi.Request_notify = in["request_notify"].(int)
		oi.Request_publish = in["request_publish"].(int)
		oi.Request_info = in["request_info"].(int)
		oi.Request_refer = in["request_refer"].(int)
		oi.Request_message = in["request_message"].(int)
		oi.Request_update = in["request_update"].(int)
		oi.Response_unknown = in["response_unknown"].(int)
		oi.Response_1xx = in["response_1xx"].(int)
		oi.Response_2xx = in["response_2xx"].(int)
		oi.Response_3xx = in["response_3xx"].(int)
		oi.Response_4xx = in["response_4xx"].(int)
		oi.Response_5xx = in["response_5xx"].(int)
		oi.Response_6xx = in["response_6xx"].(int)
		oi.Ha_send_sip_session = in["ha_send_sip_session"].(int)
		oi.Ha_send_sip_session_ok = in["ha_send_sip_session_ok"].(int)
		oi.Ha_fail_get_msg_header = in["ha_fail_get_msg_header"].(int)
		oi.Ha_recv_sip_session = in["ha_recv_sip_session"].(int)
		oi.Ha_insert_sip_session_ok = in["ha_insert_sip_session_ok"].(int)
		oi.Ha_update_sip_session_ok = in["ha_update_sip_session_ok"].(int)
		oi.Ha_invalid_pkt = in["ha_invalid_pkt"].(int)
		oi.Ha_fail_alloc_sip_session = in["ha_fail_alloc_sip_session"].(int)
		oi.Ha_fail_alloc_call_id = in["ha_fail_alloc_call_id"].(int)
		oi.Ha_fail_clone_sip_session = in["ha_fail_clone_sip_session"].(int)
		oi.Save_smp_call_id_rtp = in["save_smp_call_id_rtp"].(int)
		oi.Update_smp_call_id_rtp = in["update_smp_call_id_rtp"].(int)
		oi.Smp_call_id_rtp_session_match = in["smp_call_id_rtp_session_match"].(int)
		oi.Smp_call_id_rtp_session_not_match = in["smp_call_id_rtp_session_not_match"].(int)
		oi.Process_error_when_message_switch = in["process_error_when_message_switch"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSipOper(d *schema.ResourceData) edpt.SlbSipOper {
	var ret edpt.SlbSipOper

	ret.Oper = getObjectSlbSipOperOper(d.Get("oper").([]interface{}))
	return ret
}
