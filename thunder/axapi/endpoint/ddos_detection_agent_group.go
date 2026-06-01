package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosDetectionAgentGroup struct {
	Inst struct {
		Agent []DdosDetectionAgentGroupAgent `json:"agent"`

		AgentGroupName string `json:"agent-group-name"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"agent-group"`
}

type DdosDetectionAgentGroupAgent struct {
	AgentName string `json:"agent-name"`
}

func (p *DdosDetectionAgentGroup) GetId() string {
	return url.QueryEscape(p.Inst.AgentGroupName)
}

func (p *DdosDetectionAgentGroup) getPath() string {
	return "ddos/detection/agent-group"
}

func (p *DdosDetectionAgentGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionAgentGroup::Post")
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

func (p *DdosDetectionAgentGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionAgentGroup::Get")
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
func (p *DdosDetectionAgentGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionAgentGroup::Put")
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

func (p *DdosDetectionAgentGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionAgentGroup::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
