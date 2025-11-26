package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc struct {
	Inst struct {
		SynCookieVerificationFailed int `json:"syn-cookie-verification-failed"`

		Uuid string `json:"uuid"`

		Rule_set_rule_tmpl_name string
	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) getPath() string {
	return "visibility/packet-capture/object-templates/rule-set-rule-tmpl/" + p.Inst.Rule_set_rule_tmpl_name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplTriggerStatsInc::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
