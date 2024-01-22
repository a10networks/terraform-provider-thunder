

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSportRateLimitStats struct {
    
    Stats SlbSportRateLimitStatsStats `json:"stats"`

}
type DataSlbSportRateLimitStats struct {
    DtSlbSportRateLimitStats SlbSportRateLimitStats `json:"sport-rate-limit"`
}


type SlbSportRateLimitStatsStats struct {
    Alloc_sport int `json:"alloc_sport"`
    Alloc_sportip int `json:"alloc_sportip"`
    Freed_sport int `json:"freed_sport"`
    Freed_sportip int `json:"freed_sportip"`
    Total_drop int `json:"total_drop"`
    Total_reset int `json:"total_reset"`
    Total_log int `json:"total_log"`
}

func (p *SlbSportRateLimitStats) GetId() string{
    return "1"
}

func (p *SlbSportRateLimitStats) getPath() string{
    return "slb/sport-rate-limit/stats"
}

func (p *SlbSportRateLimitStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSportRateLimitStats,error) {
logger.Println("SlbSportRateLimitStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSportRateLimitStats
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
