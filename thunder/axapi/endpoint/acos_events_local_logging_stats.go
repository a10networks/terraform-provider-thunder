

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsLocalLoggingStats struct {
    
    Stats AcosEventsLocalLoggingStatsStats `json:"stats"`

}
type DataAcosEventsLocalLoggingStats struct {
    DtAcosEventsLocalLoggingStats AcosEventsLocalLoggingStats `json:"local-logging"`
}


type AcosEventsLocalLoggingStatsStats struct {
    InitPass int `json:"init-pass"`
    InitFail int `json:"init-fail"`
    Freed int `json:"freed"`
    DiskOverThres int `json:"disk-over-thres"`
    RateLimited int `json:"rate-limited"`
    NotInited int `json:"not-inited"`
    SentToStore int `json:"sent-to-store"`
    SentToStoreFail int `json:"sent-to-store-fail"`
    StoreFail int `json:"store-fail"`
    InLogs int `json:"in-logs"`
    InBytes int `json:"in-bytes"`
    InLogsBacklog int `json:"in-logs-backlog"`
    InBytesBacklog int `json:"in-bytes-backlog"`
    InStoreFailNoSpace int `json:"in-store-fail-no-space"`
    InDiscardLogs int `json:"in-discard-logs"`
    InDiscardBytes int `json:"in-discard-bytes"`
    OutLogs int `json:"out-logs"`
    OutBytes int `json:"out-bytes"`
    OutError int `json:"out-error"`
    RemainingLogs int `json:"remaining-logs"`
    RemainingBytes int `json:"remaining-bytes"`
}

func (p *AcosEventsLocalLoggingStats) GetId() string{
    return "1"
}

func (p *AcosEventsLocalLoggingStats) getPath() string{
    return "acos-events/local-logging/stats"
}

func (p *AcosEventsLocalLoggingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsLocalLoggingStats,error) {
logger.Println("AcosEventsLocalLoggingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsLocalLoggingStats
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
