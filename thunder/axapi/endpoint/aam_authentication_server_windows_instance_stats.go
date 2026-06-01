package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type AamAuthenticationServerWindowsInstanceStats struct {
	Name string `json:"name"`

	Stats AamAuthenticationServerWindowsInstanceStatsStats `json:"stats"`
}
type DataAamAuthenticationServerWindowsInstanceStats struct {
	DtAamAuthenticationServerWindowsInstanceStats AamAuthenticationServerWindowsInstanceStats `json:"instance"`
}

type AamAuthenticationServerWindowsInstanceStatsStats struct {
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

func (p *AamAuthenticationServerWindowsInstanceStats) GetId() string {
	return "1"
}

func (p *AamAuthenticationServerWindowsInstanceStats) getPath() string {
	return "aam/authentication/server/windows/instance/" + p.Name + "/stats"
}

func (p *AamAuthenticationServerWindowsInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerWindowsInstanceStats, error) {
	logger.Println("AamAuthenticationServerWindowsInstanceStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataAamAuthenticationServerWindowsInstanceStats
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
