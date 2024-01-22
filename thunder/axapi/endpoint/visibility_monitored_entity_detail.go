

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityDetail struct {
	Inst struct {

    Debug VisibilityMonitoredEntityDetailDebug1924 `json:"debug"`
    Uuid string `json:"uuid"`

	} `json:"detail"`
}


type VisibilityMonitoredEntityDetailDebug1924 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityMonitoredEntityDetail) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityDetail) getPath() string{
    return "visibility/monitored-entity/detail"
}

func (p *VisibilityMonitoredEntityDetail) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityDetail::Post")
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

func (p *VisibilityMonitoredEntityDetail) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityDetail::Get")
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
func (p *VisibilityMonitoredEntityDetail) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityDetail::Put")
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

func (p *VisibilityMonitoredEntityDetail) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityDetail::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
