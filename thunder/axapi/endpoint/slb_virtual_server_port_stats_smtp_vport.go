

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats54 struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStats54Stats `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbVirtualServerPortStats54Stats struct {
    Smtp_vport SlbVirtualServerPortStats54StatsSmtp_vport `json:"smtp_vport"`
}


type SlbVirtualServerPortStats54StatsSmtp_vport struct {
    Cmd_ehlo int `json:"cmd_ehlo"`
    Cmd_helo int `json:"cmd_helo"`
    Cmd_mail int `json:"cmd_mail"`
    Cmd_rcpt int `json:"cmd_rcpt"`
    Cmd_data int `json:"cmd_data"`
    Cmd_rset int `json:"cmd_rset"`
    Cmd_vrfy int `json:"cmd_vrfy"`
    Cmd_expn int `json:"cmd_expn"`
    Cmd_help int `json:"cmd_help"`
    Cmd_noop int `json:"cmd_noop"`
    Cmd_starttls int `json:"cmd_starttls"`
    Cmd_turn int `json:"cmd_turn"`
    Cmd_etrn int `json:"cmd_etrn"`
    Cmd_quit int `json:"cmd_quit"`
    Cmd_local int `json:"cmd_local"`
    Cmd_unknown int `json:"cmd_unknown"`
    Code_200 int `json:"code_200"`
    Code_211 int `json:"code_211"`
    Code_214 int `json:"code_214"`
    Code_220 int `json:"code_220"`
    Code_221 int `json:"code_221"`
    Code_250 int `json:"code_250"`
    Code_251 int `json:"code_251"`
    Code_252 int `json:"code_252"`
    Code_354 int `json:"code_354"`
    Code_421 int `json:"code_421"`
    Code_450 int `json:"code_450"`
    Code_451 int `json:"code_451"`
    Code_452 int `json:"code_452"`
    Code_455 int `json:"code_455"`
    Code_500 int `json:"code_500"`
    Code_501 int `json:"code_501"`
    Code_502 int `json:"code_502"`
    Code_503 int `json:"code_503"`
    Code_504 int `json:"code_504"`
    Code_521 int `json:"code_521"`
    Code_530 int `json:"code_530"`
    Code_550 int `json:"code_550"`
    Code_551 int `json:"code_551"`
    Code_552 int `json:"code_552"`
    Code_553 int `json:"code_553"`
    Code_554 int `json:"code_554"`
    Code_555 int `json:"code_555"`
    Code_556 int `json:"code_556"`
    Code_2xx int `json:"code_2xx"`
    Code_3xx int `json:"code_3xx"`
    Code_4xx int `json:"code_4xx"`
    Code_5xx int `json:"code_5xx"`
    Code_unknown int `json:"code_unknown"`
    Reply_10u int `json:"reply_10u"`
    Reply_20u int `json:"reply_20u"`
    Reply_50u int `json:"reply_50u"`
    Reply_100u int `json:"reply_100u"`
    Reply_200u int `json:"reply_200u"`
    Reply_500u int `json:"reply_500u"`
    Reply_1m int `json:"reply_1m"`
    Reply_2m int `json:"reply_2m"`
    Reply_5m int `json:"reply_5m"`
    Reply_10m int `json:"reply_10m"`
    Reply_20m int `json:"reply_20m"`
    Reply_50m int `json:"reply_50m"`
    Reply_100m int `json:"reply_100m"`
    Reply_200m int `json:"reply_200m"`
    Reply_500m int `json:"reply_500m"`
    Reply_1s int `json:"reply_1s"`
    Reply_2s int `json:"reply_2s"`
    Reply_5s int `json:"reply_5s"`
    Reply_over_5s int `json:"reply_over_5s"`
    Data_sz_1k int `json:"data_sz_1k"`
    Data_sz_2k int `json:"data_sz_2k"`
    Data_sz_4k int `json:"data_sz_4k"`
    Data_sz_8k int `json:"data_sz_8k"`
    Data_sz_16k int `json:"data_sz_16k"`
    Data_sz_32k int `json:"data_sz_32k"`
    Data_sz_64k int `json:"data_sz_64k"`
    Data_sz_128k int `json:"data_sz_128k"`
    Data_sz_256k int `json:"data_sz_256k"`
    Data_sz_512k int `json:"data_sz_512k"`
    Data_sz_1m int `json:"data_sz_1m"`
    Data_sz_gt_1m int `json:"data_sz_gt_1m"`
    Total_commands int `json:"total_commands"`
    Total_replies int `json:"total_replies"`
    Curr_conn int `json:"curr_conn"`
    Total_conn int `json:"total_conn"`
    Peak_conn int `json:"peak_conn"`
    Client_bytes int `json:"client_bytes"`
    Server_bytes int `json:"server_bytes"`
    Aflex_switch int `json:"aflex_switch"`
    Aflex_switch_ok int `json:"aflex_switch_ok"`
    Aflex_ehlo_sent int `json:"aflex_ehlo_sent"`
    Send_server_ehlo int `json:"send_server_ehlo"`
    Fail_to_save_client_ehlo int `json:"fail_to_save_client_ehlo"`
    Aflex_mail_fail int `json:"aflex_mail_fail"`
    Drop_server_ehlo_ok int `json:"drop_server_ehlo_ok"`
    Client_ehlo_saved int `json:"client_ehlo_saved"`
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
    Recv_server_service_not_ready int `json:"recv_server_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
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
    Recv_client_command_turn int `json:"recv_client_command_TURN"`
    Recv_client_command_etrn int `json:"recv_client_command_ETRN"`
    Recv_client_command_others int `json:"recv_client_command_others"`
    Recv_client_command_starttls int `json:"recv_client_command_STARTTLS"`
    Recv_client_command_rcpt int `json:"recv_client_command_RCPT"`
}

func (p *SlbVirtualServerPortStats54) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats54) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?smtp_vport=true"
}

func (p *SlbVirtualServerPortStats54) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats54::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SlbVirtualServerPortStats54) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats54::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SlbVirtualServerPortStats54) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats54::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SlbVirtualServerPortStats54) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats54::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
