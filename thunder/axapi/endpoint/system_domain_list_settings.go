package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemDomainListSettings struct {
	Inst struct {
		ConcurrentTask int `json:"concurrent-task" dval:"6"`

		DomainListPerGroup string `json:"domain-list-per-group" dval:"16"`

		PollingInterval string `json:"polling-interval" dval:"10-second"`

		Uuid string `json:"uuid"`
	} `json:"domain-list-settings"`
}

func (p *SystemDomainListSettings) GetId() string {
	return "1"
}

func (p *SystemDomainListSettings) getPath() string {
	return "system/domain-list-settings"
}

func (p *SystemDomainListSettings) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemDomainListSettings::Post")
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

func (p *SystemDomainListSettings) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemDomainListSettings::Get")
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
func (p *SystemDomainListSettings) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemDomainListSettings::Put")
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

func (p *SystemDomainListSettings) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemDomainListSettings::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
