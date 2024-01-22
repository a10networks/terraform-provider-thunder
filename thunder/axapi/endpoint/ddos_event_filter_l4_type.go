

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosEventFilterL4Type struct {
	Inst struct {

    OutOfSeq int `json:"out-of-seq"`
    Protocol string `json:"protocol"`
    RetransSynCfg DdosEventFilterL4TypeRetransSynCfg `json:"retrans-syn-cfg"`
    TcpAuth DdosEventFilterL4TypeTcpAuth `json:"tcp-auth"`
    UdpAuth DdosEventFilterL4TypeUdpAuth `json:"udp-auth"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZeroWindow int `json:"zero-window"`
    FilterName string 

	} `json:"l4-type"`
}


type DdosEventFilterL4TypeRetransSynCfg struct {
    RetransSyn int `json:"retrans-syn"`
    RetransSynExceed int `json:"retrans-syn-exceed"`
}


type DdosEventFilterL4TypeTcpAuth struct {
    TcpAuthInit int `json:"tcp-auth-init"`
    TcpAuthPass int `json:"tcp-auth-pass"`
    TcpAuthFail int `json:"tcp-auth-fail"`
}


type DdosEventFilterL4TypeUdpAuth struct {
    UdpAuthInit int `json:"udp-auth-init"`
    UdpAuthPass int `json:"udp-auth-pass"`
}

func (p *DdosEventFilterL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosEventFilterL4Type) getPath() string{
    return "ddos/event-filter/" +p.Inst.FilterName + "/l4-type"
}

func (p *DdosEventFilterL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilterL4Type::Post")
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

func (p *DdosEventFilterL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilterL4Type::Get")
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
func (p *DdosEventFilterL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilterL4Type::Put")
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

func (p *DdosEventFilterL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosEventFilterL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
