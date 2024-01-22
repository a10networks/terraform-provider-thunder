

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TrafficControlRuleSetOper struct {
    
    Name string `json:"name"`
    Oper TrafficControlRuleSetOperOper `json:"oper"`
    RuleList []TrafficControlRuleSetOperRuleList `json:"rule-list"`

}
type DataTrafficControlRuleSetOper struct {
    DtTrafficControlRuleSetOper TrafficControlRuleSetOper `json:"rule-set"`
}


type TrafficControlRuleSetOperOper struct {
    PolicyStatus string `json:"policy-status"`
    PolicyRuleCount int `json:"policy-rule-count"`
    RuleStats []TrafficControlRuleSetOperOperRuleStats `json:"rule-stats"`
}


type TrafficControlRuleSetOperOperRuleStats struct {
    RuleName string `json:"rule-name"`
    RuleStatus string `json:"rule-status"`
    RuleHitcount int `json:"rule-hitcount"`
}


type TrafficControlRuleSetOperRuleList struct {
    Name string `json:"name"`
    Oper TrafficControlRuleSetOperRuleListOper `json:"oper"`
}


type TrafficControlRuleSetOperRuleListOper struct {
    Status string `json:"status"`
    Hitcount int `json:"hitcount"`
}

func (p *TrafficControlRuleSetOper) GetId() string{
    return "1"
}

func (p *TrafficControlRuleSetOper) getPath() string{
    
    return "traffic-control/rule-set/"+p.Name+"/oper"
}

func (p *TrafficControlRuleSetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTrafficControlRuleSetOper,error) {
logger.Println("TrafficControlRuleSetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTrafficControlRuleSetOper
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
