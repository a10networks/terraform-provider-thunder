package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosNetworkIpEntityStats struct {
	Stats DdosNetworkIpEntityStatsStats `json:"stats"`
}
type DataDdosNetworkIpEntityStats struct {
	DtDdosNetworkIpEntityStats DdosNetworkIpEntityStats `json:"network-ip-entity"`
}

type DdosNetworkIpEntityStatsStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

func (p *DdosNetworkIpEntityStats) GetId() string {
	return "1"
}

func (p *DdosNetworkIpEntityStats) getPath() string {
	return "ddos/network-ip-entity/stats"
}

func (p *DdosNetworkIpEntityStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkIpEntityStats, error) {
	logger.Println("DdosNetworkIpEntityStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkIpEntityStats
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
