

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AutomaticUpdateInfoStats struct {
    
    Stats AutomaticUpdateInfoStatsStats `json:"stats"`

}
type DataAutomaticUpdateInfoStats struct {
    DtAutomaticUpdateInfoStats AutomaticUpdateInfoStats `json:"info"`
}


type AutomaticUpdateInfoStatsStats struct {
    Dummy int `json:"dummy"`
}

func (p *AutomaticUpdateInfoStats) GetId() string{
    return "1"
}

func (p *AutomaticUpdateInfoStats) getPath() string{
    return "automatic-update/info/stats"
}

func (p *AutomaticUpdateInfoStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAutomaticUpdateInfoStats,error) {
logger.Println("AutomaticUpdateInfoStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAutomaticUpdateInfoStats
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
