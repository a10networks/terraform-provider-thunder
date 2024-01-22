

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerWindowsStats struct {
    
    InstanceList []AamAuthenticationServerWindowsStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationServerWindowsStatsStats `json:"stats"`

}
type DataAamAuthenticationServerWindowsStats struct {
    DtAamAuthenticationServerWindowsStats AamAuthenticationServerWindowsStats `json:"windows"`
}


type AamAuthenticationServerWindowsStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationServerWindowsStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationServerWindowsStatsInstanceListStats struct {
    Krb_send_req_success int `json:"krb_send_req_success"`
    Krb_get_resp_success int `json:"krb_get_resp_success"`
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_success int `json:"krb_pw_change_success"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_success int `json:"ntlm_proto_nego_success"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_success int `json:"ntlm_session_setup_success"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_success int `json:"ntlm_prepare_req_success"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_success int `json:"ntlm_auth_success"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_success int `json:"krb_validate_kdc_success"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
}


type AamAuthenticationServerWindowsStatsStats struct {
    KerberosRequestSend int `json:"kerberos-request-send"`
    KerberosResponseGet int `json:"kerberos-response-get"`
    KerberosTimeoutError int `json:"kerberos-timeout-error"`
    KerberosOtherError int `json:"kerberos-other-error"`
    NtlmAuthenticationSuccess int `json:"ntlm-authentication-success"`
    NtlmAuthenticationFailure int `json:"ntlm-authentication-failure"`
    NtlmProtoNegotiationSuccess int `json:"ntlm-proto-negotiation-success"`
    NtlmProtoNegotiationFailure int `json:"ntlm-proto-negotiation-failure"`
    NtlmSessionSetupSuccess int `json:"ntlm-session-setup-success"`
    NtlmSessionSetupFailed int `json:"ntlm-session-setup-failed"`
    KerberosRequestNormal int `json:"kerberos-request-normal"`
    KerberosRequestDropped int `json:"kerberos-request-dropped"`
    KerberosResponseSuccess int `json:"kerberos-response-success"`
    KerberosResponseFailure int `json:"kerberos-response-failure"`
    KerberosResponseError int `json:"kerberos-response-error"`
    KerberosResponseTimeout int `json:"kerberos-response-timeout"`
    KerberosResponseOther int `json:"kerberos-response-other"`
    KerberosJobStartError int `json:"kerberos-job-start-error"`
    KerberosPollingControlError int `json:"kerberos-polling-control-error"`
    NtlmPrepareReqSuccess int `json:"ntlm-prepare-req-success"`
    NtlmPrepareReqFailed int `json:"ntlm-prepare-req-failed"`
    NtlmTimeoutError int `json:"ntlm-timeout-error"`
    NtlmOtherError int `json:"ntlm-other-error"`
    NtlmRequestNormal int `json:"ntlm-request-normal"`
    NtlmRequestDropped int `json:"ntlm-request-dropped"`
    NtlmResponseSuccess int `json:"ntlm-response-success"`
    NtlmResponseFailure int `json:"ntlm-response-failure"`
    NtlmResponseError int `json:"ntlm-response-error"`
    NtlmResponseTimeout int `json:"ntlm-response-timeout"`
    NtlmResponseOther int `json:"ntlm-response-other"`
    NtlmJobStartError int `json:"ntlm-job-start-error"`
    NtlmPollingControlError int `json:"ntlm-polling-control-error"`
    KerberosPwExpiry int `json:"kerberos-pw-expiry"`
    KerberosPwChangeSuccess int `json:"kerberos-pw-change-success"`
    KerberosPwChangeFailure int `json:"kerberos-pw-change-failure"`
    KerberosValidateKdcSuccess int `json:"kerberos-validate-kdc-success"`
    KerberosValidateKdcFailure int `json:"kerberos-validate-kdc-failure"`
    KerberosGenerateKdcKeytabSuccess int `json:"kerberos-generate-kdc-keytab-success"`
    KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
    KerberosDeleteKdcKeytabSuccess int `json:"kerberos-delete-kdc-keytab-success"`
    KerberosDeleteKdcKeytabFailure int `json:"kerberos-delete-kdc-keytab-failure"`
    KerberosKdcKeytabCount int `json:"kerberos-kdc-keytab-count"`
}

func (p *AamAuthenticationServerWindowsStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerWindowsStats) getPath() string{
    return "aam/authentication/server/windows/stats"
}

func (p *AamAuthenticationServerWindowsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerWindowsStats,error) {
logger.Println("AamAuthenticationServerWindowsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerWindowsStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
