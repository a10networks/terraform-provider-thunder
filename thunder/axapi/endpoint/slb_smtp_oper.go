

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSmtpOper struct {
    
    Oper SlbSmtpOperOper `json:"oper"`

}
type DataSlbSmtpOper struct {
    DtSlbSmtpOper SlbSmtpOper `json:"smtp"`
}


type SlbSmtpOperOper struct {
    SmtpCpuList []SlbSmtpOperOperSmtpCpuList `json:"smtp-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbSmtpOperOperSmtpCpuList struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Request int `json:"request"`
    Request_success int `json:"request_success"`
    No_proxy int `json:"no_proxy"`
    Client_reset int `json:"client_reset"`
    Server_reset int `json:"server_reset"`
    No_tuple int `json:"no_tuple"`
    Parse_req_fail int `json:"parse_req_fail"`
    Server_select_fail int `json:"server_select_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Req_retran int `json:"req_retran"`
    Req_ofo int `json:"req_ofo"`
    Server_reselect int `json:"server_reselect"`
    Server_prem_close int `json:"server_prem_close"`
    New_server_conn int `json:"new_server_conn"`
    Snat_fail int `json:"snat_fail"`
    Tcp_out_reset int `json:"tcp_out_reset"`
    Recv_client_command_ehlo int `json:"recv_client_command_EHLO"`
    Recv_client_command_helo int `json:"recv_client_command_HELO"`
    Recv_client_command_mail int `json:"recv_client_command_MAIL"`
    Recv_client_command_rcpt int `json:"recv_client_command_RCPT"`
    Recv_client_command_data int `json:"recv_client_command_DATA"`
    Recv_client_command_rset int `json:"recv_client_command_RSET"`
    Recv_client_command_vrfy int `json:"recv_client_command_VRFY"`
    Recv_client_command_expn int `json:"recv_client_command_EXPN"`
    Recv_client_command_help int `json:"recv_client_command_HELP"`
    Recv_client_command_noop int `json:"recv_client_command_NOOP"`
    Recv_client_command_quit int `json:"recv_client_command_QUIT"`
    Recv_client_command_starttls int `json:"recv_client_command_STARTTLS"`
    Recv_client_command_turn int `json:"recv_client_command_TURN"`
    Recv_client_command_etrn int `json:"recv_client_command_ETRN"`
    Recv_client_command_others int `json:"recv_client_command_others"`
    Recv_server_service_not_ready int `json:"recv_server_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Send_client_service_ready int `json:"send_client_service_ready"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Send_client_close_connection int `json:"send_client_close_connection"`
    Send_client_go_ahead int `json:"send_client_go_ahead"`
    Send_client_start_tls_first int `json:"send_client_start_TLS_first"`
    Send_client_tls_not_available int `json:"send_client_TLS_not_available"`
    Send_client_no_command int `json:"send_client_no_command"`
    Send_server_cmd_reset int `json:"send_server_cmd_reset"`
    Tls_established int `json:"TLS_established"`
    L4_switch int `json:"L4_switch"`
    Aflex_switch int `json:"Aflex_switch"`
    Aflex_switch_ok int `json:"Aflex_switch_ok"`
    Client_domain_switch int `json:"client_domain_switch"`
    Client_domain_switch_ok int `json:"client_domain_switch_ok"`
    Lb_switch int `json:"LB_switch"`
    Lb_switch_ok int `json:"LB_switch_ok"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Line_across_packet int `json:"line_across_packet"`
    Line_extend int `json:"line_extend"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend int `json:"line_table_extend"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Aflex_lb_reselect int `json:"Aflex_lb_reselect"`
    Aflex_lb_reselect_ok int `json:"Aflex_lb_reselect_ok"`
    Server_starttls_init int `json:"server_STARTTLS_init"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Rserver_starttls_disable int `json:"rserver_STARTTLS_disable"`
    Send_server_ehlo int `json:"send_server_ehlo"`
    Fail_to_save_client_ehlo int `json:"fail_to_save_client_ehlo"`
    Aflex_mail_fail int `json:"aflex_mail_fail"`
    Drop_server_ehlo_ok int `json:"drop_server_ehlo_ok"`
    Client_ehlo_saved int `json:"client_ehlo_saved"`
}

func (p *SlbSmtpOper) GetId() string{
    return "1"
}

func (p *SlbSmtpOper) getPath() string{
    return "slb/smtp/oper"
}

func (p *SlbSmtpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSmtpOper,error) {
logger.Println("SlbSmtpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSmtpOper
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
