

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingAuditlog struct {
	Inst struct {

    AuditFacility string `json:"audit-facility"`
    Host4 string `json:"host4"`
    Host6 string `json:"host6"`
    PartitionName string `json:"partition-name"`
    Port int `json:"port" dval:"514"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`

	} `json:"auditlog"`
}

func (p *LoggingAuditlog) GetId() string{
    return "1"
}

func (p *LoggingAuditlog) getPath() string{
    return "logging/auditlog"
}

func (p *LoggingAuditlog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingAuditlog::Post")
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

func (p *LoggingAuditlog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingAuditlog::Get")
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
func (p *LoggingAuditlog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingAuditlog::Put")
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

func (p *LoggingAuditlog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingAuditlog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
