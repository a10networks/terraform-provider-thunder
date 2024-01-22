

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDefaultOper struct {
    
    Oper Cgnv6LsnRuleListDefaultOperOper `json:"oper"`
    Name string 

}
type DataCgnv6LsnRuleListDefaultOper struct {
    DtCgnv6LsnRuleListDefaultOper Cgnv6LsnRuleListDefaultOper `json:"default"`
}


type Cgnv6LsnRuleListDefaultOperOper struct {
    RuleList []Cgnv6LsnRuleListDefaultOperOperRuleList `json:"rule-list"`
    RuleCount int `json:"rule-count"`
}


type Cgnv6LsnRuleListDefaultOperOperRuleList struct {
    Hits int `json:"hits"`
    Proto string `json:"proto"`
    StartPort int `json:"start-port"`
    EndPort int `json:"end-port"`
    DscpMatch string `json:"dscp-match"`
    Action string `json:"action"`
    ActionType string `json:"action-type"`
    DomainName string `json:"domain-name"`
    DnatDomainUnresolvedDrops int `json:"dnat-domain-unresolved-drops"`
    Ipv4List string `json:"ipv4-list"`
    PortList string `json:"port-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Pool string `json:"pool"`
    PoolShared int `json:"pool-shared"`
    HttpAlg string `json:"http-alg"`
    TimeoutVal int `json:"timeout-val"`
    Fast int `json:"fast"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}

func (p *Cgnv6LsnRuleListDefaultOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDefaultOper) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/default/oper"
}

func (p *Cgnv6LsnRuleListDefaultOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDefaultOper,error) {
logger.Println("Cgnv6LsnRuleListDefaultOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDefaultOper
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
