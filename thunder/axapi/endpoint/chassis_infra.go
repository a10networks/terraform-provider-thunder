

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ChassisInfra struct {
	Inst struct {

    DebugDisable int `json:"debug-disable"`
    DebugEnable int `json:"debug-enable"`
    DebugStatus int `json:"debug-status"`
    Detailed int `json:"detailed"`
    SysSync int `json:"sys-sync"`
    SystemSyncVerify int `json:"system-sync-verify"`

	} `json:"chassis-infra"`
}

func (p *ChassisInfra) GetId() string{
    return "1"
}

func (p *ChassisInfra) getPath() string{
    return "chassis-infra"
}

func (p *ChassisInfra) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ChassisInfra::Post")
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

func (p *ChassisInfra) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ChassisInfra::Get")
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
func (p *ChassisInfra) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ChassisInfra::Put")
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

func (p *ChassisInfra) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ChassisInfra::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
