

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwTcpWindowCheckStats struct {
    
    Stats FwTcpWindowCheckStatsStats `json:"stats"`

}
type DataFwTcpWindowCheckStats struct {
    DtFwTcpWindowCheckStats FwTcpWindowCheckStats `json:"tcp-window-check"`
}


type FwTcpWindowCheckStatsStats struct {
    OutsideWindow int `json:"outside-window"`
}

func (p *FwTcpWindowCheckStats) GetId() string{
    return "1"
}

func (p *FwTcpWindowCheckStats) getPath() string{
    return "fw/tcp-window-check/stats"
}

func (p *FwTcpWindowCheckStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwTcpWindowCheckStats,error) {
logger.Println("FwTcpWindowCheckStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwTcpWindowCheckStats
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
