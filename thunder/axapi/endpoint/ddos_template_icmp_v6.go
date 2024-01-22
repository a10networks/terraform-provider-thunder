

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateIcmpV6 struct {
	Inst struct {

    IcmpTmplName string `json:"icmp-tmpl-name"`
    TypeList []DdosTemplateIcmpV6TypeList `json:"type-list"`
    TypeOther DdosTemplateIcmpV6TypeOther297 `json:"type-other"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"icmp-v6"`
}


type DdosTemplateIcmpV6TypeList struct {
    TypeNumber int `json:"type-number"`
    TypeDeny int `json:"type-deny"`
    TypeRate int `json:"type-rate"`
    Code []DdosTemplateIcmpV6TypeListCode `json:"code"`
    CodeOther DdosTemplateIcmpV6TypeListCodeOther `json:"code-other"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosTemplateIcmpV6TypeListCode struct {
    CodeNumber int `json:"code-number"`
    CodeRate int `json:"code-rate"`
}


type DdosTemplateIcmpV6TypeListCodeOther struct {
    CodeOtherRate int `json:"code-other-rate"`
}


type DdosTemplateIcmpV6TypeOther297 struct {
    TypeOtherDeny int `json:"type-other-deny"`
    TypeOtherRate int `json:"type-other-rate"`
    Uuid string `json:"uuid"`
}

func (p *DdosTemplateIcmpV6) GetId() string{
    return url.QueryEscape(p.Inst.IcmpTmplName)
}

func (p *DdosTemplateIcmpV6) getPath() string{
    return "ddos/template/icmp-v6"
}

func (p *DdosTemplateIcmpV6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV6::Post")
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

func (p *DdosTemplateIcmpV6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV6::Get")
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
func (p *DdosTemplateIcmpV6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV6::Put")
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

func (p *DdosTemplateIcmpV6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateIcmpV6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
