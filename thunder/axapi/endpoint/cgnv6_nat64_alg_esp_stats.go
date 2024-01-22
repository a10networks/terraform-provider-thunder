

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64AlgEspStats struct {
    
    Stats Cgnv6Nat64AlgEspStatsStats `json:"stats"`

}
type DataCgnv6Nat64AlgEspStats struct {
    DtCgnv6Nat64AlgEspStats Cgnv6Nat64AlgEspStats `json:"esp"`
}


type Cgnv6Nat64AlgEspStatsStats struct {
    SessionCreated int `json:"session-created"`
    NatIpConflict int `json:"nat-ip-conflict"`
}

func (p *Cgnv6Nat64AlgEspStats) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64AlgEspStats) getPath() string{
    return "cgnv6/nat64/alg/esp/stats"
}

func (p *Cgnv6Nat64AlgEspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat64AlgEspStats,error) {
logger.Println("Cgnv6Nat64AlgEspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat64AlgEspStats
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
