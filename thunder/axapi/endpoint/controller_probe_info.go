package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ControllerProbeInfo struct {
	Inst struct {
		SamplingEnable []ControllerProbeInfoSamplingEnable `json:"sampling-enable"`

		Uuid string `json:"uuid"`
	} `json:"probe-info"`
}

type ControllerProbeInfoSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *ControllerProbeInfo) GetId() string {
	return "1"
}

func (p *ControllerProbeInfo) getPath() string {
	return "controller/probe-info"
}

func (p *ControllerProbeInfo) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProbeInfo::Post")
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

func (p *ControllerProbeInfo) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProbeInfo::Get")
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
func (p *ControllerProbeInfo) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProbeInfo::Put")
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

func (p *ControllerProbeInfo) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProbeInfo::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
