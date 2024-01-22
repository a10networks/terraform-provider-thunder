

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TrafficControlRuleSetRuleOper struct {
    
    Name string `json:"name"`
    Oper TrafficControlRuleSetRuleOperOper `json:"oper"`

}
type DataTrafficControlRuleSetRuleOper struct {
    DtTrafficControlRuleSetRuleOper TrafficControlRuleSetRuleOper `json:"rule"`
}


type TrafficControlRuleSetRuleOperOper struct {
    Status string `json:"status"`
    Hitcount int `json:"hitcount"`
}

func (p *TrafficControlRuleSetRuleOper) GetId() string{
    return "1"
}

func (p *TrafficControlRuleSetRuleOper) getPath() string{
    
    return "traffic-control/rule-set/"+p.Name+"/rule/"+p.Name+"/oper"
}

func (p *TrafficControlRuleSetRuleOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTrafficControlRuleSetRuleOper,error) {
logger.Println("TrafficControlRuleSetRuleOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTrafficControlRuleSetRuleOper
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
