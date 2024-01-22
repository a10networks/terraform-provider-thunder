

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPerformanceStats struct {
    
    Stats Cgnv6LsnPerformanceStatsStats `json:"stats"`

}
type DataCgnv6LsnPerformanceStats struct {
    DtCgnv6LsnPerformanceStats Cgnv6LsnPerformanceStats `json:"performance"`
}


type Cgnv6LsnPerformanceStatsStats struct {
    DataSessionsCurrentEpoch int `json:"data-sessions-current-epoch"`
    FullconeCreatedCurrentEpoch int `json:"fullcone-created-current-epoch"`
    UserQuoteCreatedCurrentEpoch int `json:"user-quote-created-current-epoch"`
    DataSessionsPreviousEpochFirst int `json:"data-sessions-previous-epoch-first"`
    FullconeCreatedPreviousEpochFirst int `json:"fullcone-created-previous-epoch-first"`
    UserQuoteCreatedPreviousEpochFirst int `json:"user-quote-created-previous-epoch-first"`
    DataSessionsPreviousEpochLast int `json:"data-sessions-previous-epoch-last"`
    FullconeCreatedPreviousEpochLast int `json:"fullcone-created-previous-epoch-last"`
    UserQuoteCreatedPreviousEpochLast int `json:"user-quote-created-previous-epoch-last"`
}

func (p *Cgnv6LsnPerformanceStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnPerformanceStats) getPath() string{
    return "cgnv6/lsn/performance/stats"
}

func (p *Cgnv6LsnPerformanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnPerformanceStats,error) {
logger.Println("Cgnv6LsnPerformanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnPerformanceStats
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
