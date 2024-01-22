

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDefaultStats struct {
    
    Stats Cgnv6LsnRuleListDefaultStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6LsnRuleListDefaultStats struct {
    DtCgnv6LsnRuleListDefaultStats Cgnv6LsnRuleListDefaultStats `json:"default"`
}


type Cgnv6LsnRuleListDefaultStatsStats struct {
}

func (p *Cgnv6LsnRuleListDefaultStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDefaultStats) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/default/stats"
}

func (p *Cgnv6LsnRuleListDefaultStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDefaultStats,error) {
logger.Println("Cgnv6LsnRuleListDefaultStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDefaultStats
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
