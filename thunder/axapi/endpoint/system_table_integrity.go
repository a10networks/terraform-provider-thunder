

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTableIntegrity struct {
	Inst struct {

    AuditAction string `json:"audit-action" dval:"enable"`
    AutoSyncAction string `json:"auto-sync-action" dval:"enable"`
    SamplingEnable []SystemTableIntegritySamplingEnable `json:"sampling-enable"`
    Table string `json:"table" dval:"all"`
    Uuid string `json:"uuid"`

	} `json:"table-integrity"`
}


type SystemTableIntegritySamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
}

func (p *SystemTableIntegrity) GetId() string{
    return "1"
}

func (p *SystemTableIntegrity) getPath() string{
    return "system/table-integrity"
}

func (p *SystemTableIntegrity) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTableIntegrity::Post")
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

func (p *SystemTableIntegrity) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTableIntegrity::Get")
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
func (p *SystemTableIntegrity) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTableIntegrity::Put")
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

func (p *SystemTableIntegrity) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTableIntegrity::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
