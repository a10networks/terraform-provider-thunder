package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type VisibilityTopnTemplGtpPlcyTopnTmpl struct {
	Inst struct {
		Interval string `json:"interval"`

		Metrics VisibilityTopnTemplGtpPlcyTopnTmplMetrics3233 `json:"metrics"`

		Name string `json:"name"`

		TopnSize int `json:"topn-size"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"templ-gtp-plcy-topn-tmpl"`
}

type VisibilityTopnTemplGtpPlcyTopnTmplMetrics3233 struct {
	RlMessageMonitor int    `json:"rl-message-monitor"`
	Uuid             string `json:"uuid"`
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) getPath() string {
	return "visibility/topn/templ-gtp-plcy-topn-tmpl"
}

func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmpl::Post")
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

func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmpl::Get")
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
func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmpl::Put")
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

func (p *VisibilityTopnTemplGtpPlcyTopnTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityTopnTemplGtpPlcyTopnTmpl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
