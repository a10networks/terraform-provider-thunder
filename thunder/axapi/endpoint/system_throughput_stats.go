package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type SystemThroughputStats struct {
	Stats SystemThroughputStatsStats `json:"stats"`
}

type SystemThroughputt struct {
	Throughput SystemThroughputStats `json:"throughput"`
}

type SystemThroughputStatsStats struct {
	GlobalSystemThroughputBitsPerSec int `json:"global-system-throughput-bits-per-sec"`
	PerPartThroughputBitsPerSec      int `json:"per-part-throughput-bits-per-sec"`
}

func (p *SystemThroughputStats) GetId() string {
	return "1"
}

func (p *SystemThroughputStats) getPath() string {
	return "system/throughput/stats"
}

func (p *SystemThroughputStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (SystemThroughputt, error) {
	logger.Println("SystemThroughputStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var st SystemThroughputt
	if err == nil {
		err = json.Unmarshal(axResp, &st)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return st, err
}
