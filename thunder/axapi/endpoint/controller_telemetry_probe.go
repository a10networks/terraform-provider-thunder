package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ControllerTelemetryProbe struct {
	Inst struct {
		Action string `json:"action" dval:"disable"`

		ExportPolicy string `json:"export-policy" dval:"snapshots-new"`

		Interval int `json:"interval" dval:"15"`

		LogLevel string `json:"log-level" dval:"ERROR"`

		Target string `json:"target" dval:"remote"`

		Uuid string `json:"uuid"`
	} `json:"probe"`
}

func (p *ControllerTelemetryProbe) GetId() string {
	return "1"
}

func (p *ControllerTelemetryProbe) getPath() string {
	return "controller/telemetry/probe"
}

func (p *ControllerTelemetryProbe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetryProbe::Post")
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

func (p *ControllerTelemetryProbe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetryProbe::Get")
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
func (p *ControllerTelemetryProbe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetryProbe::Put")
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

func (p *ControllerTelemetryProbe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerTelemetryProbe::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
