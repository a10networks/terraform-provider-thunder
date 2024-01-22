

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc1949 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate1950 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"aam-auth-server-win"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc1949 struct {
    KerberosTimeoutError int `json:"kerberos-timeout-error"`
    KerberosOtherError int `json:"kerberos-other-error"`
    NtlmAuthenticationFailure int `json:"ntlm-authentication-failure"`
    NtlmProtoNegotiationFailure int `json:"ntlm-proto-negotiation-failure"`
    NtlmSessionSetupFailed int `json:"ntlm-session-setup-failed"`
    KerberosRequestDropped int `json:"kerberos-request-dropped"`
    KerberosResponseFailure int `json:"kerberos-response-failure"`
    KerberosResponseError int `json:"kerberos-response-error"`
    KerberosResponseTimeout int `json:"kerberos-response-timeout"`
    KerberosJobStartError int `json:"kerberos-job-start-error"`
    KerberosPollingControlError int `json:"kerberos-polling-control-error"`
    NtlmPrepareReqFailed int `json:"ntlm-prepare-req-failed"`
    NtlmTimeoutError int `json:"ntlm-timeout-error"`
    NtlmOtherError int `json:"ntlm-other-error"`
    NtlmRequestDropped int `json:"ntlm-request-dropped"`
    NtlmResponseFailure int `json:"ntlm-response-failure"`
    NtlmResponseError int `json:"ntlm-response-error"`
    NtlmResponseTimeout int `json:"ntlm-response-timeout"`
    NtlmJobStartError int `json:"ntlm-job-start-error"`
    NtlmPollingControlError int `json:"ntlm-polling-control-error"`
    KerberosPwExpiry int `json:"kerberos-pw-expiry"`
    KerberosPwChangeFailure int `json:"kerberos-pw-change-failure"`
    KerberosValidateKdcFailure int `json:"kerberos-validate-kdc-failure"`
    KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
    KerberosDeleteKdcKeytabFailure int `json:"kerberos-delete-kdc-keytab-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate1950 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    KerberosTimeoutError int `json:"kerberos-timeout-error"`
    KerberosOtherError int `json:"kerberos-other-error"`
    NtlmAuthenticationFailure int `json:"ntlm-authentication-failure"`
    NtlmProtoNegotiationFailure int `json:"ntlm-proto-negotiation-failure"`
    NtlmSessionSetupFailed int `json:"ntlm-session-setup-failed"`
    KerberosRequestDropped int `json:"kerberos-request-dropped"`
    KerberosResponseFailure int `json:"kerberos-response-failure"`
    KerberosResponseError int `json:"kerberos-response-error"`
    KerberosResponseTimeout int `json:"kerberos-response-timeout"`
    KerberosJobStartError int `json:"kerberos-job-start-error"`
    KerberosPollingControlError int `json:"kerberos-polling-control-error"`
    NtlmPrepareReqFailed int `json:"ntlm-prepare-req-failed"`
    NtlmTimeoutError int `json:"ntlm-timeout-error"`
    NtlmOtherError int `json:"ntlm-other-error"`
    NtlmRequestDropped int `json:"ntlm-request-dropped"`
    NtlmResponseFailure int `json:"ntlm-response-failure"`
    NtlmResponseError int `json:"ntlm-response-error"`
    NtlmResponseTimeout int `json:"ntlm-response-timeout"`
    NtlmJobStartError int `json:"ntlm-job-start-error"`
    NtlmPollingControlError int `json:"ntlm-polling-control-error"`
    KerberosPwExpiry int `json:"kerberos-pw-expiry"`
    KerberosPwChangeFailure int `json:"kerberos-pw-change-failure"`
    KerberosValidateKdcFailure int `json:"kerberos-validate-kdc-failure"`
    KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
    KerberosDeleteKdcKeytabFailure int `json:"kerberos-delete-kdc-keytab-failure"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/aam-auth-server-win"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
