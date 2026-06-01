package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectIpv6Stats struct {
	Stats DdosNetworkObjectIpv6StatsStats `json:"stats"`

	SubnetIpv6Addr string `json:"subnet-ipv6-addr"`

	ObjectName string
}
type DataDdosNetworkObjectIpv6Stats struct {
	DtDdosNetworkObjectIpv6Stats DdosNetworkObjectIpv6Stats `json:"ipv6"`
}

type DdosNetworkObjectIpv6StatsStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

func (p *DdosNetworkObjectIpv6Stats) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectIpv6Stats) getPath() string {

	return "ddos/network-object/" + p.ObjectName + "/ipv6/" + p.SubnetIpv6Addr + "/stats"
}

func (p *DdosNetworkObjectIpv6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectIpv6Stats, error) {
	logger.Println("DdosNetworkObjectIpv6Stats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectIpv6Stats
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
