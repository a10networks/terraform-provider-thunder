

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateSipMalformedSip struct {
	Inst struct {

    MalformedSipCallIdMaxLength int `json:"malformed-sip-call-id-max-length" dval:"32511"`
    MalformedSipCheck string `json:"malformed-sip-check"`
    MalformedSipMaxHeaderNameLength int `json:"malformed-sip-max-header-name-length" dval:"63"`
    MalformedSipMaxHeaderValueLength int `json:"malformed-sip-max-header-value-length" dval:"32511"`
    MalformedSipMaxLineSize int `json:"malformed-sip-max-line-size" dval:"32511"`
    MalformedSipMaxUriLength int `json:"malformed-sip-max-uri-length" dval:"32511"`
    MalformedSipSdpMaxLength int `json:"malformed-sip-sdp-max-length" dval:"32511"`
    Uuid string `json:"uuid"`
    SipTmplName string 

	} `json:"malformed-sip"`
}

func (p *DdosTemplateSipMalformedSip) GetId() string{
    return "1"
}

func (p *DdosTemplateSipMalformedSip) getPath() string{
    return "ddos/template/sip/" +p.Inst.SipTmplName + "/malformed-sip"
}

func (p *DdosTemplateSipMalformedSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipMalformedSip::Post")
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

func (p *DdosTemplateSipMalformedSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipMalformedSip::Get")
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
func (p *DdosTemplateSipMalformedSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipMalformedSip::Put")
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

func (p *DdosTemplateSipMalformedSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipMalformedSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
