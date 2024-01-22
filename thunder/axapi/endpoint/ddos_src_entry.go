

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcEntry struct {
	Inst struct {

    AppTypeList []DdosSrcEntryAppTypeList `json:"app-type-list"`
    Bypass int `json:"bypass"`
    Description string `json:"description"`
    ExceedLogCfg DdosSrcEntryExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    HwBlacklistBlocking DdosSrcEntryHwBlacklistBlocking293 `json:"hw-blacklist-blocking"`
    IpAddr string `json:"ip-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4TypeList []DdosSrcEntryL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    SrcEntryName string `json:"src-entry-name"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Template DdosSrcEntryTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"entry"`
}


type DdosSrcEntryAppTypeList struct {
    Protocol string `json:"protocol"`
    Template DdosSrcEntryAppTypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcEntryAppTypeListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosSrcEntryExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosSrcEntryHwBlacklistBlocking293 struct {
    SrcEnable int `json:"src-enable"`
    Uuid string `json:"uuid"`
}


type DdosSrcEntryL4TypeList struct {
    Protocol string `json:"protocol"`
    Action string `json:"action"`
    Glid string `json:"glid"`
    Template DdosSrcEntryL4TypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcEntryL4TypeListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosSrcEntryTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosSrcEntry) GetId() string{
    return url.QueryEscape(p.Inst.SrcEntryName)
}

func (p *DdosSrcEntry) getPath() string{
    return "ddos/src/entry"
}

func (p *DdosSrcEntry) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntry::Post")
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

func (p *DdosSrcEntry) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntry::Get")
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
func (p *DdosSrcEntry) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntry::Put")
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

func (p *DdosSrcEntry) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntry::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
