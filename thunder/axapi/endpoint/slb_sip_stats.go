

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSipStats struct {
    
    Stats SlbSipStatsStats `json:"stats"`

}
type DataSlbSipStats struct {
    DtSlbSipStats SlbSipStats `json:"sip"`
}


type SlbSipStatsStats struct {
    Msg_proxy_current int `json:"msg_proxy_current"`
    Msg_proxy_total int `json:"msg_proxy_total"`
    Msg_proxy_client_recv int `json:"msg_proxy_client_recv"`
    Msg_proxy_client_send_success int `json:"msg_proxy_client_send_success"`
    Msg_proxy_client_incomplete int `json:"msg_proxy_client_incomplete"`
    Msg_proxy_client_drop int `json:"msg_proxy_client_drop"`
    Msg_proxy_client_connection int `json:"msg_proxy_client_connection"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_recv int `json:"msg_proxy_server_recv"`
    Msg_proxy_server_send_success int `json:"msg_proxy_server_send_success"`
    Msg_proxy_server_incomplete int `json:"msg_proxy_server_incomplete"`
    Msg_proxy_server_drop int `json:"msg_proxy_server_drop"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_create_server_conn int `json:"msg_proxy_create_server_conn"`
    Msg_proxy_start_server_conn int `json:"msg_proxy_start_server_conn"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Session_created int `json:"session_created"`
    Session_freed int `json:"session_freed"`
}

func (p *SlbSipStats) GetId() string{
    return "1"
}

func (p *SlbSipStats) getPath() string{
    return "slb/sip/stats"
}

func (p *SlbSipStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSipStats,error) {
logger.Println("SlbSipStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSipStats
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
