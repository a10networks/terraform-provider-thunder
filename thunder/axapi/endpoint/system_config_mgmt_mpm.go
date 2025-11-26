package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SystemConfigMgmtMpm struct {
	Inst struct {
		MaxWorkers int `json:"max-workers" dval:"1"`

		MinIdleWorkers int `json:"min-idle-workers" dval:"1"`

		StartWorkers int `json:"start-workers" dval:"1"`

		Uuid string `json:"uuid"`
	} `json:"mpm"`
}

func (p *SystemConfigMgmtMpm) GetId() string {
	return "1"
}

func (p *SystemConfigMgmtMpm) getPath() string {
	return "system/config-mgmt/mpm"
}

func (p *SystemConfigMgmtMpm) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtMpm::Post")
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

func (p *SystemConfigMgmtMpm) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtMpm::Get")
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
func (p *SystemConfigMgmtMpm) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtMpm::Put")
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

func (p *SystemConfigMgmtMpm) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtMpm::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
