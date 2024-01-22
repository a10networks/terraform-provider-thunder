

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtCommon struct {
	Inst struct {

    Delay int `json:"delay"`
    ProceedOnError int `json:"proceed-on-error"`
    Uuid string `json:"uuid"`

	} `json:"common"`
}

func (p *SysUtCommon) GetId() string{
    return "1"
}

func (p *SysUtCommon) getPath() string{
    return "sys-ut/common"
}

func (p *SysUtCommon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtCommon::Post")
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

func (p *SysUtCommon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtCommon::Get")
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
func (p *SysUtCommon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtCommon::Put")
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

func (p *SysUtCommon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtCommon::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
