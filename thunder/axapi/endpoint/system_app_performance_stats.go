package endpoint

import (
	"encoding/json"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
)

// based on ACOS 6_0_0-P1_10
type SystemAppPerformanceStats struct {
	Stats SystemAppPerformanceStatsStats `json:"stats"`
}

type SystemAppPerformancee struct {
	AppPerformance SystemAppPerformanceStats `json:"app-performance"`
}

type SystemAppPerformanceStatsStats struct {
	TotalThroughputBitsPerSec int `json:"total-throughput-bits-per-sec"`
	L4ConnsPerSec             int `json:"l4-conns-per-sec"`
	L7ConnsPerSec             int `json:"l7-conns-per-sec"`
	L7TransPerSec             int `json:"l7-trans-per-sec"`
	SslConnsPerSec            int `json:"ssl-conns-per-sec"`
	IpNatConnsPerSec          int `json:"ip-nat-conns-per-sec"`
	TotalNewConnsPerSec       int `json:"total-new-conns-per-sec"`
	TotalCurrConns            int `json:"total-curr-conns"`
	L4Bandwidth               int `json:"l4-bandwidth"`
	L7Bandwidth               int `json:"l7-bandwidth"`
	ServSslConnsPerSec        int `json:"serv-ssl-conns-per-sec"`
	FwConnsPerSec             int `json:"fw-conns-per-sec"`
	GifwConnsPerSec           int `json:"gifw-conns-per-sec"`
}

func (p *SystemAppPerformanceStats) GetId() string {
	return "1"
}

func (p *SystemAppPerformanceStats) getPath() string {
	return "system/app-performance/stats"
}

func (p *SystemAppPerformanceStats) Get(authToken string, host string, logger *axapi.ThunderLog) (SystemAppPerformancee, error) {
	logger.Println("SystemAppPerformanceStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)

	var sp SystemAppPerformancee
	if err == nil {
		err = json.Unmarshal(axResp, &sp)

		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return sp, err
}
