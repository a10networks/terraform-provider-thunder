

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgRtspStats struct {
    
    Stats FwAlgRtspStatsStats `json:"stats"`

}
type DataFwAlgRtspStats struct {
    DtFwAlgRtspStats FwAlgRtspStats `json:"rtsp"`
}


type FwAlgRtspStatsStats struct {
    TransportInserted int `json:"transport-inserted"`
    TransportFreed int `json:"transport-freed"`
    TransportAllocFailure int `json:"transport-alloc-failure"`
    DataSessionCreated int `json:"data-session-created"`
    DataSessionFreed int `json:"data-session-freed"`
}

func (p *FwAlgRtspStats) GetId() string{
    return "1"
}

func (p *FwAlgRtspStats) getPath() string{
    return "fw/alg/rtsp/stats"
}

func (p *FwAlgRtspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgRtspStats,error) {
logger.Println("FwAlgRtspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgRtspStats
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
