package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc struct {
	Inst struct {
		JsonFail int `json:"json-fail"`

		OtherError int `json:"other-error"`

		ParseFail int `json:"parse-fail"`

		TimeoutError int `json:"timeout-error"`

		Uuid string `json:"uuid"`

		Captcha_template_inst_tmpl_name string
	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) getPath() string {
	return "visibility/packet-capture/object-templates/captcha-template-inst-tmpl/" + p.Inst.Captcha_template_inst_tmpl_name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
