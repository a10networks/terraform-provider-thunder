

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosEventFilter struct {
	Inst struct {

    BlackList DdosEventFilterBlackList `json:"black-list"`
    Drop DdosEventFilterDrop `json:"drop"`
    FilterName string `json:"filter-name"`
    L4TypeList []DdosEventFilterL4TypeList `json:"l4-type-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    WhiteList DdosEventFilterWhiteList `json:"white-list"`

	} `json:"event-filter"`
}


type DdosEventFilterBlackList struct {
    BlackListDst int `json:"black-list-dst"`
    BlackListSrc int `json:"black-list-src"`
}


type DdosEventFilterDrop struct {
    DropSrc int `json:"drop-src"`
    DropDst int `json:"drop-dst"`
    DropBlackList int `json:"drop-black-list"`
}


type DdosEventFilterL4TypeList struct {
    Protocol string `json:"protocol"`
    TcpAuth DdosEventFilterL4TypeListTcpAuth `json:"tcp-auth"`
    RetransSynCfg DdosEventFilterL4TypeListRetransSynCfg `json:"retrans-syn-cfg"`
    OutOfSeq int `json:"out-of-seq"`
    ZeroWindow int `json:"zero-window"`
    UdpAuth DdosEventFilterL4TypeListUdpAuth `json:"udp-auth"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosEventFilterL4TypeListTcpAuth struct {
    TcpAuthInit int `json:"tcp-auth-init"`
    TcpAuthPass int `json:"tcp-auth-pass"`
    TcpAuthFail int `json:"tcp-auth-fail"`
}


type DdosEventFilterL4TypeListRetransSynCfg struct {
    RetransSyn int `json:"retrans-syn"`
    RetransSynExceed int `json:"retrans-syn-exceed"`
}


type DdosEventFilterL4TypeListUdpAuth struct {
    UdpAuthInit int `json:"udp-auth-init"`
    UdpAuthPass int `json:"udp-auth-pass"`
}


type DdosEventFilterWhiteList struct {
    WhiteListDst int `json:"white-list-dst"`
    WhiteListSrc int `json:"white-list-src"`
}

func (p *DdosEventFilter) GetId() string{
    return url.QueryEscape(p.Inst.FilterName)
}

func (p *DdosEventFilter) getPath() string{
    return "ddos/event-filter"
}

func (p *DdosEventFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilter::Post")
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

func (p *DdosEventFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilter::Get")
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
func (p *DdosEventFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilter::Put")
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

func (p *DdosEventFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
