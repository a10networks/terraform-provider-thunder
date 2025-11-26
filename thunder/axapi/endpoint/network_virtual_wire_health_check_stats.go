package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type NetworkVirtualWireHealthCheckStats struct {
	Stats NetworkVirtualWireHealthCheckStatsStats `json:"stats"`

	Vlan int `json:"vlan"`
}
type DataNetworkVirtualWireHealthCheckStats struct {
	DtNetworkVirtualWireHealthCheckStats NetworkVirtualWireHealthCheckStats `json:"virtual-wire-health-check"`
}

type NetworkVirtualWireHealthCheckStatsStats struct {
	ActEvent    int `json:"act-event"`
	SbyEvent    int `json:"sby-event"`
	PacketCount int `json:"packet-count"`
}

func (p *NetworkVirtualWireHealthCheckStats) GetId() string {
	return "1"
}

func (p *NetworkVirtualWireHealthCheckStats) getPath() string {
	return "network/virtual-wire-health-check/" + strconv.Itoa(p.Vlan) + "/stats"
}

func (p *NetworkVirtualWireHealthCheckStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireHealthCheckStats, error) {
	logger.Println("NetworkVirtualWireHealthCheckStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataNetworkVirtualWireHealthCheckStats
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
