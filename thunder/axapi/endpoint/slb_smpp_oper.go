

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSmppOper struct {
    
    Oper SlbSmppOperOper `json:"oper"`

}
type DataSlbSmppOper struct {
    DtSlbSmppOper SlbSmppOper `json:"smpp"`
}


type SlbSmppOperOper struct {
    SmppCpuFields []SlbSmppOperOperSmppCpuFields `json:"smpp-cpu-fields"`
    CpuCount int `json:"cpu-count"`
    Filter_type string `json:"filter_type"`
}


type SlbSmppOperOperSmppCpuFields struct {
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
    Payload_allocd int `json:"payload_allocd"`
    Payload_freed int `json:"payload_freed"`
    Pkt_too_small int `json:"pkt_too_small"`
    Invalid_seq int `json:"invalid_seq"`
    Ax_response_directly int `json:"AX_response_directly"`
    Select_client_conn int `json:"select_client_conn"`
    Select_client_by_req int `json:"select_client_by_req"`
    Select_client_from_list int `json:"select_client_from_list"`
    Select_client_by_conn int `json:"select_client_by_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_conn int `json:"select_server_conn"`
    Select_server_by_req int `json:"select_server_by_req"`
    Select_server_from_list int `json:"select_server_from_list"`
    Select_server_by_conn int `json:"select_server_by_conn"`
    Select_server_fail int `json:"select_server_fail"`
    Bind_conn int `json:"bind_conn"`
    Unbind_conn int `json:"unbind_conn"`
    Enquire_link_recv int `json:"enquire_link_recv"`
    Enquire_link_resp_recv int `json:"enquire_link_resp_recv"`
    Enquire_link_send int `json:"enquire_link_send"`
    Enquire_link_resp_send int `json:"enquire_link_resp_send"`
    Client_conn_put_in_list int `json:"client_conn_put_in_list"`
    Client_conn_get_from_list int `json:"client_conn_get_from_list"`
    Server_conn_put_in_list int `json:"server_conn_put_in_list"`
    Server_conn_get_from_list int `json:"server_conn_get_from_list"`
    Server_conn_fail_bind int `json:"server_conn_fail_bind"`
    Single_msg int `json:"single_msg"`
    Fail_bind_msg int `json:"fail_bind_msg"`
}

func (p *SlbSmppOper) GetId() string{
    return "1"
}

func (p *SlbSmppOper) getPath() string{
    return "slb/smpp/oper"
}

func (p *SlbSmppOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSmppOper,error) {
logger.Println("SlbSmppOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSmppOper
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
