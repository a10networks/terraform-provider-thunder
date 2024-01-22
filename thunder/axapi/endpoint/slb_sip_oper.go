

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSipOper struct {
    
    Oper SlbSipOperOper `json:"oper"`

}
type DataSlbSipOper struct {
    DtSlbSipOper SlbSipOper `json:"sip"`
}


type SlbSipOperOper struct {
    SipCpuList []SlbSipOperOperSipCpuList `json:"sip-cpu-list"`
    CpuCount int `json:"cpu-count"`
    Filter_type string `json:"filter_type"`
}


type SlbSipOperOperSipCpuList struct {
    Msg_proxy_current int `json:"msg_proxy_current"`
    Msg_proxy_total int `json:"msg_proxy_total"`
    Msg_proxy_mem_allocd int `json:"msg_proxy_mem_allocd"`
    Msg_proxy_mem_cached int `json:"msg_proxy_mem_cached"`
    Msg_proxy_mem_freed int `json:"msg_proxy_mem_freed"`
    Msg_proxy_client_recv int `json:"msg_proxy_client_recv"`
    Msg_proxy_client_send_success int `json:"msg_proxy_client_send_success"`
    Msg_proxy_client_incomplete int `json:"msg_proxy_client_incomplete"`
    Msg_proxy_client_drop int `json:"msg_proxy_client_drop"`
    Msg_proxy_client_connection int `json:"msg_proxy_client_connection"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_client_fail_parse int `json:"msg_proxy_client_fail_parse"`
    Msg_proxy_client_fail_process int `json:"msg_proxy_client_fail_process"`
    Msg_proxy_client_fail_snat int `json:"msg_proxy_client_fail_snat"`
    Msg_proxy_client_exceed_tmp_buff int `json:"msg_proxy_client_exceed_tmp_buff"`
    Msg_proxy_client_fail_send_pkt int `json:"msg_proxy_client_fail_send_pkt"`
    Msg_proxy_client_fail_start_server_conn int `json:"msg_proxy_client_fail_start_server_Conn"`
    Msg_proxy_server_recv int `json:"msg_proxy_server_recv"`
    Msg_proxy_server_send_success int `json:"msg_proxy_server_send_success"`
    Msg_proxy_server_incomplete int `json:"msg_proxy_server_incomplete"`
    Msg_proxy_server_drop int `json:"msg_proxy_server_drop"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_server_fail_parse int `json:"msg_proxy_server_fail_parse"`
    Msg_proxy_server_fail_process int `json:"msg_proxy_server_fail_process"`
    Msg_proxy_server_fail_selec_connt int `json:"msg_proxy_server_fail_selec_connt"`
    Msg_proxy_server_fail_snat int `json:"msg_proxy_server_fail_snat"`
    Msg_proxy_server_exceed_tmp_buff int `json:"msg_proxy_server_exceed_tmp_buff"`
    Msg_proxy_server_fail_send_pkt int `json:"msg_proxy_server_fail_send_pkt"`
    Msg_proxy_create_server_conn int `json:"msg_proxy_create_server_conn"`
    Msg_proxy_start_server_conn int `json:"msg_proxy_start_server_conn"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Msg_proxy_server_conn_fail_snat int `json:"msg_proxy_server_conn_fail_snat"`
    Msg_proxy_fail_construct_server_conn int `json:"msg_proxy_fail_construct_server_conn"`
    Msg_proxy_fail_reserve_pconn int `json:"msg_proxy_fail_reserve_pconn"`
    Msg_proxy_start_server_conn_failed int `json:"msg_proxy_start_server_conn_failed"`
    Msg_proxy_server_conn_already_exists int `json:"msg_proxy_server_conn_already_exists"`
    Msg_proxy_fail_insert_server_conn int `json:"msg_proxy_fail_insert_server_conn"`
    Msg_proxy_parse_msg_fail int `json:"msg_proxy_parse_msg_fail"`
    Msg_proxy_process_msg_fail int `json:"msg_proxy_process_msg_fail"`
    Msg_proxy_no_vport int `json:"msg_proxy_no_vport"`
    Msg_proxy_fail_select_server int `json:"msg_proxy_fail_select_server"`
    Msg_proxy_fail_alloc_mem int `json:"msg_proxy_fail_alloc_mem"`
    Msg_proxy_unexpected_err int `json:"msg_proxy_unexpected_err"`
    Msg_proxy_l7_cpu_failed int `json:"msg_proxy_l7_cpu_failed"`
    Msg_proxy_l4_to_l7 int `json:"msg_proxy_l4_to_l7"`
    Msg_proxy_l4_from_l7 int `json:"msg_proxy_l4_from_l7"`
    Msg_proxy_to_l4_send_pkt int `json:"msg_proxy_to_l4_send_pkt"`
    Msg_proxy_l4_from_l4_send int `json:"msg_proxy_l4_from_l4_send"`
    Msg_proxy_l7_to_l4 int `json:"msg_proxy_l7_to_L4"`
    Msg_proxy_mag_back int `json:"msg_proxy_mag_back"`
    Msg_proxy_fail_dcmsg int `json:"msg_proxy_fail_dcmsg"`
    Msg_proxy_deprecated_conn int `json:"msg_proxy_deprecated_conn"`
    Msg_proxy_hold_msg int `json:"msg_proxy_hold_msg"`
    Msg_proxy_split_pkt int `json:"msg_proxy_split_pkt"`
    Msg_proxy_pipline_msg int `json:"msg_proxy_pipline_msg"`
    Msg_proxy_client_reset int `json:"msg_proxy_client_reset"`
    Msg_proxy_server_reset int `json:"msg_proxy_server_reset"`
    Session_created int `json:"session_created"`
    Session_freed int `json:"session_freed"`
    Session_in_rml int `json:"session_in_rml"`
    Session_invalid int `json:"session_invalid"`
    Conn_allocd int `json:"conn_allocd"`
    Conn_freed int `json:"conn_freed"`
    Session_callid_allocd int `json:"session_callid_allocd"`
    Session_callid_freed int `json:"session_callid_freed"`
    Line_mem_allocd int `json:"line_mem_allocd"`
    Line_mem_freed int `json:"line_mem_freed"`
    Table_mem_allocd int `json:"table_mem_allocd"`
    Table_mem_freed int `json:"table_mem_freed"`
    Cmsg_no_uri_header int `json:"cmsg_no_uri_header"`
    Cmsg_no_uri_session int `json:"cmsg_no_uri_session"`
    Sg_no_uri_header int `json:"sg_no_uri_header"`
    Smsg_no_uri_session int `json:"smsg_no_uri_session"`
    Line_too_long int `json:"line_too_long"`
    Fail_read_start_line int `json:"fail_read_start_line"`
    Fail_parse_start_line int `json:"fail_parse_start_line"`
    Invalid_start_line int `json:"invalid_start_line"`
    Request_unknown_version int `json:"request_unknown_version"`
    Response_unknown_version int `json:"response_unknown_version"`
    Request_unknown int `json:"request_unknown"`
    Fail_parse_headers int `json:"fail_parse_headers"`
    Too_many_headers int `json:"too_many_headers"`
    Invalid_header int `json:"invalid_header"`
    Header_name_too_long int `json:"header_name_too_long"`
    Body_too_big int `json:"body_too_big"`
    Fail_get_counter int `json:"fail_get_counter"`
    Msg_no_call_id int `json:"msg_no_call_id"`
    Identify_dir_failed int `json:"identify_dir_failed"`
    No_sip_request int `json:"no_sip_request"`
    Deprecated_msg int `json:"deprecated_msg"`
    Fail_insert_callid_session int `json:"fail_insert_callid_session"`
    Fail_insert_uri_session int `json:"fail_insert_uri_session"`
    Fail_insert_header int `json:"fail_insert_header"`
    Select_server_conn int `json:"select_server_conn"`
    Select_server_conn_by_callid int `json:"select_server_conn_by_callid"`
    Select_server_conn_by_uri int `json:"select_server_conn_by_uri"`
    Select_server_conn_by_rev_tuple int `json:"select_server_conn_by_rev_tuple"`
    Select_server_conn_failed int `json:"select_server_conn_failed"`
    Select_client_conn int `json:"select_client_conn"`
    X_forward_for_select_client int `json:"X_forward_for_select_client"`
    Call_id_select_client int `json:"call_id_select_client"`
    Uri_select_client int `json:"uri_select_client"`
    Client_select_failed int `json:"client_select_failed"`
    Acl_denied int `json:"acl_denied"`
    Assemble_frag_failed int `json:"assemble_frag_failed"`
    Wrong_ip_version int `json:"wrong_ip_version"`
    Size_too_large int `json:"size_too_large"`
    Fail_split_fragment int `json:"fail_split_fragment"`
    Client_keepalive_received int `json:"client_keepalive_received"`
    Server_keepalive_received int `json:"server_keepalive_received"`
    Client_keepalive_send int `json:"client_keepalive_send"`
    Server_keepalive_send int `json:"server_keepalive_send"`
    Ax_health_check_received int `json:"ax_health_check_received"`
    Client_request int `json:"client_request"`
    Client_request_ok int `json:"client_request_ok"`
    Concatenate_msg int `json:"concatenate_msg"`
    Save_uri int `json:"save_uri"`
    Save_uri_ok int `json:"save_uri_ok"`
    Save_call_id int `json:"save_call_id"`
    Save_call_id_ok int `json:"save_call_id_ok"`
    Msg_translation int `json:"msg_translation"`
    Msg_translation_fail int `json:"msg_translation_fail"`
    Msg_trans_start_line int `json:"msg_trans_start_line"`
    Msg_trans_start_headers int `json:"msg_trans_start_headers"`
    Msg_trans_body int `json:"msg_trans_body"`
    Request_register int `json:"request_register"`
    Request_invite int `json:"request_invite"`
    Request_ack int `json:"request_ack"`
    Request_cancel int `json:"request_cancel"`
    Request_bye int `json:"request_bye"`
    Request_options int `json:"request_options"`
    Request_prack int `json:"request_prack"`
    Request_subscribe int `json:"request_subscribe"`
    Request_notify int `json:"request_notify"`
    Request_publish int `json:"request_publish"`
    Request_info int `json:"request_info"`
    Request_refer int `json:"request_refer"`
    Request_message int `json:"request_message"`
    Request_update int `json:"request_update"`
    Response_unknown int `json:"response_unknown"`
    Response_1xx int `json:"response_1XX"`
    Response_2xx int `json:"response_2XX"`
    Response_3xx int `json:"response_3XX"`
    Response_4xx int `json:"response_4XX"`
    Response_5xx int `json:"response_5XX"`
    Response_6xx int `json:"response_6XX"`
    Ha_send_sip_session int `json:"ha_send_sip_session"`
    Ha_send_sip_session_ok int `json:"ha_send_sip_session_ok"`
    Ha_fail_get_msg_header int `json:"ha_fail_get_msg_header"`
    Ha_recv_sip_session int `json:"ha_recv_sip_session"`
    Ha_insert_sip_session_ok int `json:"ha_insert_sip_session_ok"`
    Ha_update_sip_session_ok int `json:"ha_update_sip_session_ok"`
    Ha_invalid_pkt int `json:"ha_invalid_pkt"`
    Ha_fail_alloc_sip_session int `json:"ha_fail_alloc_sip_session"`
    Ha_fail_alloc_call_id int `json:"ha_fail_alloc_call_id"`
    Ha_fail_clone_sip_session int `json:"ha_fail_clone_sip_session"`
    Save_smp_call_id_rtp int `json:"save_smp_call_id_rtp"`
    Update_smp_call_id_rtp int `json:"update_smp_call_id_rtp"`
    Smp_call_id_rtp_session_match int `json:"smp_call_id_rtp_session_match"`
    Smp_call_id_rtp_session_not_match int `json:"smp_call_id_rtp_session_not_match"`
    Process_error_when_message_switch int `json:"process_error_when_message_switch"`
}

func (p *SlbSipOper) GetId() string{
    return "1"
}

func (p *SlbSipOper) getPath() string{
    return "slb/sip/oper"
}

func (p *SlbSipOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSipOper,error) {
logger.Println("SlbSipOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSipOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
