

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsLoggingResponseTypeType struct {
	Inst struct {

    CaaTypeLimitNum int `json:"caa-type-limit-num"`
    CaaTypeNoLimit int `json:"caa-type-no-limit"`
    Digest int `json:"digest"`
    DnskeyTypeLimitNum int `json:"dnskey-type-limit-num"`
    DnskeyTypeNoLimit int `json:"dnskey-type-no-limit"`
    DsTypeLimitNum int `json:"ds-type-limit-num"`
    DsTypeNoLimit int `json:"ds-type-no-limit"`
    LengthLimitFlag int `json:"length-limit-flag"`
    NaptrTypeLimitNum int `json:"naptr-type-limit-num"`
    NaptrTypeNoLimit int `json:"naptr-type-no-limit"`
    OptTypeLimitNum int `json:"opt-type-limit-num"`
    OptTypeNoLimit int `json:"opt-type-no-limit"`
    OtherData int `json:"other-data"`
    PublicKey int `json:"public-key"`
    RdataField int `json:"rdata-field"`
    ResponseTypeName string `json:"response-type-name"`
    RrsigTypeLimitNum int `json:"rrsig-type-limit-num"`
    RrsigTypeNoLimit int `json:"rrsig-type-no-limit"`
    ServiceField int `json:"service-field"`
    Signature int `json:"signature"`
    TsigTypeLimitNum int `json:"tsig-type-limit-num"`
    TsigTypeNoLimit int `json:"tsig-type-no-limit"`
    TxtData int `json:"txt-data"`
    TxtTypeLimitNum int `json:"txt-type-limit-num"`
    TxtTypeNoLimit int `json:"txt-type-no-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ValueField int `json:"value-field"`
    Name string 

	} `json:"type"`
}

func (p *SlbTemplateDnsLoggingResponseTypeType) GetId() string{
    return p.Inst.ResponseTypeName
}

func (p *SlbTemplateDnsLoggingResponseTypeType) getPath() string{
    return "slb/template/dns-logging/" +p.Inst.Name + "/response-type/type"
}

func (p *SlbTemplateDnsLoggingResponseTypeType) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLoggingResponseTypeType::Post")
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

func (p *SlbTemplateDnsLoggingResponseTypeType) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLoggingResponseTypeType::Get")
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
func (p *SlbTemplateDnsLoggingResponseTypeType) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLoggingResponseTypeType::Put")
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

func (p *SlbTemplateDnsLoggingResponseTypeType) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLoggingResponseTypeType::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
