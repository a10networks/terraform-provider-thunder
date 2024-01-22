

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgEspStats struct {
    
    Stats FwAlgEspStatsStats `json:"stats"`

}
type DataFwAlgEspStats struct {
    DtFwAlgEspStats FwAlgEspStats `json:"esp"`
}


type FwAlgEspStatsStats struct {
    SessionCreated int `json:"session-created"`
    HelperCreated int `json:"helper-created"`
    HelperFreed int `json:"helper-freed"`
}

func (p *FwAlgEspStats) GetId() string{
    return "1"
}

func (p *FwAlgEspStats) getPath() string{
    return "fw/alg/esp/stats"
}

func (p *FwAlgEspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgEspStats,error) {
logger.Println("FwAlgEspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgEspStats
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
