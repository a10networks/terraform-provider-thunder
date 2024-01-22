

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugSip struct {
	Inst struct {

    Ack int `json:"ACK"`
    Bye int `json:"BYE"`
    Cancel int `json:"CANCEL"`
    Info int `json:"INFO"`
    Invite int `json:"INVITE"`
    Message int `json:"MESSAGE"`
    Method int `json:"method"`
    Notify int `json:"NOTIFY"`
    Options int `json:"OPTIONS"`
    Prack int `json:"PRACK"`
    Publish int `json:"PUBLISH"`
    Refer int `json:"REFER"`
    Register int `json:"REGISTER"`
    Subscribe int `json:"SUBSCRIBE"`
    Update int `json:"UPDATE"`
    Uuid string `json:"uuid"`

	} `json:"sip"`
}

func (p *DebugSip) GetId() string{
    return "1"
}

func (p *DebugSip) getPath() string{
    return "debug/sip"
}

func (p *DebugSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugSip::Post")
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

func (p *DebugSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugSip::Get")
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
func (p *DebugSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugSip::Put")
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

func (p *DebugSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
