

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgEspStats struct {
    
    Stats Cgnv6LsnAlgEspStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgEspStats struct {
    DtCgnv6LsnAlgEspStats Cgnv6LsnAlgEspStats `json:"esp"`
}


type Cgnv6LsnAlgEspStatsStats struct {
    SessionCreated int `json:"session-created"`
    NatIpConflict int `json:"nat-ip-conflict"`
}

func (p *Cgnv6LsnAlgEspStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgEspStats) getPath() string{
    return "cgnv6/lsn/alg/esp/stats"
}

func (p *Cgnv6LsnAlgEspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgEspStats,error) {
logger.Println("Cgnv6LsnAlgEspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgEspStats
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
