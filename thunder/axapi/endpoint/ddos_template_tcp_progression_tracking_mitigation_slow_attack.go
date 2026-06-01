package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosTemplateTcpProgressionTrackingMitigationSlowAttack struct {
	Inst struct {
		InitRequestMaxTime int `json:"init-request-max-time"`

		InitResponseMaxTime int `json:"init-response-max-time"`

		ProgressionTrackingSlowAction string `json:"progression-tracking-slow-action" dval:"drop"`

		ProgressionTrackingSlowActionListName string `json:"progression-tracking-slow-action-list-name"`

		ResponsePktRateMax int `json:"response-pkt-rate-max"`

		Uuid string `json:"uuid"`

		Tcp_name string
	} `json:"slow-attack"`
}

func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) GetId() string {
	return "1"
}

func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) getPath() string {
	return "ddos/template/tcp/" + p.Inst.Tcp_name + "/progression-tracking/mitigation/slow-attack"
}

func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationSlowAttack::Post")
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

func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationSlowAttack::Get")
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
func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationSlowAttack::Put")
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

func (p *DdosTemplateTcpProgressionTrackingMitigationSlowAttack) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationSlowAttack::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
