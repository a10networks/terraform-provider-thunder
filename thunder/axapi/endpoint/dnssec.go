

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Dnssec struct {
	Inst struct {

    Dnskey DnssecDnskey343 `json:"dnskey"`
    Ds DnssecDs344 `json:"ds"`
    KeyRollover DnssecKeyRollover345 `json:"key-rollover"`
    SignZoneNow DnssecSignZoneNow346 `json:"sign-zone-now"`
    Standalone int `json:"standalone"`
    TemplateList []DnssecTemplateList `json:"template-list"`
    Uuid string `json:"uuid"`

	} `json:"dnssec"`
}


type DnssecDnskey343 struct {
    KeyDelete int `json:"key-delete"`
    ZoneName string `json:"zone-name"`
}


type DnssecDs344 struct {
    DsDelete int `json:"ds-delete"`
    ZoneName string `json:"zone-name"`
}


type DnssecKeyRollover345 struct {
    ZoneName string `json:"zone-name"`
    DnssecKeyType string `json:"dnssec-key-type"`
    ZskStart int `json:"zsk-start"`
    KskStart int `json:"ksk-start"`
    DsReadyInParentZone int `json:"ds-ready-in-parent-zone"`
}


type DnssecSignZoneNow346 struct {
    ZoneName string `json:"zone-name"`
}


type DnssecTemplateList struct {
    DnssecTempName string `json:"dnssec-temp-name"`
    Algorithm string `json:"algorithm"`
    CombinationsLimit int `json:"combinations-limit"`
    DnskeyTtlK int `json:"dnskey-ttl-k"`
    DnskeyTtlV int `json:"dnskey-ttl-v" dval:"14400"`
    EnableNsec3 int `json:"enable-nsec3"`
    ReturnNsecOnFailure int `json:"return-nsec-on-failure" dval:"1"`
    SignatureValidityPeriodK int `json:"signature-validity-period-k"`
    SignatureValidityPeriodV int `json:"signature-validity-period-v" dval:"10"`
    Hsm string `json:"hsm"`
    DnssecTemplateZsk DnssecTemplateListDnssecTemplateZsk `json:"dnssec-template-zsk"`
    DnssecTemplateKsk DnssecTemplateListDnssecTemplateKsk `json:"dnssec-template-ksk"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DnssecTemplateListDnssecTemplateZsk struct {
    ZskKeysizeK int `json:"zsk-keysize-k"`
    ZskKeysizeV int `json:"zsk-keysize-v"`
    ZskLifetimeK int `json:"zsk-lifetime-k"`
    ZskLifetimeV int `json:"zsk-lifetime-v" dval:"90"`
    ZskRolloverTimeK int `json:"zsk-rollover-time-k"`
    ZskRolloverTimeV int `json:"zsk-rollover-time-v" dval:"83"`
}


type DnssecTemplateListDnssecTemplateKsk struct {
    KskKeysizeK int `json:"ksk-keysize-k"`
    KskKeysizeV int `json:"ksk-keysize-v"`
    KskLifetimeK int `json:"ksk-lifetime-k"`
    KskLifetimeV int `json:"ksk-lifetime-v"`
    KskRolloverTimeK int `json:"ksk-rollover-time-k"`
    ZskRolloverTimeV int `json:"zsk-rollover-time-v" dval:"358"`
}

func (p *Dnssec) GetId() string{
    return "1"
}

func (p *Dnssec) getPath() string{
    return "dnssec"
}

func (p *Dnssec) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Dnssec::Post")
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

func (p *Dnssec) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Dnssec::Get")
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
func (p *Dnssec) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Dnssec::Put")
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

func (p *Dnssec) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Dnssec::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
