

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntity struct {
	Inst struct {

    Detail VisibilityMonitoredEntityDetail1927 `json:"detail"`
    MonTopk VisibilityMonitoredEntityMonTopk1929 `json:"mon-topk"`
    Secondary VisibilityMonitoredEntitySecondary1931 `json:"secondary"`
    Sessions VisibilityMonitoredEntitySessions1934 `json:"sessions"`
    Uuid string `json:"uuid"`

	} `json:"monitored-entity"`
}


type VisibilityMonitoredEntityDetail1927 struct {
    Uuid string `json:"uuid"`
    Debug VisibilityMonitoredEntityDetailDebug1928 `json:"debug"`
}


type VisibilityMonitoredEntityDetailDebug1928 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntityMonTopk1929 struct {
    Uuid string `json:"uuid"`
    Sources VisibilityMonitoredEntityMonTopkSources1930 `json:"sources"`
}


type VisibilityMonitoredEntityMonTopkSources1930 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntitySecondary1931 struct {
    MonTopk VisibilityMonitoredEntitySecondaryMonTopk1932 `json:"mon-topk"`
}


type VisibilityMonitoredEntitySecondaryMonTopk1932 struct {
    Uuid string `json:"uuid"`
    Sources VisibilityMonitoredEntitySecondaryMonTopkSources1933 `json:"sources"`
}


type VisibilityMonitoredEntitySecondaryMonTopkSources1933 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntitySessions1934 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityMonitoredEntity) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntity) getPath() string{
    return "visibility/monitored-entity"
}

func (p *VisibilityMonitoredEntity) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntity::Post")
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

func (p *VisibilityMonitoredEntity) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntity::Get")
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
func (p *VisibilityMonitoredEntity) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntity::Put")
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

func (p *VisibilityMonitoredEntity) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitoredEntity::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
