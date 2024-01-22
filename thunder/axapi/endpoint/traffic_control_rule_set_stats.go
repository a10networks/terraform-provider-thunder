

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TrafficControlRuleSetStats struct {
    
    Name string `json:"name"`
    RuleList []TrafficControlRuleSetStatsRuleList `json:"rule-list"`
    Stats TrafficControlRuleSetStatsStats `json:"stats"`

}
type DataTrafficControlRuleSetStats struct {
    DtTrafficControlRuleSetStats TrafficControlRuleSetStats `json:"rule-set"`
}


type TrafficControlRuleSetStatsRuleList struct {
    Name string `json:"name"`
    Stats TrafficControlRuleSetStatsRuleListStats `json:"stats"`
}


type TrafficControlRuleSetStatsRuleListStats struct {
    HitCount int `json:"hit-count"`
}


type TrafficControlRuleSetStatsStats struct {
    HitCount int `json:"hit-count"`
}

func (p *TrafficControlRuleSetStats) GetId() string{
    return "1"
}

func (p *TrafficControlRuleSetStats) getPath() string{
    
    return "traffic-control/rule-set/"+p.Name+"/stats"
}

func (p *TrafficControlRuleSetStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTrafficControlRuleSetStats,error) {
logger.Println("TrafficControlRuleSetStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTrafficControlRuleSetStats
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
