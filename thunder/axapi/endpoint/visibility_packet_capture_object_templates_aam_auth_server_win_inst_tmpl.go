

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"aam-auth-server-win-inst-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654 struct {
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656 struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/aam-auth-server-win-inst-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
