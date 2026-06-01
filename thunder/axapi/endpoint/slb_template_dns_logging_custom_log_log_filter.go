package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsLoggingCustomLogLogFilter struct {
	Inst struct {
		Feature string `json:"feature"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		TriggerReason string

		Dns_logging_name string
	} `json:"log-filter"`
}

func (p *SlbTemplateDnsLoggingCustomLogLogFilter) GetId() string {
	return p.Inst.Feature
}

func (p *SlbTemplateDnsLoggingCustomLogLogFilter) getPath() string {
	return "slb/template/dns-logging/" + p.Inst.Dns_logging_name + "/custom-log/" + p.Inst.TriggerReason + "/log-filter"
}

func (p *SlbTemplateDnsLoggingCustomLogLogFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLogLogFilter::Post")
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

func (p *SlbTemplateDnsLoggingCustomLogLogFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLogLogFilter::Get")
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
func (p *SlbTemplateDnsLoggingCustomLogLogFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLogLogFilter::Put")
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

func (p *SlbTemplateDnsLoggingCustomLogLogFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLogLogFilter::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
