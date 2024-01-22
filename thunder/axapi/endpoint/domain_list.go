

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DomainList struct {
	Inst struct {

    DomainNameList []DomainListDomainNameList `json:"domain-name-list"`
    File int `json:"file"`
    MatchTypeAxfr []DomainListMatchTypeAxfr `json:"match-type-axfr"`
    MatchTypeEquals []DomainListMatchTypeEquals `json:"match-type-equals"`
    MatchTypeSuffix []DomainListMatchTypeSuffix `json:"match-type-suffix"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"domain-list"`
}


type DomainListDomainNameList struct {
    DomainName string `json:"domain-name"`
    Interval int `json:"interval" dval:"10"`
}


type DomainListMatchTypeAxfr struct {
    AxfrDomain string `json:"axfr-domain"`
    AxfrIpAddress string `json:"axfr-ip-address"`
    AxfrIpv6Address string `json:"axfr-ipv6-address"`
    IpAxfrPortNum int `json:"ip-axfr-port-num" dval:"53"`
    Ipv6AxfrPortNum int `json:"ipv6-axfr-port-num" dval:"53"`
    IpRefreshIntvl int `json:"ip-refresh-intvl" dval:"30"`
    Ipv6RefreshIntvl int `json:"ipv6-refresh-intvl" dval:"30"`
}


type DomainListMatchTypeEquals struct {
    Equals string `json:"equals"`
}


type DomainListMatchTypeSuffix struct {
    Suffix string `json:"suffix"`
}

func (p *DomainList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DomainList) getPath() string{
    return "domain-list"
}

func (p *DomainList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainList::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *DomainList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainList::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *DomainList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainList::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DomainList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
