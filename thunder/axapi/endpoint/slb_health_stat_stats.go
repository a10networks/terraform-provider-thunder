

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthStatStats struct {
    
    Stats SlbHealthStatStatsStats `json:"stats"`

}
type DataSlbHealthStatStats struct {
    DtSlbHealthStatStats SlbHealthStatStats `json:"health-stat"`
}


type SlbHealthStatStatsStats struct {
    Num_burst int `json:"num_burst"`
    Max_jiffie int `json:"max_jiffie"`
    Min_jiffie int `json:"min_jiffie"`
    Avg_jiffie int `json:"avg_jiffie"`
    Open_socket int `json:"open_socket"`
    Open_socket_failed int `json:"open_socket_failed"`
    Close_socket int `json:"close_socket"`
    Connect_failed int `json:"connect_failed"`
    Send_packet int `json:"send_packet"`
    Send_packet_failed int `json:"send_packet_failed"`
    Recv_packet int `json:"recv_packet"`
    Recv_packet_failed int `json:"recv_packet_failed"`
    Retry_times int `json:"retry_times"`
    Timeout int `json:"timeout"`
    Unexpected_error int `json:"unexpected_error"`
    Conn_imdt_succ int `json:"conn_imdt_succ"`
    Sock_close_before_17 int `json:"sock_close_before_17"`
    Sock_close_without_notify int `json:"sock_close_without_notify"`
    Curr_health_rate int `json:"curr_health_rate"`
    Ext_health_rate int `json:"ext_health_rate"`
    Ext_health_rate_val int `json:"ext_health_rate_val"`
    Total_number int `json:"total_number"`
    Status_up int `json:"status_up"`
    Status_down int `json:"status_down"`
    Status_unkn int `json:"status_unkn"`
    Status_other int `json:"status_other"`
    Running_time int `json:"running_time"`
    Config_health_rate int `json:"config_health_rate"`
    Ssl_post_handshake_packet int `json:"ssl_post_handshake_packet"`
    Timeout_with_packet int `json:"timeout_with_packet"`
}

func (p *SlbHealthStatStats) GetId() string{
    return "1"
}

func (p *SlbHealthStatStats) getPath() string{
    return "slb/health-stat/stats"
}

func (p *SlbHealthStatStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthStatStats,error) {
logger.Println("SlbHealthStatStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthStatStats
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
