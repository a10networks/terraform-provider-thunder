package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type SlbServerService struct {
	Inst struct {
		Action string `json:"action" dval:"enable"`

		HealthCheck string `json:"health-check"`

		HealthCheckDisable int `json:"health-check-disable"`

		Label string `json:"label"`

		PacketCaptureTemplate string `json:"packet-capture-template"`

		PortNumber int `json:"port-number"`

		Protocol string `json:"protocol"`

		SamplingEnable []SlbServerServiceSamplingEnable `json:"sampling-enable"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Server_name string
	} `json:"service"`
}

type SlbServerServiceSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *SlbServerService) GetId() string {
	return strconv.Itoa(p.Inst.PortNumber) + "+" + p.Inst.Protocol + "+" + p.Inst.Label
}

func (p *SlbServerService) getPath() string {
	return "slb/server/" + p.Inst.Server_name + "/service"
}

func (p *SlbServerService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerService::Post")
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

func (p *SlbServerService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerService::Get")
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
func (p *SlbServerService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerService::Put")
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

func (p *SlbServerService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerService::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
