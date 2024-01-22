

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureAutomatedCapturesStats struct {
    
    Stats VisibilityPacketCaptureAutomatedCapturesStatsStats `json:"stats"`

}
type DataVisibilityPacketCaptureAutomatedCapturesStats struct {
    DtVisibilityPacketCaptureAutomatedCapturesStats VisibilityPacketCaptureAutomatedCapturesStats `json:"automated-captures"`
}


type VisibilityPacketCaptureAutomatedCapturesStatsStats struct {
    TotalFailure int `json:"total-failure"`
}

func (p *VisibilityPacketCaptureAutomatedCapturesStats) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureAutomatedCapturesStats) getPath() string{
    return "visibility/packet-capture/automated-captures/stats"
}

func (p *VisibilityPacketCaptureAutomatedCapturesStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPacketCaptureAutomatedCapturesStats,error) {
logger.Println("VisibilityPacketCaptureAutomatedCapturesStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPacketCaptureAutomatedCapturesStats
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
