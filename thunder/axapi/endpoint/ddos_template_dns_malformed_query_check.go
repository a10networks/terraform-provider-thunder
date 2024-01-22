

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateDnsMalformedQueryCheck struct {
	Inst struct {

    NonQueryOpcodeCheck string `json:"non-query-opcode-check"`
    SkipMultiPacketCheck int `json:"skip-multi-packet-check"`
    Uuid string `json:"uuid"`
    ValidationType string `json:"validation-type"`
    Name string 

	} `json:"malformed-query-check"`
}

func (p *DdosTemplateDnsMalformedQueryCheck) GetId() string{
    return "1"
}

func (p *DdosTemplateDnsMalformedQueryCheck) getPath() string{
    return "ddos/template/dns/" +p.Inst.Name + "/malformed-query-check"
}

func (p *DdosTemplateDnsMalformedQueryCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDnsMalformedQueryCheck::Post")
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

func (p *DdosTemplateDnsMalformedQueryCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDnsMalformedQueryCheck::Get")
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
func (p *DdosTemplateDnsMalformedQueryCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDnsMalformedQueryCheck::Put")
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

func (p *DdosTemplateDnsMalformedQueryCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDnsMalformedQueryCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
