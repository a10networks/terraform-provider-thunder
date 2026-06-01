package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type NetworkVirtualWireHealthCheckOper struct {
	Oper NetworkVirtualWireHealthCheckOperOper `json:"oper"`

	Vlan int `json:"vlan"`
}
type DataNetworkVirtualWireHealthCheckOper struct {
	DtNetworkVirtualWireHealthCheckOper NetworkVirtualWireHealthCheckOper `json:"virtual-wire-health-check"`
}

type NetworkVirtualWireHealthCheckOperOper struct {
	EntryState string `json:"entry-state"`
	VlanState  string `json:"vlan-state"`
}

func (p *NetworkVirtualWireHealthCheckOper) GetId() string {
	return "1"
}

func (p *NetworkVirtualWireHealthCheckOper) getPath() string {
	return "network/virtual-wire-health-check/" + strconv.Itoa(p.Vlan) + "/oper"
}

func (p *NetworkVirtualWireHealthCheckOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireHealthCheckOper, error) {
	logger.Println("NetworkVirtualWireHealthCheckOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataNetworkVirtualWireHealthCheckOper
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
