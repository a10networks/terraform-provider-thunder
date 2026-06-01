package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityTopnTemplGtpPlcyTopnTmplMetrics struct {
	Inst struct {
		RlMessageMonitor int `json:"rl-message-monitor"`

		Uuid string `json:"uuid"`

		Templ_gtp_plcy_topn_tmpl_name string
	} `json:"metrics"`
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) GetId() string {
	return "1"
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) getPath() string {
	return "visibility/topn/templ-gtp-plcy-topn-tmpl/" + p.Inst.Templ_gtp_plcy_topn_tmpl_name + "/metrics"
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmplMetrics::Post")
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

func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmplMetrics::Get")
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
func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmplMetrics::Put")
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

func (p *VisibilityTopnTemplGtpPlcyTopnTmplMetrics) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmplMetrics::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
