

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgRtspStats struct {
    
    Stats Cgnv6LsnAlgRtspStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgRtspStats struct {
    DtCgnv6LsnAlgRtspStats Cgnv6LsnAlgRtspStats `json:"rtsp"`
}


type Cgnv6LsnAlgRtspStatsStats struct {
    StreamsCreated int `json:"streams-created"`
    StreamsFreed int `json:"streams-freed"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortsAllocated int `json:"ports-allocated"`
    PortsFreed int `json:"ports-freed"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    DataSessionCreated int `json:"data-session-created"`
    DataSessionFreed int `json:"data-session-freed"`
    NoSessionMem int `json:"no-session-mem"`
}

func (p *Cgnv6LsnAlgRtspStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgRtspStats) getPath() string{
    return "cgnv6/lsn/alg/rtsp/stats"
}

func (p *Cgnv6LsnAlgRtspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgRtspStats,error) {
logger.Println("Cgnv6LsnAlgRtspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgRtspStats
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
