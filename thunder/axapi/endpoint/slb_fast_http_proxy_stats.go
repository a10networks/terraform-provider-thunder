

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFastHttpProxyStats struct {
    
    Stats SlbFastHttpProxyStatsStats `json:"stats"`

}
type DataSlbFastHttpProxyStats struct {
    DtSlbFastHttpProxyStats SlbFastHttpProxyStats `json:"fast-http-proxy"`
}


type SlbFastHttpProxyStatsStats struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Req int `json:"req"`
    Req_succ int `json:"req_succ"`
    Noproxy int `json:"noproxy"`
    Client_rst int `json:"client_rst"`
    Server_rst int `json:"server_rst"`
    Notuple int `json:"notuple"`
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Req_retran int `json:"req_retran"`
    Req_ofo int `json:"req_ofo"`
    Server_resel int `json:"server_resel"`
    Svr_prem_close int `json:"svr_prem_close"`
    New_svrconn int `json:"new_svrconn"`
    Snat_fail int `json:"snat_fail"`
    Tcpoutrst int `json:"tcpoutrst"`
    Full_proxy int `json:"full_proxy"`
    Full_proxy_post int `json:"full_proxy_post"`
    Full_proxy_pipeline int `json:"full_proxy_pipeline"`
    Full_proxy_fpga_err int `json:"full_proxy_fpga_err"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Close_on_ddos int `json:"close_on_ddos"`
    Full_proxy_put int `json:"full_proxy_put"`
}

func (p *SlbFastHttpProxyStats) GetId() string{
    return "1"
}

func (p *SlbFastHttpProxyStats) getPath() string{
    return "slb/fast-http-proxy/stats"
}

func (p *SlbFastHttpProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFastHttpProxyStats,error) {
logger.Println("SlbFastHttpProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFastHttpProxyStats
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
