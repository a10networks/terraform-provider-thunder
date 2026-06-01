package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type AamAuthenticationServerWindowsStats struct {
	InstanceList []AamAuthenticationServerWindowsStatsInstanceList `json:"instance-list"`

	Stats AamAuthenticationServerWindowsStatsStats `json:"stats"`
}
type DataAamAuthenticationServerWindowsStats struct {
	DtAamAuthenticationServerWindowsStats AamAuthenticationServerWindowsStats `json:"windows"`
}

type AamAuthenticationServerWindowsStatsInstanceList struct {
	Name  string                                               `json:"name"`
	Stats AamAuthenticationServerWindowsStatsInstanceListStats `json:"stats"`
}

type AamAuthenticationServerWindowsStatsInstanceListStats struct {
	Krb_send_req_success     int `json:"krb_send_req_success"`
	Krb_get_resp_success     int `json:"krb_get_resp_success"`
	Krb_timeout_error        int `json:"krb_timeout_error"`
	Krb_other_error          int `json:"krb_other_error"`
	Krb_pw_expiry            int `json:"krb_pw_expiry"`
	Krb_pw_change_success    int `json:"krb_pw_change_success"`
	Krb_pw_change_failure    int `json:"krb_pw_change_failure"`
	Krb_validate_kdc_success int `json:"krb_validate_kdc_success"`
	Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
}

type AamAuthenticationServerWindowsStatsStats struct {
	KerberosRequestSend              int `json:"kerberos-request-send"`
	KerberosResponseGet              int `json:"kerberos-response-get"`
	KerberosTimeoutError             int `json:"kerberos-timeout-error"`
	KerberosOtherError               int `json:"kerberos-other-error"`
	KerberosRequestNormal            int `json:"kerberos-request-normal"`
	KerberosRequestDropped           int `json:"kerberos-request-dropped"`
	KerberosResponseSuccess          int `json:"kerberos-response-success"`
	KerberosResponseFailure          int `json:"kerberos-response-failure"`
	KerberosResponseError            int `json:"kerberos-response-error"`
	KerberosResponseTimeout          int `json:"kerberos-response-timeout"`
	KerberosResponseOther            int `json:"kerberos-response-other"`
	KerberosJobStartError            int `json:"kerberos-job-start-error"`
	KerberosPollingControlError      int `json:"kerberos-polling-control-error"`
	KerberosPwExpiry                 int `json:"kerberos-pw-expiry"`
	KerberosPwChangeSuccess          int `json:"kerberos-pw-change-success"`
	KerberosPwChangeFailure          int `json:"kerberos-pw-change-failure"`
	KerberosValidateKdcSuccess       int `json:"kerberos-validate-kdc-success"`
	KerberosValidateKdcFailure       int `json:"kerberos-validate-kdc-failure"`
	KerberosGenerateKdcKeytabSuccess int `json:"kerberos-generate-kdc-keytab-success"`
	KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
	KerberosDeleteKdcKeytabSuccess   int `json:"kerberos-delete-kdc-keytab-success"`
	KerberosDeleteKdcKeytabFailure   int `json:"kerberos-delete-kdc-keytab-failure"`
	KerberosKdcKeytabCount           int `json:"kerberos-kdc-keytab-count"`
}

func (p *AamAuthenticationServerWindowsStats) GetId() string {
	return "1"
}

func (p *AamAuthenticationServerWindowsStats) getPath() string {
	return "aam/authentication/server/windows/stats"
}

func (p *AamAuthenticationServerWindowsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerWindowsStats, error) {
	logger.Println("AamAuthenticationServerWindowsStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataAamAuthenticationServerWindowsStats
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
