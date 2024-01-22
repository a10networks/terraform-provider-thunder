

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DnssecTemplate struct {
	Inst struct {

    Algorithm string `json:"algorithm"`
    CombinationsLimit int `json:"combinations-limit"`
    DnskeyTtlK int `json:"dnskey-ttl-k"`
    DnskeyTtlV int `json:"dnskey-ttl-v" dval:"14400"`
    DnssecTempName string `json:"dnssec-temp-name"`
    DnssecTemplateKsk DnssecTemplateDnssecTemplateKsk `json:"dnssec-template-ksk"`
    DnssecTemplateZsk DnssecTemplateDnssecTemplateZsk `json:"dnssec-template-zsk"`
    EnableNsec3 int `json:"enable-nsec3"`
    Hsm string `json:"hsm"`
    ReturnNsecOnFailure int `json:"return-nsec-on-failure" dval:"1"`
    SignatureValidityPeriodK int `json:"signature-validity-period-k"`
    SignatureValidityPeriodV int `json:"signature-validity-period-v" dval:"10"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type DnssecTemplateDnssecTemplateKsk struct {
    KskKeysizeK int `json:"ksk-keysize-k"`
    KskKeysizeV int `json:"ksk-keysize-v"`
    KskLifetimeK int `json:"ksk-lifetime-k"`
    KskLifetimeV int `json:"ksk-lifetime-v"`
    KskRolloverTimeK int `json:"ksk-rollover-time-k"`
    ZskRolloverTimeV int `json:"zsk-rollover-time-v" dval:"358"`
}


type DnssecTemplateDnssecTemplateZsk struct {
    ZskKeysizeK int `json:"zsk-keysize-k"`
    ZskKeysizeV int `json:"zsk-keysize-v"`
    ZskLifetimeK int `json:"zsk-lifetime-k"`
    ZskLifetimeV int `json:"zsk-lifetime-v" dval:"90"`
    ZskRolloverTimeK int `json:"zsk-rollover-time-k"`
    ZskRolloverTimeV int `json:"zsk-rollover-time-v" dval:"83"`
}

func (p *DnssecTemplate) GetId() string{
    return p.Inst.DnssecTempName
}

func (p *DnssecTemplate) getPath() string{
    return "dnssec/template"
}

func (p *DnssecTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecTemplate::Post")
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

func (p *DnssecTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecTemplate::Get")
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
func (p *DnssecTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecTemplate::Put")
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

func (p *DnssecTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
