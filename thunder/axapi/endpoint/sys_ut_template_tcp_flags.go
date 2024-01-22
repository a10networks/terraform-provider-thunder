

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateTcpFlags struct {
	Inst struct {

    Ack int `json:"ack"`
    Cwr int `json:"cwr"`
    Ece int `json:"ece"`
    Fin int `json:"fin"`
    Psh int `json:"psh"`
    Rst int `json:"rst"`
    Syn int `json:"syn"`
    Urg int `json:"urg"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"flags"`
}

func (p *SysUtTemplateTcpFlags) GetId() string{
    return "1"
}

func (p *SysUtTemplateTcpFlags) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/tcp/flags"
}

func (p *SysUtTemplateTcpFlags) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpFlags::Post")
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

func (p *SysUtTemplateTcpFlags) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpFlags::Get")
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
func (p *SysUtTemplateTcpFlags) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpFlags::Put")
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

func (p *SysUtTemplateTcpFlags) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpFlags::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
