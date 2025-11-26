package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ConfigReplace struct {
	Inst struct {
		Action string `json:"action"`

		Filter string `json:"filter"`

		GslbSyncingOff int `json:"gslb-syncing-off"`

		IgnoreError int `json:"ignore-error"`

		LogError int `json:"log-error"`
	} `json:"config-replace"`
}

func (p *ConfigReplace) GetId() string {
	return "1"
}

func (p *ConfigReplace) getPath() string {
	return "config-replace"
}

func (p *ConfigReplace) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ConfigReplace::Post")
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

func (p *ConfigReplace) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ConfigReplace::Get")
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
func (p *ConfigReplace) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ConfigReplace::Put")
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

func (p *ConfigReplace) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ConfigReplace::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
