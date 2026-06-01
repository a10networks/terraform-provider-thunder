package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemMonTemplateMonitoringMode struct {
	Inst struct {
		Mmode string `json:"mmode" dval:"interdependent"`

		Uuid string `json:"uuid"`
	} `json:"monitoring-mode"`
}

func (p *SystemMonTemplateMonitoringMode) GetId() string {
	return "1"
}

func (p *SystemMonTemplateMonitoringMode) getPath() string {
	return "system/mon-template/monitoring-mode"
}

func (p *SystemMonTemplateMonitoringMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemMonTemplateMonitoringMode::Post")
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

func (p *SystemMonTemplateMonitoringMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemMonTemplateMonitoringMode::Get")
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
func (p *SystemMonTemplateMonitoringMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemMonTemplateMonitoringMode::Put")
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

func (p *SystemMonTemplateMonitoringMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemMonTemplateMonitoringMode::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
