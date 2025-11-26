package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SystemConfigMgmtPuSyncDetection struct {
	Inst struct {
		Action string `json:"action" dval:"disable"`

		Interval int `json:"interval" dval:"30"`

		Uuid string `json:"uuid"`
	} `json:"pu-sync-detection"`
}

func (p *SystemConfigMgmtPuSyncDetection) GetId() string {
	return "1"
}

func (p *SystemConfigMgmtPuSyncDetection) getPath() string {
	return "system/config-mgmt/pu-sync-detection"
}

func (p *SystemConfigMgmtPuSyncDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtPuSyncDetection::Post")
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

func (p *SystemConfigMgmtPuSyncDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtPuSyncDetection::Get")
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
func (p *SystemConfigMgmtPuSyncDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtPuSyncDetection::Put")
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

func (p *SystemConfigMgmtPuSyncDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmtPuSyncDetection::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
