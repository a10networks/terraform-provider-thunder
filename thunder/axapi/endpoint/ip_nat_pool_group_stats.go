

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatPoolGroupStats struct {
    
    PoolGroupName string `json:"pool-group-name"`
    Stats IpNatPoolGroupStatsStats `json:"stats"`

}
type DataIpNatPoolGroupStats struct {
    DtIpNatPoolGroupStats IpNatPoolGroupStats `json:"pool-group"`
}


type IpNatPoolGroupStatsStats struct {
    Failed int `json:"Failed"`
}

func (p *IpNatPoolGroupStats) GetId() string{
    return "1"
}

func (p *IpNatPoolGroupStats) getPath() string{
    
    return "ip/nat/pool-group/"+p.PoolGroupName+"/stats"
}

func (p *IpNatPoolGroupStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpNatPoolGroupStats,error) {
logger.Println("IpNatPoolGroupStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpNatPoolGroupStats
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
