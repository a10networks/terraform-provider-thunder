

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDistributedForwardingFw struct {
	Inst struct {

    FwValue string `json:"fw-value" dval:"disable"`
    SessionOffloadDirection string `json:"session-offload-direction" dval:"both"`
    Threshold []ScaleoutDistributedForwardingFwThreshold `json:"threshold"`
    Uuid string `json:"uuid"`

	} `json:"fw"`
}


type ScaleoutDistributedForwardingFwThreshold struct {
    ThresholdValue int `json:"threshold-value" dval:"5"`
    ProtocolValue string `json:"protocol-value"`
}

func (p *ScaleoutDistributedForwardingFw) GetId() string{
    return "1"
}

func (p *ScaleoutDistributedForwardingFw) getPath() string{
    return "scaleout/distributed-forwarding/fw"
}

func (p *ScaleoutDistributedForwardingFw) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingFw::Post")
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

func (p *ScaleoutDistributedForwardingFw) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingFw::Get")
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
func (p *ScaleoutDistributedForwardingFw) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingFw::Put")
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

func (p *ScaleoutDistributedForwardingFw) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingFw::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
