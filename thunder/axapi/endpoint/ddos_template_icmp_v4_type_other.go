

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateIcmpV4TypeOther struct {
	Inst struct {

    TypeOtherDeny int `json:"type-other-deny"`
    TypeOtherRate int `json:"type-other-rate"`
    Uuid string `json:"uuid"`
    IcmpTmplName string 

	} `json:"type-other"`
}

func (p *DdosTemplateIcmpV4TypeOther) GetId() string{
    return "1"
}

func (p *DdosTemplateIcmpV4TypeOther) getPath() string{
    return "ddos/template/icmp-v4/" +p.Inst.IcmpTmplName + "/type-other"
}

func (p *DdosTemplateIcmpV4TypeOther) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4TypeOther::Post")
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

func (p *DdosTemplateIcmpV4TypeOther) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4TypeOther::Get")
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
func (p *DdosTemplateIcmpV4TypeOther) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4TypeOther::Put")
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

func (p *DdosTemplateIcmpV4TypeOther) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4TypeOther::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
