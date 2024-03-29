

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainNameOper struct {
    
    NameDomain string `json:"name-domain"`
    Oper Cgnv6LsnRuleListDomainNameOperOper `json:"oper"`
    Name string 

}
type DataCgnv6LsnRuleListDomainNameOper struct {
    DtCgnv6LsnRuleListDomainNameOper Cgnv6LsnRuleListDomainNameOper `json:"domain-name"`
}


type Cgnv6LsnRuleListDomainNameOperOper struct {
    RuleList []Cgnv6LsnRuleListDomainNameOperOperRuleList `json:"rule-list"`
    RuleCount int `json:"rule-count"`
}


type Cgnv6LsnRuleListDomainNameOperOperRuleList struct {
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

func (p *Cgnv6LsnRuleListDomainNameOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainNameOper) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-name/"+p.NameDomain+"/oper"
}

func (p *Cgnv6LsnRuleListDomainNameOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainNameOper,error) {
logger.Println("Cgnv6LsnRuleListDomainNameOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainNameOper
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
