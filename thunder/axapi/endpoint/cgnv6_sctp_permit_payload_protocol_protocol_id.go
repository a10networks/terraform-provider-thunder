

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SctpPermitPayloadProtocolProtocolId struct {
	Inst struct {

    Id1 int `json:"id1"`
    Uuid string `json:"uuid"`

	} `json:"protocol-id"`
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) GetId() string{
    return strconv.Itoa(p.Inst.Id1)
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) getPath() string{
    return "cgnv6/sctp/permit-payload-protocol/protocol-id"
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolId::Post")
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

func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolId::Get")
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
func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolId::Put")
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

func (p *Cgnv6SctpPermitPayloadProtocolProtocolId) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolId::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
