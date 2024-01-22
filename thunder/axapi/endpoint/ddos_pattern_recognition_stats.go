

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosPatternRecognitionStats struct {
    
    Stats DdosPatternRecognitionStatsStats `json:"stats"`

}
type DataDdosPatternRecognitionStats struct {
    DtDdosPatternRecognitionStats DdosPatternRecognitionStats `json:"pattern-recognition"`
}


type DdosPatternRecognitionStatsStats struct {
    Proceeded int `json:"proceeded"`
    Not_found int `json:"not_found"`
    Generic_error int `json:"generic_error"`
    Filter_match int `json:"filter_match"`
    Filter_drop int `json:"filter_drop"`
}

func (p *DdosPatternRecognitionStats) GetId() string{
    return "1"
}

func (p *DdosPatternRecognitionStats) getPath() string{
    return "ddos/pattern-recognition/stats"
}

func (p *DdosPatternRecognitionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosPatternRecognitionStats,error) {
logger.Println("DdosPatternRecognitionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosPatternRecognitionStats
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
