

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateIcmpV4Type struct {
	Inst struct {

    Code []DdosTemplateIcmpV4TypeCode `json:"code"`
    CodeOther DdosTemplateIcmpV4TypeCodeOther `json:"code-other"`
    TypeDeny int `json:"type-deny"`
    TypeNumber int `json:"type-number"`
    TypeRate int `json:"type-rate"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    IcmpTmplName string 

	} `json:"type"`
}


type DdosTemplateIcmpV4TypeCode struct {
    CodeNumber int `json:"code-number"`
    CodeRate int `json:"code-rate"`
}


type DdosTemplateIcmpV4TypeCodeOther struct {
    CodeOtherRate int `json:"code-other-rate"`
}

func (p *DdosTemplateIcmpV4Type) GetId() string{
    return strconv.Itoa(p.Inst.TypeNumber)
}

func (p *DdosTemplateIcmpV4Type) getPath() string{
    return "ddos/template/icmp-v4/" +p.Inst.IcmpTmplName + "/type"
}

func (p *DdosTemplateIcmpV4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4Type::Post")
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

func (p *DdosTemplateIcmpV4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4Type::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosTemplateIcmpV4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4Type::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosTemplateIcmpV4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
