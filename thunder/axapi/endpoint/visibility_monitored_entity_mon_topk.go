

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityMonTopk struct {
	Inst struct {

    Sources VisibilityMonitoredEntityMonTopkSources1925 `json:"sources"`
    Uuid string `json:"uuid"`

	} `json:"mon-topk"`
}


type VisibilityMonitoredEntityMonTopkSources1925 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityMonitoredEntityMonTopk) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityMonTopk) getPath() string{
    return "visibility/monitored-entity/mon-topk"
}

func (p *VisibilityMonitoredEntityMonTopk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityMonTopk::Post")
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

func (p *VisibilityMonitoredEntityMonTopk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityMonTopk::Get")
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
func (p *VisibilityMonitoredEntityMonTopk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityMonTopk::Put")
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

func (p *VisibilityMonitoredEntityMonTopk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntityMonTopk::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
