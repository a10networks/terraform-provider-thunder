

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatPoolStats struct {
    
    PoolName string `json:"pool-name"`
    Stats Ipv6NatPoolStatsStats `json:"stats"`

}
type DataIpv6NatPoolStats struct {
    DtIpv6NatPoolStats Ipv6NatPoolStats `json:"pool"`
}


type Ipv6NatPoolStatsStats struct {
    PortUsage int `json:"Port-Usage"`
    TotalUsed int `json:"Total-Used"`
    TotalFreed int `json:"Total-Freed"`
    Failed int `json:"Failed"`
}

func (p *Ipv6NatPoolStats) GetId() string{
    return "1"
}

func (p *Ipv6NatPoolStats) getPath() string{
    
    return "ipv6/nat/pool/"+p.PoolName+"/stats"
}

func (p *Ipv6NatPoolStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6NatPoolStats,error) {
logger.Println("Ipv6NatPoolStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6NatPoolStats
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
