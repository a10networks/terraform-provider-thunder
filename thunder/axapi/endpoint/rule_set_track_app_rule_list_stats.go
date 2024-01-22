

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetTrackAppRuleListStats struct {
    
    Stats RuleSetTrackAppRuleListStatsStats `json:"stats"`
    Name string 

}
type DataRuleSetTrackAppRuleListStats struct {
    DtRuleSetTrackAppRuleListStats RuleSetTrackAppRuleListStats `json:"track-app-rule-list"`
}


type RuleSetTrackAppRuleListStatsStats struct {
    Dummy int `json:"dummy"`
}

func (p *RuleSetTrackAppRuleListStats) GetId() string{
    return "1"
}

func (p *RuleSetTrackAppRuleListStats) getPath() string{
    
    return "rule-set/" +p.Name + "/track-app-rule-list/stats"
}

func (p *RuleSetTrackAppRuleListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetTrackAppRuleListStats,error) {
logger.Println("RuleSetTrackAppRuleListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetTrackAppRuleListStats
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
