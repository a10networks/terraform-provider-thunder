package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SystemIp struct {
	Inst struct {
		IcmpRedirectDisable int `json:"icmp-redirect-disable"`

		IcmpUnreachableDisable int `json:"icmp-unreachable-disable"`

		RpfCheckEnable int `json:"rpf-check-enable"`

		SourceRoutePktDropEnable int `json:"source-route-pkt-drop-enable"`

		Uuid string `json:"uuid"`
	} `json:"ip"`
}

func (p *SystemIp) GetId() string {
	return "1"
}

func (p *SystemIp) getPath() string {
	return "system/ip"
}

func (p *SystemIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIp::Post")
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

func (p *SystemIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIp::Get")
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
func (p *SystemIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIp::Put")
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

func (p *SystemIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
