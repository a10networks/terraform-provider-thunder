

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRuleOper struct {
    
    Name string `json:"name"`
    Oper RuleSetRuleOperOper `json:"oper"`

}
type DataRuleSetRuleOper struct {
    DtRuleSetRuleOper RuleSetRuleOper `json:"rule"`
}


type RuleSetRuleOperOper struct {
    Hitcount int `json:"hitcount"`
    LastHitcountTime string `json:"last-hitcount-time"`
    Action string `json:"action"`
    Status string `json:"status"`
    Permitbytes int `json:"permitbytes"`
    Denybytes int `json:"denybytes"`
    Resetbytes int `json:"resetbytes"`
    Totalbytes int `json:"totalbytes"`
    Permitpackets int `json:"permitpackets"`
    Denypackets int `json:"denypackets"`
    Resetpackets int `json:"resetpackets"`
    Totalpackets int `json:"totalpackets"`
    Activesessiontcp int `json:"activesessiontcp"`
    Activesessionudp int `json:"activesessionudp"`
    Activesessionicmp int `json:"activesessionicmp"`
    Activesessionsctp int `json:"activesessionsctp"`
    Activesessionother int `json:"activesessionother"`
    Activesessiontotal int `json:"activesessiontotal"`
    Sessiontcp int `json:"sessiontcp"`
    Sessionudp int `json:"sessionudp"`
    Sessionicmp int `json:"sessionicmp"`
    Sessionsctp int `json:"sessionsctp"`
    Sessionother int `json:"sessionother"`
    Sessiontotal int `json:"sessiontotal"`
    Ratelimitdrops int `json:"ratelimitdrops"`
}

func (p *RuleSetRuleOper) GetId() string{
    return "1"
}

func (p *RuleSetRuleOper) getPath() string{
    
    return "rule-set/"+p.Name+"/rule/"+p.Name+"/oper"
}

func (p *RuleSetRuleOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetRuleOper,error) {
logger.Println("RuleSetRuleOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetRuleOper
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
