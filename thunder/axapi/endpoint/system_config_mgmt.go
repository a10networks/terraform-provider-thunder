package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemConfigMgmt struct {
	Inst struct {
		DeleteReferencedTaggedObjects string `json:"delete-referenced-tagged-objects" dval:"enable"`

		Mpm SystemConfigMgmtMpm1674 `json:"mpm"`

		Notification SystemConfigMgmtNotification1675 `json:"notification"`

		PuSyncDetection SystemConfigMgmtPuSyncDetection1676 `json:"pu-sync-detection"`

		Uuid string `json:"uuid"`
	} `json:"config-mgmt"`
}

type SystemConfigMgmtMpm1674 struct {
	MaxWorkers     int    `json:"max-workers" dval:"1"`
	MinIdleWorkers int    `json:"min-idle-workers" dval:"1"`
	StartWorkers   int    `json:"start-workers" dval:"1"`
	Uuid           string `json:"uuid"`
}

type SystemConfigMgmtNotification1675 struct {
	Period int    `json:"period" dval:"15"`
	Uuid   string `json:"uuid"`
}

type SystemConfigMgmtPuSyncDetection1676 struct {
	Interval int    `json:"interval" dval:"30"`
	Action   string `json:"action" dval:"disable"`
	Uuid     string `json:"uuid"`
}

func (p *SystemConfigMgmt) GetId() string {
	return "1"
}

func (p *SystemConfigMgmt) getPath() string {
	return "system/config-mgmt"
}

func (p *SystemConfigMgmt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmt::Post")
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

func (p *SystemConfigMgmt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmt::Get")
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
func (p *SystemConfigMgmt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmt::Put")
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

func (p *SystemConfigMgmt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemConfigMgmt::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
