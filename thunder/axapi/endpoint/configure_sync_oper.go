package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ConfigureSyncOper struct {
	Oper ConfigureSyncOperOper `json:"oper"`
}
type DataConfigureSyncOper struct {
	DtConfigureSyncOper ConfigureSyncOper `json:"sync"`
}

type ConfigureSyncOperOper struct {
	AllPartitions  int                                   `json:"all-partitions"`
	ConfigSyncList []ConfigureSyncOperOperConfigSyncList `json:"config-sync-list"`
}

type ConfigureSyncOperOperConfigSyncList struct {
	PartitionName     string `json:"partition-name"`
	RunSyncStatus     string `json:"run-sync-status"`
	StartupSyncStatus string `json:"startup-sync-status"`
}

func (p *ConfigureSyncOper) GetId() string {
	return "1"
}

func (p *ConfigureSyncOper) getPath() string {
	return "configure/sync/oper"
}

func (p *ConfigureSyncOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataConfigureSyncOper, error) {
	logger.Println("ConfigureSyncOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataConfigureSyncOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
