package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl struct {
	Inst struct {
		CaptureConfig string `json:"capture-config"`

		Name string `json:"name"`

		TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779 `json:"trigger-stats-inc"`

		TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780 `json:"trigger-stats-rate"`

		TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781 `json:"trigger-stats-severity"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"captcha-template-inst-tmpl"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779 struct {
	ParseFail    int    `json:"parse-fail"`
	JsonFail     int    `json:"json-fail"`
	TimeoutError int    `json:"timeout-error"`
	OtherError   int    `json:"other-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	ParseFail           int    `json:"parse-fail"`
	JsonFail            int    `json:"json-fail"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) getPath() string {
	return "visibility/packet-capture/object-templates/captcha-template-inst-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
