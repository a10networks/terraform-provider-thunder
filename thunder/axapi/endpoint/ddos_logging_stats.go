

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosLoggingStats struct {
    
    Stats DdosLoggingStatsStats `json:"stats"`

}
type DataDdosLoggingStats struct {
    DtDdosLoggingStats DdosLoggingStats `json:"logging"`
}


type DdosLoggingStatsStats struct {
    Log_msg_quota_exceed int `json:"log_msg_quota_exceed"`
    Log_msg_oom int `json:"log_msg_oom"`
    Log_queue_full int `json:"log_queue_full"`
    Log_msg_sent int `json:"log_msg_sent"`
    Log_msg_send_err int `json:"log_msg_send_err"`
}

func (p *DdosLoggingStats) GetId() string{
    return "1"
}

func (p *DdosLoggingStats) getPath() string{
    return "ddos/logging/stats"
}

func (p *DdosLoggingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosLoggingStats,error) {
logger.Println("DdosLoggingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosLoggingStats
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
