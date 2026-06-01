package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate struct {
	Inst struct {
		Duration int `json:"duration" dval:"60"`

		SynCookieVerificationFailed int `json:"syn-cookie-verification-failed"`

		ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`

		Uuid string `json:"uuid"`

		Rule_set_rule_tmpl_name string
	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) getPath() string {
	return "visibility/packet-capture/object-templates/rule-set-rule-tmpl/" + p.Inst.Rule_set_rule_tmpl_name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsRate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
