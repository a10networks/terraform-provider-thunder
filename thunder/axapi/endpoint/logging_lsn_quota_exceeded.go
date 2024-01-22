

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLsnQuotaExceeded struct {
	Inst struct {

    Custom1 int `json:"custom1"`
    Custom2 int `json:"custom2"`
    Custom3 int `json:"custom3"`
    Custom4 int `json:"custom4"`
    Custom5 int `json:"custom5"`
    Custom6 int `json:"custom6"`
    DisablePoolBased int `json:"disable-pool-based"`
    Imei int `json:"imei"`
    Imsi int `json:"imsi"`
    IpBased int `json:"ip-based"`
    Msisdn int `json:"msisdn"`
    Uuid string `json:"uuid"`
    WithRadiusAttribute int `json:"with-radius-attribute"`

	} `json:"quota-exceeded"`
}

func (p *LoggingLsnQuotaExceeded) GetId() string{
    return "1"
}

func (p *LoggingLsnQuotaExceeded) getPath() string{
    return "logging/lsn/quota-exceeded"
}

func (p *LoggingLsnQuotaExceeded) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnQuotaExceeded::Post")
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

func (p *LoggingLsnQuotaExceeded) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnQuotaExceeded::Get")
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
func (p *LoggingLsnQuotaExceeded) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnQuotaExceeded::Put")
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

func (p *LoggingLsnQuotaExceeded) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnQuotaExceeded::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
