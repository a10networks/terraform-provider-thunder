

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugGslbProtocol struct {
	Inst struct {

    ActiveRdt int `json:"active-rdt"`
    All int `json:"all"`
    Cache int `json:"cache"`
    Event int `json:"event"`
    Fsm int `json:"fsm"`
    Ip int `json:"ip"`
    Ipc int `json:"ipc"`
    KeepAlive int `json:"keep-alive"`
    Message int `json:"message"`
    MessageAll int `json:"message-all"`
    MessageArdtQuery int `json:"message-ardt-query"`
    MessageArdtReport int `json:"message-ardt-report"`
    MessageControl int `json:"message-control"`
    MessageKeepalive int `json:"message-keepalive"`
    MessageNotify int `json:"message-notify"`
    MessageOpen int `json:"message-open"`
    MessageQuery int `json:"message-query"`
    MessageUpdate int `json:"message-update"`
    Name string `json:"name"`
    Normal int `json:"normal"`
    PeerIpv4 string `json:"peer-ipv4"`
    PeerIpv6 string `json:"peer-ipv6"`
    Timer int `json:"timer"`
    Update int `json:"update"`
    Uuid string `json:"uuid"`

	} `json:"protocol"`
}

func (p *DebugGslbProtocol) GetId() string{
    return "1"
}

func (p *DebugGslbProtocol) getPath() string{
    return "debug/gslb/protocol"
}

func (p *DebugGslbProtocol) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbProtocol::Post")
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

func (p *DebugGslbProtocol) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbProtocol::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DebugGslbProtocol) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbProtocol::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *DebugGslbProtocol) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbProtocol::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
