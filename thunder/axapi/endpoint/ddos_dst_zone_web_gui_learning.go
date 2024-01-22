

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneWebGuiLearning struct {
	Inst struct {

    Duration string `json:"duration" dval:"6hour"`
    StartingTime string `json:"starting-time"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"learning"`
}

func (p *DdosDstZoneWebGuiLearning) GetId() string{
    return "1"
}

func (p *DdosDstZoneWebGuiLearning) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/web-gui/learning"
}

func (p *DdosDstZoneWebGuiLearning) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiLearning::Post")
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

func (p *DdosDstZoneWebGuiLearning) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiLearning::Get")
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
func (p *DdosDstZoneWebGuiLearning) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiLearning::Put")
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

func (p *DdosDstZoneWebGuiLearning) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiLearning::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
