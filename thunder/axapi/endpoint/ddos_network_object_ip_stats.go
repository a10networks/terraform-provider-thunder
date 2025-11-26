package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectIpStats struct {
	Stats DdosNetworkObjectIpStatsStats `json:"stats"`

	SubnetIpAddr string `json:"subnet-ip-addr"`

	ObjectName string
}
type DataDdosNetworkObjectIpStats struct {
	DtDdosNetworkObjectIpStats DdosNetworkObjectIpStats `json:"ip"`
}

type DdosNetworkObjectIpStatsStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

func (p *DdosNetworkObjectIpStats) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectIpStats) getPath() string {

	return "ddos/network-object/" + p.ObjectName + "/ip/" + p.SubnetIpAddr + "/stats"
}

func (p *DdosNetworkObjectIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectIpStats, error) {
	logger.Println("DdosNetworkObjectIpStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectIpStats
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
