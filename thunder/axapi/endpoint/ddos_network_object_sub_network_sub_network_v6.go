package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectSubNetworkSubNetworkV6 struct {
	Inst struct {
		HostAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold `json:"host-anomaly-threshold"`

		SamplingEnable []DdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable `json:"sampling-enable"`

		SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`

		SubnetBreakdown int `json:"subnet-breakdown"`

		SubnetIpv6Addr string `json:"subnet-ipv6-addr"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"sub-network-v6"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold struct {
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

type DdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold struct {
	StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
	StaticSubNetworkBitRate int `json:"static-sub-network-bit-rate"`
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6) GetId() string {
	return url.QueryEscape(p.Inst.SubnetIpv6Addr)
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/sub-network/sub-network-v6"
}

func (p *DdosNetworkObjectSubNetworkSubNetworkV6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV6::Post")
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

func (p *DdosNetworkObjectSubNetworkSubNetworkV6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV6::Get")
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
func (p *DdosNetworkObjectSubNetworkSubNetworkV6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV6::Put")
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

func (p *DdosNetworkObjectSubNetworkSubNetworkV6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSubNetworkSubNetworkV6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
