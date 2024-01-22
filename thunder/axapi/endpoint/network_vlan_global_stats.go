

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanGlobalStats struct {
    
    Stats NetworkVlanGlobalStatsStats `json:"stats"`

}
type DataNetworkVlanGlobalStats struct {
    DtNetworkVlanGlobalStats NetworkVlanGlobalStats `json:"vlan-global"`
}


type NetworkVlanGlobalStatsStats struct {
    Xparent_vlan_list_err int `json:"xparent_vlan_list_err"`
}

func (p *NetworkVlanGlobalStats) GetId() string{
    return "1"
}

func (p *NetworkVlanGlobalStats) getPath() string{
    return "network/vlan-global/stats"
}

func (p *NetworkVlanGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVlanGlobalStats,error) {
logger.Println("NetworkVlanGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVlanGlobalStats
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
