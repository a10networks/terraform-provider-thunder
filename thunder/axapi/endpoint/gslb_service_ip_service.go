package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type GslbServiceIpService struct {
	Inst struct {
		Action string `json:"action" dval:"enable"`

		HealthCheck string `json:"health-check"`

		HealthCheckDisable int `json:"health-check-disable"`

		HealthCheckProtocolDisable int `json:"health-check-protocol-disable"`

		Label string `json:"label"`

		PortNum int `json:"port-num"`

		PortProto string `json:"port-proto"`

		SamplingEnable []GslbServiceIpServiceSamplingEnable `json:"sampling-enable"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		NodeName string
	} `json:"service"`
}

type GslbServiceIpServiceSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *GslbServiceIpService) GetId() string {
	return strconv.Itoa(p.Inst.PortNum) + "+" + p.Inst.PortProto + "+" + p.Inst.Label
}

func (p *GslbServiceIpService) getPath() string {
	return "gslb/service-ip/" + p.Inst.NodeName + "/service"
}

func (p *GslbServiceIpService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceIpService::Post")
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

func (p *GslbServiceIpService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceIpService::Get")
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
func (p *GslbServiceIpService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceIpService::Put")
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

func (p *GslbServiceIpService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceIpService::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
