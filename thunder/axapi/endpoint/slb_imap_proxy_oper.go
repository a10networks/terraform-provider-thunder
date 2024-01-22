

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbImapProxyOper struct {
    
    Oper SlbImapProxyOperOper `json:"oper"`

}
type DataSlbImapProxyOper struct {
    DtSlbImapProxyOper SlbImapProxyOper `json:"imap-proxy"`
}


type SlbImapProxyOperOper struct {
    ImapProxyCpuList []SlbImapProxyOperOperImapProxyCpuList `json:"imap-proxy-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbImapProxyOperOperImapProxyCpuList struct {
    Current_proxy_conns int `json:"current_proxy_conns"`
    Total_proxy_conns int `json:"total_proxy_conns"`
    Total_imap_request int `json:"total_imap_request"`
    Server_selection_failure int `json:"server_selection_failure"`
    No_route_failure int `json:"no_route_failure"`
    Source_nat_failure int `json:"source_nat_failure"`
    Start_tls_cmd int `json:"start_tls_cmd"`
    Login_packet int `json:"login_packet"`
    Capability_packet int `json:"capability_packet"`
    Request_line_freed int `json:"request_line_freed"`
    Inv_start_line int `json:"inv_start_line"`
    Other_cmd int `json:"other_cmd"`
    Imap_line_too_long int `json:"imap_line_too_long"`
    Control_chn_ssl int `json:"control_chn_ssl"`
    Bad_seq int `json:"bad_seq"`
    Serv_sel_persist_fail int `json:"serv_sel_persist_fail"`
    Serv_sel_smpv6_fail int `json:"serv_sel_smpv6_fail"`
    Serv_sel_smpv4_fail int `json:"serv_sel_smpv4_fail"`
    Serv_sel_ins_tpl_fail int `json:"serv_sel_ins_tpl_fail"`
    Client_est_state_err int `json:"client_est_state_err"`
    Serv_ctng_state_err int `json:"serv_ctng_state_err"`
    Serv_resp_state_err int `json:"serv_resp_state_err"`
    Client_rq_state_err int `json:"client_rq_state_err"`
    Realloc_error int `json:"realloc_error"`
    Alloc_error int `json:"alloc_error"`
    Boundary_error int `json:"boundary_error"`
    Negative_error int `json:"negative_error"`
}

func (p *SlbImapProxyOper) GetId() string{
    return "1"
}

func (p *SlbImapProxyOper) getPath() string{
    return "slb/imap-proxy/oper"
}

func (p *SlbImapProxyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbImapProxyOper,error) {
logger.Println("SlbImapProxyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbImapProxyOper
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
