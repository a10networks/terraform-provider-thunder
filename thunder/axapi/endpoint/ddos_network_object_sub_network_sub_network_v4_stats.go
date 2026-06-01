package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectSubNetworkSubNetworkV4Stats struct {
	Stats DdosNetworkObjectSubNetworkSubNetworkV4StatsStats `json:"stats"`

	SubnetIpAddr string `json:"subnet-ip-addr"`

	ObjectName string
}
type DataDdosNetworkObjectSubNetworkSubNetworkV4Stats struct {
	DtDdosNetworkObjectSubNetworkSubNetworkV4Stats DdosNetworkObjectSubNetworkSubNetworkV4Stats `json:"sub-network-v4"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4StatsStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4Stats) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4Stats) getPath() string {

	return "ddos/network-object/" + p.ObjectName + "/sub-network/sub-network-v4/" + p.SubnetIpAddr + "/stats"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectSubNetworkSubNetworkV4Stats, error) {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV4Stats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectSubNetworkSubNetworkV4Stats
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
