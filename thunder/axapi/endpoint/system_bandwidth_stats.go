package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemBandwidthStats struct {
	Stats SystemBandwidthStatsStats `json:"stats"`
}
type DataSystemBandwidthStats struct {
	DtSystemBandwidthStats SystemBandwidthStats `json:"bandwidth"`
}

type SystemBandwidthStatsStats struct {
	InputBytesPerSec  int `json:"input-bytes-per-sec"`
	OutputBytesPerSec int `json:"output-bytes-per-sec"`
	Ppsl_drop_egr     int `json:"ppsl_drop_egr"`
	Ppsl_drop_ing     int `json:"ppsl_drop_ing"`
	Ppsl_ignore_limit int `json:"ppsl_ignore_limit"`
	Licexpire_drop    int `json:"licexpire_drop"`
	Bwl_drop          int `json:"bwl_drop"`
}

func (p *SystemBandwidthStats) GetId() string {
	return "1"
}

func (p *SystemBandwidthStats) getPath() string {
	return "system/bandwidth/stats"
}

func (p *SystemBandwidthStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemBandwidthStats, error) {
	logger.Println("SystemBandwidthStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSystemBandwidthStats
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
