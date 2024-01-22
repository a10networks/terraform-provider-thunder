

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpDnsCacheOper struct {
    
    Oper SystemIpDnsCacheOperOper `json:"oper"`

}
type DataSystemIpDnsCacheOper struct {
    DtSystemIpDnsCacheOper SystemIpDnsCacheOper `json:"ip-dns-cache"`
}


type SystemIpDnsCacheOperOper struct {
    DomainIpList []SystemIpDnsCacheOperOperDomainIpList `json:"domain-ip-list"`
}


type SystemIpDnsCacheOperOperDomainIpList struct {
    Domain string `json:"domain"`
    Address string `json:"address"`
    Ttl int `json:"ttl"`
    Interval int `json:"interval"`
}

func (p *SystemIpDnsCacheOper) GetId() string{
    return "1"
}

func (p *SystemIpDnsCacheOper) getPath() string{
    return "system/ip-dns-cache/oper"
}

func (p *SystemIpDnsCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIpDnsCacheOper,error) {
logger.Println("SystemIpDnsCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIpDnsCacheOper
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
