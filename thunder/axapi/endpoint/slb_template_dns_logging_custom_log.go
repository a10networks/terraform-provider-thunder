package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsLoggingCustomLog struct {
	Inst struct {
		Enable int `json:"enable"`

		Format string `json:"format"`

		LogFilterList []SlbTemplateDnsLoggingCustomLogLogFilterList `json:"log-filter-list"`

		TriggerReason string `json:"trigger-reason"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Dns_logging_name string
	} `json:"custom-log"`
}

type SlbTemplateDnsLoggingCustomLogLogFilterList struct {
	Feature string `json:"feature"`
	Uuid    string `json:"uuid"`
	UserTag string `json:"user-tag"`
}

func (p *SlbTemplateDnsLoggingCustomLog) GetId() string {
	return p.Inst.TriggerReason
}

func (p *SlbTemplateDnsLoggingCustomLog) getPath() string {
	return "slb/template/dns-logging/" + p.Inst.Dns_logging_name + "/custom-log"
}

func (p *SlbTemplateDnsLoggingCustomLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLog::Post")
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

func (p *SlbTemplateDnsLoggingCustomLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLog::Get")
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
func (p *SlbTemplateDnsLoggingCustomLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLog::Put")
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

func (p *SlbTemplateDnsLoggingCustomLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingCustomLog::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
