package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate struct {
	Inst struct {
		Duration int `json:"duration" dval:"60"`

		JsonFail int `json:"json-fail"`

		OtherError int `json:"other-error"`

		ParseFail int `json:"parse-fail"`

		ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`

		TimeoutError int `json:"timeout-error"`

		Uuid string `json:"uuid"`

		Captcha_template_inst_tmpl_name string
	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) getPath() string {
	return "visibility/packet-capture/object-templates/captcha-template-inst-tmpl/" + p.Inst.Captcha_template_inst_tmpl_name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
