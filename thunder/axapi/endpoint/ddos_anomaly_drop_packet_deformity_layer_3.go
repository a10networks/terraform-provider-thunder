

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosAnomalyDropPacketDeformityLayer3 struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Log int `json:"log"`
    Toggle string `json:"toggle" dval:"enable"`
    Uuid string `json:"uuid"`

	} `json:"packet-deformity-layer-3"`
}

func (p *DdosAnomalyDropPacketDeformityLayer3) GetId() string{
    return "1"
}

func (p *DdosAnomalyDropPacketDeformityLayer3) getPath() string{
    return "ddos/anomaly-drop/packet-deformity-layer-3"
}

func (p *DdosAnomalyDropPacketDeformityLayer3) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosAnomalyDropPacketDeformityLayer3::Post")
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

func (p *DdosAnomalyDropPacketDeformityLayer3) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosAnomalyDropPacketDeformityLayer3::Get")
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
func (p *DdosAnomalyDropPacketDeformityLayer3) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosAnomalyDropPacketDeformityLayer3::Put")
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

func (p *DdosAnomalyDropPacketDeformityLayer3) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosAnomalyDropPacketDeformityLayer3::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
