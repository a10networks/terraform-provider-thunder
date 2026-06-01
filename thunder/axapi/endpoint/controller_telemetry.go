package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ControllerTelemetry struct {
	Telemetry ControllerTelemetryInst `json:"telemetry"`
}

type ControllerTelemetryInst struct {
	LogRate int    `json:"log-rate"`
	Uuid    string `json:"uuid"`
}


func (p *ControllerTelemetry) GetId() string {
	return "1"
}

func (p *ControllerTelemetry) getPath() string {
	return "controller/telemetry"
}

func (p *ControllerTelemetry) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetry::Post")
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

func (p *ControllerTelemetry) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetry::Get")
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
func (p *ControllerTelemetry) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetry::Put")
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

func (p *ControllerTelemetry) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetry::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
