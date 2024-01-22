

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugOspfPacket struct {
	Inst struct {

    Dd int `json:"dd"`
    Detail int `json:"detail"`
    Hello int `json:"hello"`
    LsAck int `json:"ls-ack"`
    LsRequest int `json:"ls-request"`
    LsUpdate int `json:"ls-update"`
    Recv int `json:"recv"`
    Send int `json:"send"`
    Uuid string `json:"uuid"`

	} `json:"packet"`
}

func (p *DebugOspfPacket) GetId() string{
    return "1"
}

func (p *DebugOspfPacket) getPath() string{
    return "debug/ospf/packet"
}

func (p *DebugOspfPacket) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfPacket::Post")
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

func (p *DebugOspfPacket) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfPacket::Get")
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
func (p *DebugOspfPacket) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfPacket::Put")
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

func (p *DebugOspfPacket) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfPacket::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
