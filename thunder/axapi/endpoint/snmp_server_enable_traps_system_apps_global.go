

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSystemAppsGlobal struct {
	Inst struct {

    CpsThreshold int `json:"cps-threshold"`
    SessionsThreshold int `json:"sessions-threshold"`
    Uuid string `json:"uuid"`

	} `json:"apps-global"`
}

func (p *SnmpServerEnableTrapsSystemAppsGlobal) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSystemAppsGlobal) getPath() string{
    return "snmp-server/enable/traps/system/apps-global"
}

func (p *SnmpServerEnableTrapsSystemAppsGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystemAppsGlobal::Post")
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

func (p *SnmpServerEnableTrapsSystemAppsGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystemAppsGlobal::Get")
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
func (p *SnmpServerEnableTrapsSystemAppsGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystemAppsGlobal::Put")
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

func (p *SnmpServerEnableTrapsSystemAppsGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystemAppsGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
