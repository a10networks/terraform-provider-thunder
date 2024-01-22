

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FlowspecPacketLength struct {
	Inst struct {

    Length int `json:"length"`
    LengthEnd int `json:"length-end"`
    PacketLengthAttribute string `json:"packet-length-attribute"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"packet-length"`
}

func (p *FlowspecPacketLength) GetId() string{
    return p.Inst.PacketLengthAttribute+"+"+strconv.Itoa(p.Inst.Length)
}

func (p *FlowspecPacketLength) getPath() string{
    return "flowspec/" +p.Inst.Name + "/packet-length"
}

func (p *FlowspecPacketLength) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecPacketLength::Post")
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

func (p *FlowspecPacketLength) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecPacketLength::Get")
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
func (p *FlowspecPacketLength) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecPacketLength::Put")
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

func (p *FlowspecPacketLength) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecPacketLength::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
