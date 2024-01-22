

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallAlgRtspStats struct {
    
    Stats Cgnv6StatefulFirewallAlgRtspStatsStats `json:"stats"`

}
type DataCgnv6StatefulFirewallAlgRtspStats struct {
    DtCgnv6StatefulFirewallAlgRtspStats Cgnv6StatefulFirewallAlgRtspStats `json:"rtsp"`
}


type Cgnv6StatefulFirewallAlgRtspStatsStats struct {
    TransportInserted int `json:"transport-inserted"`
    TransportFreed int `json:"transport-freed"`
    TransportAllocFailure int `json:"transport-alloc-failure"`
    DataSessionCreated int `json:"data-session-created"`
    DataSessionFreed int `json:"data-session-freed"`
}

func (p *Cgnv6StatefulFirewallAlgRtspStats) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallAlgRtspStats) getPath() string{
    return "cgnv6/stateful-firewall/alg/rtsp/stats"
}

func (p *Cgnv6StatefulFirewallAlgRtspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6StatefulFirewallAlgRtspStats,error) {
logger.Println("Cgnv6StatefulFirewallAlgRtspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6StatefulFirewallAlgRtspStats
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
