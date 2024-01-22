

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugGslb struct {
	Inst struct {

    Glname string `json:"glname"`
    Group DebugGslbGroup323 `json:"group"`
    Id1 int `json:"id1"`
    IpAddr string `json:"ip-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Memory int `json:"memory"`
    OneShot int `json:"one-shot"`
    Protocol DebugGslbProtocol324 `json:"protocol"`
    State int `json:"state"`
    Uuid string `json:"uuid"`

	} `json:"gslb"`
}


type DebugGslbGroup323 struct {
    Cache int `json:"cache"`
    Event int `json:"event"`
    All int `json:"all"`
    Fsm int `json:"fsm"`
    Ip int `json:"ip"`
    PeerIpv4 string `json:"peer-ipv4"`
    PeerIpv6 string `json:"peer-ipv6"`
    Ipc int `json:"ipc"`
    KeepAlive int `json:"keep-alive"`
    Message int `json:"message"`
    MessageAll int `json:"message-all"`
    MessageKeepalive int `json:"message-keepalive"`
    MessageNotify int `json:"message-notify"`
    MessageControl int `json:"message-control"`
    MessageQuery int `json:"message-query"`
    MessageOpen int `json:"message-open"`
    MessageGroup int `json:"message-group"`
    Name string `json:"name"`
    Normal int `json:"normal"`
    Timer int `json:"timer"`
    Update int `json:"update"`
    Uuid string `json:"uuid"`
}


type DebugGslbProtocol324 struct {
    Cache int `json:"cache"`
    Event int `json:"event"`
    All int `json:"all"`
    ActiveRdt int `json:"active-rdt"`
    Fsm int `json:"fsm"`
    Ip int `json:"ip"`
    PeerIpv4 string `json:"peer-ipv4"`
    PeerIpv6 string `json:"peer-ipv6"`
    Ipc int `json:"ipc"`
    KeepAlive int `json:"keep-alive"`
    Message int `json:"message"`
    MessageAll int `json:"message-all"`
    MessageKeepalive int `json:"message-keepalive"`
    MessageNotify int `json:"message-notify"`
    MessageControl int `json:"message-control"`
    MessageQuery int `json:"message-query"`
    MessageOpen int `json:"message-open"`
    MessageUpdate int `json:"message-update"`
    MessageArdtQuery int `json:"message-ardt-query"`
    MessageArdtReport int `json:"message-ardt-report"`
    Name string `json:"name"`
    Normal int `json:"normal"`
    Timer int `json:"timer"`
    Update int `json:"update"`
    Uuid string `json:"uuid"`
}

func (p *DebugGslb) GetId() string{
    return "1"
}

func (p *DebugGslb) getPath() string{
    return "debug/gslb"
}

func (p *DebugGslb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslb::Post")
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

func (p *DebugGslb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslb::Get")
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
func (p *DebugGslb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslb::Put")
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

func (p *DebugGslb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
