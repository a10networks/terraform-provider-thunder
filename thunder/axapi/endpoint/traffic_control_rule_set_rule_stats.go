

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TrafficControlRuleSetRuleStats struct {
    
    Name string `json:"name"`
    Stats TrafficControlRuleSetRuleStatsStats `json:"stats"`

}
type DataTrafficControlRuleSetRuleStats struct {
    DtTrafficControlRuleSetRuleStats TrafficControlRuleSetRuleStats `json:"rule"`
}


type TrafficControlRuleSetRuleStatsStats struct {
    HitCount int `json:"hit-count"`
}

func (p *TrafficControlRuleSetRuleStats) GetId() string{
    return "1"
}

func (p *TrafficControlRuleSetRuleStats) getPath() string{
    
    return "traffic-control/rule-set/"+p.Name+"/rule/"+p.Name+"/stats"
}

func (p *TrafficControlRuleSetRuleStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTrafficControlRuleSetRuleStats,error) {
logger.Println("TrafficControlRuleSetRuleStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTrafficControlRuleSetRuleStats
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
