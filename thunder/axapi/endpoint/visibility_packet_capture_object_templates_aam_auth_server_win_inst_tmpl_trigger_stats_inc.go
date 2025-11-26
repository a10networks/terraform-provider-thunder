package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc struct {
	Inst struct {
		Krb_other_error int `json:"krb_other_error"`

		Krb_pw_change_failure int `json:"krb_pw_change_failure"`

		Krb_pw_expiry int `json:"krb_pw_expiry"`

		Krb_timeout_error int `json:"krb_timeout_error"`

		Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`

		Uuid string `json:"uuid"`

		Aam_auth_server_win_inst_tmpl_name string
	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) getPath() string {
	return "visibility/packet-capture/object-templates/aam-auth-server-win-inst-tmpl/" + p.Inst.Aam_auth_server_win_inst_tmpl_name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
