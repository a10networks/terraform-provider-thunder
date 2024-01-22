

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReportingStats struct {
    
    Stats VisibilityReportingStatsStats `json:"stats"`

}
type DataVisibilityReportingStats struct {
    DtVisibilityReportingStats VisibilityReportingStats `json:"reporting"`
}


type VisibilityReportingStatsStats struct {
    LogTransmitFailure int `json:"log-transmit-failure"`
    BufferAllocFailure int `json:"buffer-alloc-failure"`
    NotifJobsInQueue int `json:"notif-jobs-in-queue"`
    EnqueueFail int `json:"enqueue-fail"`
    EnqueuePass int `json:"enqueue-pass"`
    Dequeued int `json:"dequeued"`
}

func (p *VisibilityReportingStats) GetId() string{
    return "1"
}

func (p *VisibilityReportingStats) getPath() string{
    return "visibility/reporting/stats"
}

func (p *VisibilityReportingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityReportingStats,error) {
logger.Println("VisibilityReportingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityReportingStats
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
