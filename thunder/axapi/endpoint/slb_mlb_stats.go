

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbMlbStats struct {
    
    Stats SlbMlbStatsStats `json:"stats"`

}
type DataSlbMlbStats struct {
    DtSlbMlbStats SlbMlbStats `json:"mlb"`
}


type SlbMlbStatsStats struct {
    Client_msg_sent int `json:"client_msg_sent"`
    Server_msg_received int `json:"server_msg_received"`
    Server_conn_created int `json:"server_conn_created"`
    Server_conn_rst int `json:"server_conn_rst"`
    Server_conn_failed int `json:"server_conn_failed"`
    Server_conn_closed int `json:"server_conn_closed"`
    Client_conn_created int `json:"client_conn_created"`
    Client_conn_closed int `json:"client_conn_closed"`
    Client_conn_not_found int `json:"client_conn_not_found"`
    Msg_dropped int `json:"msg_dropped"`
    Msg_rerouted int `json:"msg_rerouted"`
    Mlb_dcmsg_sent int `json:"mlb_dcmsg_sent"`
    Mlb_dcmsg_received int `json:"mlb_dcmsg_received"`
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Mlb_dcmsg_alloc int `json:"mlb_dcmsg_alloc"`
    Mlb_dcmsg_free int `json:"mlb_dcmsg_free"`
    Mlb_server_probe int `json:"mlb_server_probe"`
    Mlb_server_down int `json:"mlb_server_down"`
}

func (p *SlbMlbStats) GetId() string{
    return "1"
}

func (p *SlbMlbStats) getPath() string{
    return "slb/mlb/stats"
}

func (p *SlbMlbStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbMlbStats,error) {
logger.Println("SlbMlbStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbMlbStats
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
