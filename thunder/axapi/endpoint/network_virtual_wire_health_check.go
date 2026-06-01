package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type NetworkVirtualWireHealthCheck struct {
	Inst struct {
		ActiveThreshold int `json:"active-threshold"`

		Enable int `json:"enable"`

		Ethernet int `json:"ethernet"`

		GarpInterval int `json:"garp-interval"`

		InnerVlan int `json:"inner-vlan"`

		InnerVlanPacketCount int `json:"inner-vlan-packet-count"`

		Interval int `json:"interval"`

		L3Packet int `json:"l3-packet"`

		Method string `json:"method"`

		NexthopIp string `json:"nexthop-ip"`

		NexthopMac string `json:"nexthop-mac"`

		PartitionHealthCheck int `json:"partition-health-check"`

		SamplingEnable []NetworkVirtualWireHealthCheckSamplingEnable `json:"sampling-enable"`

		SbyEthernet int `json:"sby-ethernet"`

		SbyTrunk int `json:"sby-trunk"`

		SourceIp string `json:"source-ip"`

		SourceMac string `json:"source-mac"`

		Trunk int `json:"trunk"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Vlan int `json:"vlan"`
	} `json:"virtual-wire-health-check"`
}

type NetworkVirtualWireHealthCheckSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *NetworkVirtualWireHealthCheck) GetId() string {
	return strconv.Itoa(p.Inst.Vlan)
}

func (p *NetworkVirtualWireHealthCheck) getPath() string {
	return "network/virtual-wire-health-check"
}

func (p *NetworkVirtualWireHealthCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("NetworkVirtualWireHealthCheck::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *NetworkVirtualWireHealthCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("NetworkVirtualWireHealthCheck::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *NetworkVirtualWireHealthCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("NetworkVirtualWireHealthCheck::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *NetworkVirtualWireHealthCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("NetworkVirtualWireHealthCheck::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
