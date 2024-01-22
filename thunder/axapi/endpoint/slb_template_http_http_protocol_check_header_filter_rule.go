

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateHttpHttpProtocolCheckHeaderFilterRule struct {
	Inst struct {

    ActionValue string `json:"action-value"`
    HeaderNameValue string `json:"header-name-value"`
    HeaderValueValue string `json:"header-value-value"`
    MatchTypeValue string `json:"match-type-value"`
    SeqNum int `json:"seq-num"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"header-filter-rule"`
}

func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) GetId() string{
    return strconv.Itoa(p.Inst.SeqNum)
}

func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) getPath() string{
    return "slb/template/http/" +p.Inst.Name + "/http-protocol-check/header-filter-rule"
}

func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheckHeaderFilterRule::Post")
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

func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheckHeaderFilterRule::Get")
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
func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheckHeaderFilterRule::Put")
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

func (p *SlbTemplateHttpHttpProtocolCheckHeaderFilterRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheckHeaderFilterRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
