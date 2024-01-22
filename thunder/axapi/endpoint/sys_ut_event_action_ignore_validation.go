

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtEventActionIgnoreValidation struct {
	Inst struct {

    All int `json:"all"`
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    Uuid string `json:"uuid"`
    EventNumber string 
    Direction string 

	} `json:"ignore-validation"`
}

func (p *SysUtEventActionIgnoreValidation) GetId() string{
    return "1"
}

func (p *SysUtEventActionIgnoreValidation) getPath() string{
    return "sys-ut/event/" +p.Inst.EventNumber + "/action/" +p.Inst.Direction + "/ignore-validation"
}

func (p *SysUtEventActionIgnoreValidation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionIgnoreValidation::Post")
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

func (p *SysUtEventActionIgnoreValidation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionIgnoreValidation::Get")
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
func (p *SysUtEventActionIgnoreValidation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionIgnoreValidation::Put")
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

func (p *SysUtEventActionIgnoreValidation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionIgnoreValidation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
