

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateTcpOptions struct {
	Inst struct {

    Mss int `json:"mss"`
    Nop int `json:"nop"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Uuid string `json:"uuid"`
    Wscale int `json:"wscale"`
    Name string 

	} `json:"options"`
}

func (p *SysUtTemplateTcpOptions) GetId() string{
    return "1"
}

func (p *SysUtTemplateTcpOptions) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/tcp/options"
}

func (p *SysUtTemplateTcpOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpOptions::Post")
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

func (p *SysUtTemplateTcpOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpOptions::Get")
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
func (p *SysUtTemplateTcpOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpOptions::Put")
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

func (p *SysUtTemplateTcpOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcpOptions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
