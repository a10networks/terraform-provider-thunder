

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetTrackAppRuleListOper struct {
    
    Oper RuleSetTrackAppRuleListOperOper `json:"oper"`
    Name string 

}
type DataRuleSetTrackAppRuleListOper struct {
    DtRuleSetTrackAppRuleListOper RuleSetTrackAppRuleListOper `json:"track-app-rule-list"`
}


type RuleSetTrackAppRuleListOperOper struct {
    RuleList []RuleSetTrackAppRuleListOperOperRuleList `json:"rule-list"`
}


type RuleSetTrackAppRuleListOperOperRuleList struct {
    Name string `json:"name"`
}

func (p *RuleSetTrackAppRuleListOper) GetId() string{
    return "1"
}

func (p *RuleSetTrackAppRuleListOper) getPath() string{
    
    return "rule-set/" +p.Name + "/track-app-rule-list/oper"
}

func (p *RuleSetTrackAppRuleListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetTrackAppRuleListOper,error) {
logger.Println("RuleSetTrackAppRuleListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetTrackAppRuleListOper
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
