

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAppStats struct {
    
    Stats FwAppStatsStats `json:"stats"`

}
type DataFwAppStats struct {
    DtFwAppStats FwAppStats `json:"app"`
}


type FwAppStatsStats struct {
    Dummy int `json:"dummy"`
}

func (p *FwAppStats) GetId() string{
    return "1"
}

func (p *FwAppStats) getPath() string{
    return "fw/app/stats"
}

func (p *FwAppStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAppStats,error) {
logger.Println("FwAppStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAppStats
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
