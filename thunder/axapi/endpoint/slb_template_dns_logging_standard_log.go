package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsLoggingStandardLog struct {
	Inst struct {
		LogFilterList []SlbTemplateDnsLoggingStandardLogLogFilterList `json:"log-filter-list"`

		TriggerReason string `json:"trigger-reason"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Dns_logging_name string
	} `json:"standard-log"`
}

type SlbTemplateDnsLoggingStandardLogLogFilterList struct {
	Feature string `json:"feature"`
	Uuid    string `json:"uuid"`
	UserTag string `json:"user-tag"`
}

func (p *SlbTemplateDnsLoggingStandardLog) GetId() string {
	return p.Inst.TriggerReason
}

func (p *SlbTemplateDnsLoggingStandardLog) getPath() string {
	return "slb/template/dns-logging/" + p.Inst.Dns_logging_name + "/standard-log"
}

func (p *SlbTemplateDnsLoggingStandardLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingStandardLog::Post")
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

func (p *SlbTemplateDnsLoggingStandardLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingStandardLog::Get")
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
func (p *SlbTemplateDnsLoggingStandardLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingStandardLog::Put")
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

func (p *SlbTemplateDnsLoggingStandardLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLoggingStandardLog::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
