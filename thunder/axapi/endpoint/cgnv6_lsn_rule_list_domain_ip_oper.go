

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainIpOper struct {
    
    Oper Cgnv6LsnRuleListDomainIpOperOper `json:"oper"`
    Name string 

}
type DataCgnv6LsnRuleListDomainIpOper struct {
    DtCgnv6LsnRuleListDomainIpOper Cgnv6LsnRuleListDomainIpOper `json:"domain-ip"`
}


type Cgnv6LsnRuleListDomainIpOperOper struct {
    IpList []Cgnv6LsnRuleListDomainIpOperOperIpList `json:"ip-list"`
}


type Cgnv6LsnRuleListDomainIpOperOperIpList struct {
    Domain string `json:"domain"`
    DomainList string `json:"domain-list"`
    IpAddress string `json:"ip-address"`
    Ttl int `json:"ttl"`
}

func (p *Cgnv6LsnRuleListDomainIpOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainIpOper) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-ip/oper"
}

func (p *Cgnv6LsnRuleListDomainIpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainIpOper,error) {
logger.Println("Cgnv6LsnRuleListDomainIpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainIpOper
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
