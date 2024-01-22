

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatPoolStats struct {
    
    PoolName string `json:"pool-name"`
    Stats IpNatPoolStatsStats `json:"stats"`

}
type DataIpNatPoolStats struct {
    DtIpNatPoolStats IpNatPoolStats `json:"pool"`
}


type IpNatPoolStatsStats struct {
    PortUsage int `json:"Port-Usage"`
    TotalUsed int `json:"Total-Used"`
    TotalFreed int `json:"Total-Freed"`
    Failed int `json:"Failed"`
}

func (p *IpNatPoolStats) GetId() string{
    return "1"
}

func (p *IpNatPoolStats) getPath() string{
    
    return "ip/nat/pool/"+p.PoolName+"/stats"
}

func (p *IpNatPoolStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpNatPoolStats,error) {
logger.Println("IpNatPoolStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpNatPoolStats
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
