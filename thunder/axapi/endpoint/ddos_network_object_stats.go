package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectStats struct {
	IpList []DdosNetworkObjectStatsIpList `json:"ip-list"`

	Ipv6List []DdosNetworkObjectStatsIpv6List `json:"ipv6-list"`

	ObjectName string `json:"object-name"`

	Stats DdosNetworkObjectStatsStats `json:"stats"`
}
type DataDdosNetworkObjectStats struct {
	DtDdosNetworkObjectStats DdosNetworkObjectStats `json:"network-object"`
}

type DdosNetworkObjectStatsIpList struct {
	SubnetIpAddr string                            `json:"subnet-ip-addr"`
	Stats        DdosNetworkObjectStatsIpListStats `json:"stats"`
}

type DdosNetworkObjectStatsIpListStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

type DdosNetworkObjectStatsIpv6List struct {
	SubnetIpv6Addr string                              `json:"subnet-ipv6-addr"`
	Stats          DdosNetworkObjectStatsIpv6ListStats `json:"stats"`
}

type DdosNetworkObjectStatsIpv6ListStats struct {
	Packet_rate int `json:"packet_rate"`
	Bit_rate    int `json:"bit_rate"`
}

type DdosNetworkObjectStatsStats struct {
	Subnet_learned           int `json:"subnet_learned"`
	Subnet_aged              int `json:"subnet_aged"`
	Subnet_create_fail       int `json:"subnet_create_fail"`
	Ip_learned               int `json:"ip_learned"`
	Ip_aged                  int `json:"ip_aged"`
	Ip_create_fail           int `json:"ip_create_fail"`
	Service_learned          int `json:"service_learned"`
	Service_aged             int `json:"service_aged"`
	Service_create_fail      int `json:"service_create_fail"`
	Packet_rate              int `json:"packet_rate"`
	Bit_rate                 int `json:"bit_rate"`
	Topk_allocate_fail       int `json:"topk_allocate_fail"`
	Sport_learned            int `json:"sport_learned"`
	Sport_aged               int `json:"sport_aged"`
	Sport_create_fail        int `json:"sport_create_fail"`
	Agent_group_learned      int `json:"agent_group_learned"`
	Agent_group_aged         int `json:"agent_group_aged"`
	Agent_group_create_fail  int `json:"agent_group_create_fail"`
	Duplicate_sample_pkt_rcv int `json:"duplicate_sample_pkt_rcv"`
}

func (p *DdosNetworkObjectStats) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectStats) getPath() string {
	return "ddos/network-object/" + p.ObjectName + "/stats"
}

func (p *DdosNetworkObjectStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectStats, error) {
	logger.Println("DdosNetworkObjectStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectStats
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
