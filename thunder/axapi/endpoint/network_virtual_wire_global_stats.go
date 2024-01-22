

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireGlobalStats struct {
    
    Stats NetworkVirtualWireGlobalStatsStats `json:"stats"`

}
type DataNetworkVirtualWireGlobalStats struct {
    DtNetworkVirtualWireGlobalStats NetworkVirtualWireGlobalStats `json:"virtual-wire-global"`
}


type NetworkVirtualWireGlobalStatsStats struct {
    VlanUpdate int `json:"vlan-update"`
    MacUpdate int `json:"mac-update"`
    VlanPairUpdate int `json:"vlan-pair-update"`
}

func (p *NetworkVirtualWireGlobalStats) GetId() string{
    return "1"
}

func (p *NetworkVirtualWireGlobalStats) getPath() string{
    return "network/virtual-wire-global/stats"
}

func (p *NetworkVirtualWireGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireGlobalStats,error) {
logger.Println("NetworkVirtualWireGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVirtualWireGlobalStats
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
