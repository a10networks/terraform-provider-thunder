package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDetectionSettingsZoneNotifications struct {
	Inst struct {
		SourceEntry string `json:"source-entry" dval:"disable"`

		Uuid string `json:"uuid"`
	} `json:"zone-notifications"`
}

func (p *DdosDetectionSettingsZoneNotifications) GetId() string {
	return "1"
}

func (p *DdosDetectionSettingsZoneNotifications) getPath() string {
	return "ddos/detection/settings/zone-notifications"
}

func (p *DdosDetectionSettingsZoneNotifications) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettingsZoneNotifications::Post")
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

func (p *DdosDetectionSettingsZoneNotifications) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettingsZoneNotifications::Get")
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
func (p *DdosDetectionSettingsZoneNotifications) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettingsZoneNotifications::Put")
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

func (p *DdosDetectionSettingsZoneNotifications) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettingsZoneNotifications::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
