

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SctpPermitPayloadProtocolProtocolName struct {
	Inst struct {

    Protocol string `json:"protocol"`
    Uuid string `json:"uuid"`

	} `json:"protocol-name"`
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) GetId() string{
    return p.Inst.Protocol
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) getPath() string{
    return "cgnv6/sctp/permit-payload-protocol/protocol-name"
}

func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolName::Post")
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

func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolName::Get")
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
func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolName::Put")
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

func (p *Cgnv6SctpPermitPayloadProtocolProtocolName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpPermitPayloadProtocolProtocolName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
