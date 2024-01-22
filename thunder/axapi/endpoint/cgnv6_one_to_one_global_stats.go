

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneGlobalStats struct {
    
    Stats Cgnv6OneToOneGlobalStatsStats `json:"stats"`

}
type DataCgnv6OneToOneGlobalStats struct {
    DtCgnv6OneToOneGlobalStats Cgnv6OneToOneGlobalStats `json:"global"`
}


type Cgnv6OneToOneGlobalStatsStats struct {
    TotalMapAllocated int `json:"total-map-allocated"`
    TotalMapFreed int `json:"total-map-freed"`
    MapAllocFailure int `json:"map-alloc-failure"`
}

func (p *Cgnv6OneToOneGlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneGlobalStats) getPath() string{
    return "cgnv6/one-to-one/global/stats"
}

func (p *Cgnv6OneToOneGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOneGlobalStats,error) {
logger.Println("Cgnv6OneToOneGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOneGlobalStats
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
