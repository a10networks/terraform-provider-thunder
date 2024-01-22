

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSnmp struct {
	Inst struct {

    All int `json:"all"`
    Linkdown int `json:"linkdown"`
    Linkup int `json:"linkup"`
    Uuid string `json:"uuid"`

	} `json:"snmp"`
}

func (p *SnmpServerEnableTrapsSnmp) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSnmp) getPath() string{
    return "snmp-server/enable/traps/snmp"
}

func (p *SnmpServerEnableTrapsSnmp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSnmp::Post")
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

func (p *SnmpServerEnableTrapsSnmp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSnmp::Get")
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
func (p *SnmpServerEnableTrapsSnmp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSnmp::Put")
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

func (p *SnmpServerEnableTrapsSnmp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSnmp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
