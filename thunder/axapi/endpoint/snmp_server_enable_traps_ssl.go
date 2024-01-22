

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSsl struct {
	Inst struct {

    ServerCertificateError int `json:"server-certificate-error"`
    Uuid string `json:"uuid"`

	} `json:"ssl"`
}

func (p *SnmpServerEnableTrapsSsl) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSsl) getPath() string{
    return "snmp-server/enable/traps/ssl"
}

func (p *SnmpServerEnableTrapsSsl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSsl::Post")
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

func (p *SnmpServerEnableTrapsSsl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSsl::Get")
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
func (p *SnmpServerEnableTrapsSsl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSsl::Put")
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

func (p *SnmpServerEnableTrapsSsl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSsl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
