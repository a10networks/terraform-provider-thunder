

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOnePoolStats struct {
    
    PoolName string `json:"pool-name"`
    Stats Cgnv6OneToOnePoolStatsStats `json:"stats"`

}
type DataCgnv6OneToOnePoolStats struct {
    DtCgnv6OneToOnePoolStats Cgnv6OneToOnePoolStats `json:"pool"`
}


type Cgnv6OneToOnePoolStatsStats struct {
    TotalAddress int `json:"total-address"`
    UsedAddress int `json:"used-address"`
    FreeAddress int `json:"free-address"`
}

func (p *Cgnv6OneToOnePoolStats) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOnePoolStats) getPath() string{
    
    return "cgnv6/one-to-one/pool/"+p.PoolName+"/stats"
}

func (p *Cgnv6OneToOnePoolStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOnePoolStats,error) {
logger.Println("Cgnv6OneToOnePoolStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOnePoolStats
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
