

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetApplicationOper struct {
    
    Oper RuleSetApplicationOperOper `json:"oper"`
    Name string 

}
type DataRuleSetApplicationOper struct {
    DtRuleSetApplicationOper RuleSetApplicationOper `json:"application"`
}


type RuleSetApplicationOperOper struct {
    CategoryStat string `json:"category-stat"`
    AppStat string `json:"app-stat"`
    Protocol int `json:"protocol"`
    Rule string `json:"rule"`
    RuleSetOnly int `json:"rule-set-only"`
    RuleList []RuleSetApplicationOperOperRuleList `json:"rule-list"`
}


type RuleSetApplicationOperOperRuleList struct {
    Name string `json:"name"`
    StatList []RuleSetApplicationOperOperRuleListStatList `json:"stat-list"`
}


type RuleSetApplicationOperOperRuleListStatList struct {
    Name string `json:"name"`
    Category string `json:"category"`
    Type string `json:"type"`
    Conns int `json:"conns"`
    Bytes int `json:"bytes"`
}

func (p *RuleSetApplicationOper) GetId() string{
    return "1"
}

func (p *RuleSetApplicationOper) getPath() string{
    
    return "rule-set/" +p.Name + "/application/oper"
}

func (p *RuleSetApplicationOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetApplicationOper,error) {
logger.Println("RuleSetApplicationOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetApplicationOper
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
