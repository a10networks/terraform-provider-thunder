

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgEspStats struct {
    
    Stats Cgnv6FixedNatAlgEspStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgEspStats struct {
    DtCgnv6FixedNatAlgEspStats Cgnv6FixedNatAlgEspStats `json:"esp"`
}


type Cgnv6FixedNatAlgEspStatsStats struct {
    SessionCreated int `json:"session-created"`
}

func (p *Cgnv6FixedNatAlgEspStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgEspStats) getPath() string{
    return "cgnv6/fixed-nat/alg/esp/stats"
}

func (p *Cgnv6FixedNatAlgEspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgEspStats,error) {
logger.Println("Cgnv6FixedNatAlgEspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgEspStats
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
