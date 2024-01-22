

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgRtspStats struct {
    
    Stats Cgnv6FixedNatAlgRtspStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgRtspStats struct {
    DtCgnv6FixedNatAlgRtspStats Cgnv6FixedNatAlgRtspStats `json:"rtsp"`
}


type Cgnv6FixedNatAlgRtspStatsStats struct {
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

func (p *Cgnv6FixedNatAlgRtspStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgRtspStats) getPath() string{
    return "cgnv6/fixed-nat/alg/rtsp/stats"
}

func (p *Cgnv6FixedNatAlgRtspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgRtspStats,error) {
logger.Println("Cgnv6FixedNatAlgRtspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgRtspStats
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
