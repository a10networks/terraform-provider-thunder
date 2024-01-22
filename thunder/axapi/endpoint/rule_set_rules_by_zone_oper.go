

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRulesByZoneOper struct {
    
    Oper RuleSetRulesByZoneOperOper `json:"oper"`
    Name string 

}
type DataRuleSetRulesByZoneOper struct {
    DtRuleSetRulesByZoneOper RuleSetRulesByZoneOper `json:"rules-by-zone"`
}


type RuleSetRulesByZoneOperOper struct {
    GroupList []RuleSetRulesByZoneOperOperGroupList `json:"group-list"`
}


type RuleSetRulesByZoneOperOperGroupList struct {
    From string `json:"from"`
    To string `json:"to"`
    RuleList []RuleSetRulesByZoneOperOperGroupListRuleList `json:"rule-list"`
}


type RuleSetRulesByZoneOperOperGroupListRuleList struct {
    Name string `json:"name"`
    Action string `json:"action"`
    SourceList []RuleSetRulesByZoneOperOperGroupListRuleListSourceList `json:"source-list"`
    DestList []RuleSetRulesByZoneOperOperGroupListRuleListDestList `json:"dest-list"`
    ServiceList []RuleSetRulesByZoneOperOperGroupListRuleListServiceList `json:"service-list"`
    DscpList []RuleSetRulesByZoneOperOperGroupListRuleListDscpList `json:"dscp-list"`
}


type RuleSetRulesByZoneOperOperGroupListRuleListSourceList struct {
    Source string `json:"source"`
}


type RuleSetRulesByZoneOperOperGroupListRuleListDestList struct {
    Dest string `json:"dest"`
}


type RuleSetRulesByZoneOperOperGroupListRuleListServiceList struct {
    Service string `json:"service"`
}


type RuleSetRulesByZoneOperOperGroupListRuleListDscpList struct {
    Dscp string `json:"dscp"`
}

func (p *RuleSetRulesByZoneOper) GetId() string{
    return "1"
}

func (p *RuleSetRulesByZoneOper) getPath() string{
    
    return "rule-set/" +p.Name + "/rules-by-zone/oper"
}

func (p *RuleSetRulesByZoneOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetRulesByZoneOper,error) {
logger.Println("RuleSetRulesByZoneOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetRulesByZoneOper
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
