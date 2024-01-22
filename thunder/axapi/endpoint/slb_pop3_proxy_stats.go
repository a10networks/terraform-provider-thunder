

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPop3ProxyStats struct {
    
    Stats SlbPop3ProxyStatsStats `json:"stats"`

}
type DataSlbPop3ProxyStats struct {
    DtSlbPop3ProxyStats SlbPop3ProxyStats `json:"pop3-proxy"`
}


type SlbPop3ProxyStatsStats struct {
    Num int `json:"num"`
    Curr int `json:"curr"`
    Total int `json:"total"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Line_mem_freed int `json:"line_mem_freed"`
    Invalid_start_line int `json:"invalid_start_line"`
    Stls int `json:"stls"`
    Request_dont_care int `json:"request_dont_care"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Request int `json:"request"`
    Control_to_ssl int `json:"control_to_ssl"`
}

func (p *SlbPop3ProxyStats) GetId() string{
    return "1"
}

func (p *SlbPop3ProxyStats) getPath() string{
    return "slb/pop3-proxy/stats"
}

func (p *SlbPop3ProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPop3ProxyStats,error) {
logger.Println("SlbPop3ProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPop3ProxyStats
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
