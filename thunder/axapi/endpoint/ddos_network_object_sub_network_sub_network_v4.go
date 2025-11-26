package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectSubNetworkSubNetworkV4 struct {
	Inst struct {
		BreakdownSubnetThreshold DdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold `json:"breakdown-subnet-threshold"`

		HostAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold `json:"host-anomaly-threshold"`

		SamplingEnable []DdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable `json:"sampling-enable"`

		SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`

		SubnetBreakdown int `json:"subnet-breakdown"`

		SubnetIpAddr string `json:"subnet-ip-addr"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"sub-network-v4"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold struct {
	BreakdownSubnetPktRate int `json:"breakdown-subnet-pkt-rate"`
	BreakdownSubnetBitRate int `json:"breakdown-subnet-bit-rate"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold struct {
	StaticPktRateThreshold                 int `json:"static-pkt-rate-threshold"`
	StaticRevPktRateThreshold              int `json:"static-rev-pkt-rate-threshold"`
	StaticBitRateThreshold                 int `json:"static-bit-rate-threshold"`
	StaticRevBitRateThreshold              int `json:"static-rev-bit-rate-threshold"`
	StaticUndiscoveredPktRateThreshold     int `json:"static-undiscovered-pkt-rate-threshold"`
	StaticFlowCountThreshold               int `json:"static-flow-count-threshold"`
	StaticSynRateThreshold                 int `json:"static-syn-rate-threshold"`
	StaticFinRateThreshold                 int `json:"static-fin-rate-threshold"`
	StaticRstRateThreshold                 int `json:"static-rst-rate-threshold"`
	StaticTcpPktRateThreshold              int `json:"static-tcp-pkt-rate-threshold"`
	StaticUdpPktRateThreshold              int `json:"static-udp-pkt-rate-threshold"`
	StaticIcmpPktRateThreshold             int `json:"static-icmp-pkt-rate-threshold"`
	StaticUndiscoveredHostPktRateThreshold int `json:"static-undiscovered-host-pkt-rate-threshold"`
	StaticUndiscoveredHostBitRateThreshold int `json:"static-undiscovered-host-bit-rate-threshold"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold struct {
	StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
	StaticSubNetworkBitRate int `json:"static-sub-network-bit-rate"`
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4) GetId() string {
	return url.QueryEscape(p.Inst.SubnetIpAddr)
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/sub-network/sub-network-v4"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV4::Post")
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

func (p *DdosNetworkObjectSubNetworkSubNetworkV4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV4::Get")
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
func (p *DdosNetworkObjectSubNetworkSubNetworkV4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV4::Put")
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

func (p *DdosNetworkObjectSubNetworkSubNetworkV4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV4::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
