package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectSubNetworkSubNetworkV6Stats struct {
	Stats DdosNetworkObjectSubNetworkSubNetworkV6StatsStats `json:"stats"`

	SubnetIpv6Addr string `json:"subnet-ipv6-addr"`

	ObjectName string
}
type DataDdosNetworkObjectSubNetworkSubNetworkV6Stats struct {
	DtDdosNetworkObjectSubNetworkSubNetworkV6Stats DdosNetworkObjectSubNetworkSubNetworkV6Stats `json:"sub-network-v6"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6StatsStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6Stats) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6Stats) getPath() string {

	return "ddos/network-object/" + p.ObjectName + "/sub-network/sub-network-v6/" + p.SubnetIpv6Addr + "/stats"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectSubNetworkSubNetworkV6Stats, error) {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV6Stats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectSubNetworkSubNetworkV6Stats
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
