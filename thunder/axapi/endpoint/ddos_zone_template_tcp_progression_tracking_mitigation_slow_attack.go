package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack struct {
	Inst struct {
		InitRequestMaxTime int `json:"init-request-max-time"`

		InitResponseMaxTime int `json:"init-response-max-time"`

		ProgressionTrackingSlowAction string `json:"progression-tracking-slow-action" dval:"drop"`

		ProgressionTrackingSlowActionListName string `json:"progression-tracking-slow-action-list-name"`

		ResponsePktRateMax int `json:"response-pkt-rate-max"`

		SlowAttack string `json:"slow-attack"`

		SlowAttackerIdentification DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification369 `json:"slow-attacker-identification"`

		Uuid string `json:"uuid"`

		Tcp_name string
	} `json:"slow-attack"`
}

type DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification369 struct {
	EnableIdentification int    `json:"enable-identification"`
	ActiveConnection     int    `json:"active-connection" dval:"3"`
	BadConnection        int    `json:"bad-connection" dval:"75"`
	Uuid                 string `json:"uuid"`
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) GetId() string {
	return "1"
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) getPath() string {
	return "ddos/zone-template/tcp/" + p.Inst.Tcp_name + "/progression-tracking/mitigation/slow-attack"
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack::Post")
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

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack::Get")
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
func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack::Put")
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

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
