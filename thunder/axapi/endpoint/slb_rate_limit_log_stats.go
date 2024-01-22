

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbRateLimitLogStats struct {
    
    Stats SlbRateLimitLogStatsStats `json:"stats"`

}
type DataSlbRateLimitLogStats struct {
    DtSlbRateLimitLogStats SlbRateLimitLogStats `json:"rate-limit-log"`
}


type SlbRateLimitLogStatsStats struct {
    Total_log_times int `json:"total_log_times"`
    Total_log_msg int `json:"total_log_msg"`
    Local_log_msg int `json:"local_log_msg"`
    Remote_log_msg int `json:"remote_log_msg"`
    Local_log_rate int `json:"local_log_rate"`
    Remote_log_rate int `json:"remote_log_rate"`
    Msg_too_big int `json:"msg_too_big"`
    Buff_alloc_fail int `json:"buff_alloc_fail"`
    No_route int `json:"no_route"`
    Buff_send_fail int `json:"buff_send_fail"`
    Alloc_conn int `json:"alloc_conn"`
    Free_conn int `json:"free_conn"`
    Conn_alloc_fail int `json:"conn_alloc_fail"`
    No_repeat_msg int `json:"no_repeat_msg"`
    Local_log_dropped int `json:"local_log_dropped"`
}

func (p *SlbRateLimitLogStats) GetId() string{
    return "1"
}

func (p *SlbRateLimitLogStats) getPath() string{
    return "slb/rate-limit-log/stats"
}

func (p *SlbRateLimitLogStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbRateLimitLogStats,error) {
logger.Println("SlbRateLimitLogStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbRateLimitLogStats
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
