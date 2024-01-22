

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatPoolGroupStats struct {
    
    PoolGroupName string `json:"pool-group-name"`
    Stats Ipv6NatPoolGroupStatsStats `json:"stats"`

}
type DataIpv6NatPoolGroupStats struct {
    DtIpv6NatPoolGroupStats Ipv6NatPoolGroupStats `json:"pool-group"`
}


type Ipv6NatPoolGroupStatsStats struct {
    Failed int `json:"Failed"`
}

func (p *Ipv6NatPoolGroupStats) GetId() string{
    return "1"
}

func (p *Ipv6NatPoolGroupStats) getPath() string{
    
    return "ipv6/nat/pool-group/"+p.PoolGroupName+"/stats"
}

func (p *Ipv6NatPoolGroupStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6NatPoolGroupStats,error) {
logger.Println("Ipv6NatPoolGroupStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6NatPoolGroupStats
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
