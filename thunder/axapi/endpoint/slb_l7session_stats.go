

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbL7sessionStats struct {
    
    Stats SlbL7sessionStatsStats `json:"stats"`

}
type DataSlbL7sessionStats struct {
    DtSlbL7sessionStats SlbL7sessionStats `json:"l7session"`
}


type SlbL7sessionStatsStats struct {
    Start_server_conn_succ int `json:"start_server_conn_succ"`
    Conn_not_exist int `json:"conn_not_exist"`
    Data_event int `json:"data_event"`
    Client_fin int `json:"client_fin"`
    Server_fin int `json:"server_fin"`
    Wbuf_event int `json:"wbuf_event"`
    Wbuf_cb_failed int `json:"wbuf_cb_failed"`
    Err_event int `json:"err_event"`
    Err_cb_failed int `json:"err_cb_failed"`
    Server_conn_failed int `json:"server_conn_failed"`
    Client_rst int `json:"client_rst"`
    Server_rst int `json:"server_rst"`
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Server_select_fail int `json:"server_select_fail"`
    Data_cb_failed int `json:"data_cb_failed"`
    Hps_fwdreq_fail int `json:"hps_fwdreq_fail"`
    Udp_data_event int `json:"udp_data_event"`
}

func (p *SlbL7sessionStats) GetId() string{
    return "1"
}

func (p *SlbL7sessionStats) getPath() string{
    return "slb/l7session/stats"
}

func (p *SlbL7sessionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbL7sessionStats,error) {
logger.Println("SlbL7sessionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbL7sessionStats
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
