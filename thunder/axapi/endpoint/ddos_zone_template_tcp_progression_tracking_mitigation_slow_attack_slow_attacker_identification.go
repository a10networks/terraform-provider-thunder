package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification struct {
	Inst struct {
		ActiveConnection int `json:"active-connection" dval:"3"`

		BadConnection int `json:"bad-connection" dval:"75"`

		EnableIdentification int `json:"enable-identification"`

		Uuid string `json:"uuid"`

		Tcp_name string
	} `json:"slow-attacker-identification"`
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) GetId() string {
	return "1"
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) getPath() string {
	return "ddos/zone-template/tcp/" + p.Inst.Tcp_name + "/progression-tracking/mitigation/slow-attack/slow-attacker-identification"
}

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification::Post")
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

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification::Get")
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
func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification::Put")
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

func (p *DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
