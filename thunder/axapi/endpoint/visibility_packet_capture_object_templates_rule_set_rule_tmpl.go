package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl struct {
	Inst struct {
		CaptureConfig string `json:"capture-config"`

		Name string `json:"name"`

		TriggerStatsInc VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc2821 `json:"trigger-stats-inc"`

		TriggerStatsRate VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate2822 `json:"trigger-stats-rate"`

		TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsSeverity2823 `json:"trigger-stats-severity"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"rule-set-rule-tmpl"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc2821 struct {
	SynCookieVerificationFailed int    `json:"syn-cookie-verification-failed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate2822 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	SynCookieVerificationFailed int    `json:"syn-cookie-verification-failed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsSeverity2823 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) getPath() string {
	return "visibility/packet-capture/object-templates/rule-set-rule-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmpl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
