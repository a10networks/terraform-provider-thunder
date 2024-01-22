

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSmppStats struct {
    
    Stats SlbSmppStatsStats `json:"stats"`

}
type DataSlbSmppStats struct {
    DtSlbSmppStats SlbSmppStats `json:"smpp"`
}


type SlbSmppStatsStats struct {
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
}

func (p *SlbSmppStats) GetId() string{
    return "1"
}

func (p *SlbSmppStats) getPath() string{
    return "slb/smpp/stats"
}

func (p *SlbSmppStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSmppStats,error) {
logger.Println("SlbSmppStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSmppStats
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
