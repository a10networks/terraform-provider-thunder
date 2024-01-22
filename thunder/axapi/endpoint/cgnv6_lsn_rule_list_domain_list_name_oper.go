

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainListNameOper struct {
    
    NameDomainList string `json:"name-domain-list"`
    Oper Cgnv6LsnRuleListDomainListNameOperOper `json:"oper"`
    Name string 

}
type DataCgnv6LsnRuleListDomainListNameOper struct {
    DtCgnv6LsnRuleListDomainListNameOper Cgnv6LsnRuleListDomainListNameOper `json:"domain-list-name"`
}


type Cgnv6LsnRuleListDomainListNameOperOper struct {
    RuleList []Cgnv6LsnRuleListDomainListNameOperOperRuleList `json:"rule-list"`
    RuleCount int `json:"rule-count"`
}


type Cgnv6LsnRuleListDomainListNameOperOperRuleList struct {
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

func (p *Cgnv6LsnRuleListDomainListNameOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainListNameOper) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-list-name/"+p.NameDomainList+"/oper"
}

func (p *Cgnv6LsnRuleListDomainListNameOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainListNameOper,error) {
logger.Println("Cgnv6LsnRuleListDomainListNameOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainListNameOper
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
