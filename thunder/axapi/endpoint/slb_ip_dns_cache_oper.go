

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbIpDnsCacheOper struct {
    
    Oper SlbIpDnsCacheOperOper `json:"oper"`

}
type DataSlbIpDnsCacheOper struct {
    DtSlbIpDnsCacheOper SlbIpDnsCacheOper `json:"ip-dns-cache"`
}


type SlbIpDnsCacheOperOper struct {
    DomainIpList []SlbIpDnsCacheOperOperDomainIpList `json:"domain-ip-list"`
}


type SlbIpDnsCacheOperOperDomainIpList struct {
    Domain string `json:"domain"`
    Address string `json:"address"`
    Ttl int `json:"ttl"`
    Interval int `json:"interval"`
}

func (p *SlbIpDnsCacheOper) GetId() string{
    return "1"
}

func (p *SlbIpDnsCacheOper) getPath() string{
    return "slb/ip-dns-cache/oper"
}

func (p *SlbIpDnsCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbIpDnsCacheOper,error) {
logger.Println("SlbIpDnsCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbIpDnsCacheOper
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
