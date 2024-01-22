

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPop3ProxyOper struct {
    
    Oper SlbPop3ProxyOperOper `json:"oper"`

}
type DataSlbPop3ProxyOper struct {
    DtSlbPop3ProxyOper SlbPop3ProxyOper `json:"pop3-proxy"`
}


type SlbPop3ProxyOperOper struct {
    L4CpuList []SlbPop3ProxyOperOperL4CpuList `json:"l4-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbPop3ProxyOperOperL4CpuList struct {
    Current_proxy_conns int `json:"current_proxy_conns"`
    Total_proxy_conns int `json:"total_proxy_conns"`
    Server_selection_failure int `json:"server_selection_failure"`
    No_route_failure int `json:"no_route_failure"`
    Source_nat_failure int `json:"source_nat_failure"`
    Stls_packet int `json:"stls_packet"`
    Request_line_freed int `json:"request_line_freed"`
    Inv_start_line int `json:"inv_start_line"`
    Other_cmd int `json:"other_cmd"`
    Pop3_line_too_long int `json:"pop3_line_too_long"`
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
    Total_pop3_request int `json:"total_pop3_request"`
}

func (p *SlbPop3ProxyOper) GetId() string{
    return "1"
}

func (p *SlbPop3ProxyOper) getPath() string{
    return "slb/pop3-proxy/oper"
}

func (p *SlbPop3ProxyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPop3ProxyOper,error) {
logger.Println("SlbPop3ProxyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPop3ProxyOper
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
