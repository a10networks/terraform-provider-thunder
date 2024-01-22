

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemManagementInterfaceMode struct {
	Inst struct {

    Dedicated int `json:"dedicated"`
    NonDedicated int `json:"non-dedicated"`

	} `json:"management-interface-mode"`
}

func (p *SystemManagementInterfaceMode) GetId() string{
    return "1"
}

func (p *SystemManagementInterfaceMode) getPath() string{
    return "system/management-interface-mode"
}

func (p *SystemManagementInterfaceMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemManagementInterfaceMode::Post")
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

func (p *SystemManagementInterfaceMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemManagementInterfaceMode::Get")
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
func (p *SystemManagementInterfaceMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemManagementInterfaceMode::Put")
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

func (p *SystemManagementInterfaceMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemManagementInterfaceMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
