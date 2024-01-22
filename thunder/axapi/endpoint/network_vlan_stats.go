

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanStats struct {
    
    Stats NetworkVlanStatsStats `json:"stats"`
    VlanNum int `json:"vlan-num"`

}
type DataNetworkVlanStats struct {
    DtNetworkVlanStats NetworkVlanStats `json:"vlan"`
}


type NetworkVlanStatsStats struct {
    Broadcast_count int `json:"broadcast_count"`
    Multicast_count int `json:"multicast_count"`
    Ip_multicast_count int `json:"ip_multicast_count"`
    Unknown_unicast_count int `json:"unknown_unicast_count"`
    Mac_movement_count int `json:"mac_movement_count"`
    Shared_vlan_partition_switched_counter int `json:"shared_vlan_partition_switched_counter"`
}

func (p *NetworkVlanStats) GetId() string{
    return "1"
}

func (p *NetworkVlanStats) getPath() string{
    
    return "network/vlan/" +strconv.Itoa(p.VlanNum)+"/stats"
}

func (p *NetworkVlanStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVlanStats,error) {
logger.Println("NetworkVlanStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVlanStats
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
