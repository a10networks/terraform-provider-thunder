

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsLogging struct {
	Inst struct {

    Disable int `json:"disable"`
    DnsLoggingProtocol string `json:"dns-logging-protocol"`
    DnsLoggingRequestSection string `json:"dns-logging-request-section"`
    DnsLoggingResponseSection string `json:"dns-logging-response-section"`
    DnsLoggingType string `json:"dns-logging-type"`
    Name string `json:"name"`
    ResponseType SlbTemplateDnsLoggingResponseType1417 `json:"response-type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns-logging"`
}


type SlbTemplateDnsLoggingResponseType1417 struct {
    Config int `json:"config"`
    Uuid string `json:"uuid"`
    TypeList []SlbTemplateDnsLoggingResponseTypeTypeList1418 `json:"type-list"`
}


type SlbTemplateDnsLoggingResponseTypeTypeList1418 struct {
    ResponseTypeName string `json:"response-type-name"`
    LengthLimitFlag int `json:"length-limit-flag"`
    TxtData int `json:"txt-data"`
    TxtTypeLimitNum int `json:"txt-type-limit-num"`
    TxtTypeNoLimit int `json:"txt-type-no-limit"`
    Signature int `json:"signature"`
    RrsigTypeLimitNum int `json:"rrsig-type-limit-num"`
    RrsigTypeNoLimit int `json:"rrsig-type-no-limit"`
    OtherData int `json:"other-data"`
    TsigTypeLimitNum int `json:"tsig-type-limit-num"`
    TsigTypeNoLimit int `json:"tsig-type-no-limit"`
    PublicKey int `json:"public-key"`
    DnskeyTypeLimitNum int `json:"dnskey-type-limit-num"`
    DnskeyTypeNoLimit int `json:"dnskey-type-no-limit"`
    Digest int `json:"digest"`
    DsTypeLimitNum int `json:"ds-type-limit-num"`
    DsTypeNoLimit int `json:"ds-type-no-limit"`
    ValueField int `json:"value-field"`
    CaaTypeLimitNum int `json:"caa-type-limit-num"`
    CaaTypeNoLimit int `json:"caa-type-no-limit"`
    ServiceField int `json:"service-field"`
    NaptrTypeLimitNum int `json:"naptr-type-limit-num"`
    NaptrTypeNoLimit int `json:"naptr-type-no-limit"`
    RdataField int `json:"rdata-field"`
    OptTypeLimitNum int `json:"opt-type-limit-num"`
    OptTypeNoLimit int `json:"opt-type-no-limit"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *SlbTemplateDnsLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDnsLogging) getPath() string{
    return "slb/template/dns-logging"
}

func (p *SlbTemplateDnsLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLogging::Post")
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

func (p *SlbTemplateDnsLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLogging::Get")
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
func (p *SlbTemplateDnsLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLogging::Put")
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

func (p *SlbTemplateDnsLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
