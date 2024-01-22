

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDomainListInfoOper struct {
    
    Oper SystemDomainListInfoOperOper `json:"oper"`

}
type DataSystemDomainListInfoOper struct {
    DtSystemDomainListInfoOper SystemDomainListInfoOper `json:"domain-list-info"`
}


type SystemDomainListInfoOperOper struct {
    DomainList []SystemDomainListInfoOperOperDomainList `json:"domain-list"`
}


type SystemDomainListInfoOperOperDomainList struct {
    Name string `json:"name"`
    DomainName string `json:"domain-name"`
    Resolved int `json:"resolved"`
    HitCount int `json:"hit-count"`
    IpNumber int `json:"ip-number"`
    Ip string `json:"ip"`
    IsIpv6 int `json:"is-ipv6"`
}

func (p *SystemDomainListInfoOper) GetId() string{
    return "1"
}

func (p *SystemDomainListInfoOper) getPath() string{
    return "system/domain-list-info/oper"
}

func (p *SystemDomainListInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemDomainListInfoOper,error) {
logger.Println("SystemDomainListInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemDomainListInfoOper
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
