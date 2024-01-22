

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbRpzStats struct {
    
    Stats SlbRpzStatsStats `json:"stats"`

}
type DataSlbRpzStats struct {
    DtSlbRpzStats SlbRpzStats `json:"rpz"`
}


type SlbRpzStatsStats struct {
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
}

func (p *SlbRpzStats) GetId() string{
    return "1"
}

func (p *SlbRpzStats) getPath() string{
    return "slb/rpz/stats"
}

func (p *SlbRpzStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbRpzStats,error) {
logger.Println("SlbRpzStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbRpzStats
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
