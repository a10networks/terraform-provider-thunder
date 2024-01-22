

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugGslbGroup struct {
	Inst struct {

    All int `json:"all"`
    Cache int `json:"cache"`
    Event int `json:"event"`
    Fsm int `json:"fsm"`
    Ip int `json:"ip"`
    Ipc int `json:"ipc"`
    KeepAlive int `json:"keep-alive"`
    Message int `json:"message"`
    MessageAll int `json:"message-all"`
    MessageControl int `json:"message-control"`
    MessageGroup int `json:"message-group"`
    MessageKeepalive int `json:"message-keepalive"`
    MessageNotify int `json:"message-notify"`
    MessageOpen int `json:"message-open"`
    MessageQuery int `json:"message-query"`
    Name string `json:"name"`
    Normal int `json:"normal"`
    PeerIpv4 string `json:"peer-ipv4"`
    PeerIpv6 string `json:"peer-ipv6"`
    Timer int `json:"timer"`
    Update int `json:"update"`
    Uuid string `json:"uuid"`

	} `json:"group"`
}

func (p *DebugGslbGroup) GetId() string{
    return "1"
}

func (p *DebugGslbGroup) getPath() string{
    return "debug/gslb/group"
}

func (p *DebugGslbGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbGroup::Post")
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

func (p *DebugGslbGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbGroup::Get")
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
func (p *DebugGslbGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbGroup::Put")
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

func (p *DebugGslbGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugGslbGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
