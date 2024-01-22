

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateHttpHttpProtocolCheck struct {
	Inst struct {

    GetAndPayload string `json:"get-and-payload"`
    H2upContentLengthAlias string `json:"h2up-content-length-alias"`
    H2upWithHostAndAuth string `json:"h2up-with-host-and-auth"`
    H2upWithTransferEncoding string `json:"h2up-with-transfer-encoding"`
    HeaderFilterRuleList []SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList `json:"header-filter-rule-list"`
    MalformedH2upHeaderValue string `json:"malformed-h2up-header-value"`
    MalformedH2upSchemeValue string `json:"malformed-h2up-scheme-value"`
    MultipleContentLength string `json:"multiple-content-length"`
    MultipleTransferEncoding string `json:"multiple-transfer-encoding"`
    TransferEncodingAndContentLength string `json:"transfer-encoding-and-content-length"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"http-protocol-check"`
}


type SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList struct {
    SeqNum int `json:"seq-num"`
    MatchTypeValue string `json:"match-type-value"`
    HeaderNameValue string `json:"header-name-value"`
    HeaderValueValue string `json:"header-value-value"`
    ActionValue string `json:"action-value"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *SlbTemplateHttpHttpProtocolCheck) GetId() string{
    return "1"
}

func (p *SlbTemplateHttpHttpProtocolCheck) getPath() string{
    return "slb/template/http/" +p.Inst.Name + "/http-protocol-check"
}

func (p *SlbTemplateHttpHttpProtocolCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheck::Post")
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

func (p *SlbTemplateHttpHttpProtocolCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheck::Get")
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
func (p *SlbTemplateHttpHttpProtocolCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheck::Put")
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

func (p *SlbTemplateHttpHttpProtocolCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpHttpProtocolCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
