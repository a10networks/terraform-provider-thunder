

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDomainListInfoOper struct {
    
    Oper SlbDomainListInfoOperOper `json:"oper"`

}
type DataSlbDomainListInfoOper struct {
    DtSlbDomainListInfoOper SlbDomainListInfoOper `json:"domain-list-info"`
}


type SlbDomainListInfoOperOper struct {
    DomainList []SlbDomainListInfoOperOperDomainList `json:"domain-list"`
}


type SlbDomainListInfoOperOperDomainList struct {
    Name string `json:"name"`
    DomainName string `json:"domain-name"`
    Resolved int `json:"resolved"`
    HitCount int `json:"hit-count"`
    IpNumber int `json:"ip-number"`
    Ip string `json:"ip"`
    IsIpv6 int `json:"is-ipv6"`
}

func (p *SlbDomainListInfoOper) GetId() string{
    return "1"
}

func (p *SlbDomainListInfoOper) getPath() string{
    return "slb/domain-list-info/oper"
}

func (p *SlbDomainListInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDomainListInfoOper,error) {
logger.Println("SlbDomainListInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDomainListInfoOper
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
