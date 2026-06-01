package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemConfigMgmtNotification struct {
	Inst struct {
		Period int `json:"period" dval:"15"`

		Uuid string `json:"uuid"`
	} `json:"notification"`
}

func (p *SystemConfigMgmtNotification) GetId() string {
	return "1"
}

func (p *SystemConfigMgmtNotification) getPath() string {
	return "system/config-mgmt/notification"
}

func (p *SystemConfigMgmtNotification) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtNotification::Post")
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

func (p *SystemConfigMgmtNotification) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtNotification::Get")
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
func (p *SystemConfigMgmtNotification) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtNotification::Put")
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

func (p *SystemConfigMgmtNotification) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtNotification::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
