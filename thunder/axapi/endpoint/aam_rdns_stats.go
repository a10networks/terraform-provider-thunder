

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamRdnsStats struct {
    
    Stats AamRdnsStatsStats `json:"stats"`

}
type DataAamRdnsStats struct {
    DtAamRdnsStats AamRdnsStats `json:"rdns"`
}


type AamRdnsStatsStats struct {
    RequestNormal int `json:"request-normal"`
    RequestDropped int `json:"request-dropped"`
    ResponseSuccess int `json:"response-success"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
}

func (p *AamRdnsStats) GetId() string{
    return "1"
}

func (p *AamRdnsStats) getPath() string{
    return "aam/rdns/stats"
}

func (p *AamRdnsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamRdnsStats,error) {
logger.Println("AamRdnsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamRdnsStats
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
