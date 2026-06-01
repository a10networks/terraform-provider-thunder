package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosNetworkObject struct {
	Inst struct {
		AnomalyChildPercentage int `json:"anomaly-child-percentage"`

		AnomalyDetectionTrigger string `json:"anomaly-detection-trigger"`

		EnableTopK []DdosNetworkObjectEnableTopK `json:"enable-top-k"`

		FloodingMultiplier int `json:"flooding-multiplier"`

		HistogramMode string `json:"histogram-mode"`

		HostAnomalyThreshold DdosNetworkObjectHostAnomalyThreshold `json:"host-anomaly-threshold"`

		IndicatorsToMonitor DdosNetworkObjectIndicatorsToMonitor317 `json:"indicators-to-monitor"`

		IpList []DdosNetworkObjectIpList `json:"ip-list"`

		Ipv6List []DdosNetworkObjectIpv6List `json:"ipv6-list"`

		NetworkObjectAnomalyThreshold DdosNetworkObjectNetworkObjectAnomalyThreshold `json:"network-object-anomaly-threshold"`

		NetworkObjectTemplate string `json:"network-object-template"`

		Notification DdosNetworkObjectNotification318 `json:"notification"`

		ObjectName string `json:"object-name"`

		OperationalMode string `json:"operational-mode"`

		RelativeAutoBreakDownThreshold DdosNetworkObjectRelativeAutoBreakDownThreshold `json:"relative-auto-break-down-threshold"`

		SamplingEnable []DdosNetworkObjectSamplingEnable `json:"sampling-enable"`

		ServiceBreakDownThresholdLocal DdosNetworkObjectServiceBreakDownThresholdLocal `json:"service-break-down-threshold-local"`

		ServiceDiscovery string `json:"service-discovery"`

		SrcServiceDiscovery string `json:"src-service-discovery"`

		SrcServiceDiscoveryThreshold int `json:"src-service-discovery-threshold" dval:"10"`

		StaticAutoBreakDownThreshold DdosNetworkObjectStaticAutoBreakDownThreshold `json:"static-auto-break-down-threshold"`

		SubNetwork DdosNetworkObjectSubNetwork320 `json:"sub-network"`

		ThresholdSensitivity string `json:"threshold-sensitivity"`

		TopkDestinations DdosNetworkObjectTopkDestinations321 `json:"topk-destinations"`

		Trustlist DdosNetworkObjectTrustlist322 `json:"trustlist"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"network-object"`
}

type DdosNetworkObjectEnableTopK struct {
	TopkType          string `json:"topk-type"`
	TopkDstNumRecords int    `json:"topk-dst-num-records" dval:"20"`
	TopkSortKey       string `json:"topk-sort-key" dval:"average"`
}

type DdosNetworkObjectHostAnomalyThreshold struct {
	HostPktRate                 int `json:"host-pkt-rate"`
	HostBitRate                 int `json:"host-bit-rate"`
	HostRevPktRate              int `json:"host-rev-pkt-rate"`
	HostRevBitRate              int `json:"host-rev-bit-rate"`
	HostUndiscoveredPktRate     int `json:"host-undiscovered-pkt-rate"`
	HostFlowCount               int `json:"host-flow-count"`
	HostSynRate                 int `json:"host-syn-rate"`
	HostFinRate                 int `json:"host-fin-rate"`
	HostRstRate                 int `json:"host-rst-rate"`
	HostTcpPktRate              int `json:"host-tcp-pkt-rate"`
	HostUdpPktRate              int `json:"host-udp-pkt-rate"`
	HostIcmpPktRate             int `json:"host-icmp-pkt-rate"`
	HostUndiscoveredHostPktRate int `json:"host-undiscovered-host-pkt-rate"`
	HostUndiscoveredHostBitRate int `json:"host-undiscovered-host-bit-rate"`
}

type DdosNetworkObjectIndicatorsToMonitor317 struct {
	Enable                     int    `json:"enable"`
	MonitorPktRate             int    `json:"monitor-pkt-rate"`
	MonitorBitRate             int    `json:"monitor-bit-rate"`
	MonitorRevPktRate          int    `json:"monitor-rev-pkt-rate"`
	MonitorRevBitRate          int    `json:"monitor-rev-bit-rate"`
	MonitorUndiscoveredPktRate int    `json:"monitor-undiscovered-pkt-rate"`
	MonitorFlowCount           int    `json:"monitor-flow-count"`
	MonitorSynRate             int    `json:"monitor-syn-rate"`
	MonitorFinRate             int    `json:"monitor-fin-rate"`
	MonitorRstRate             int    `json:"monitor-rst-rate"`
	MonitorTcpPktRate          int    `json:"monitor-tcp-pkt-rate"`
	MonitorUdpPktRate          int    `json:"monitor-udp-pkt-rate"`
	MonitorIcmpPktRate         int    `json:"monitor-icmp-pkt-rate"`
	Uuid                       string `json:"uuid"`
}

type DdosNetworkObjectIpList struct {
	SubnetIpAddr           string                                        `json:"subnet-ip-addr"`
	PrefixAnomalyThreshold DdosNetworkObjectIpListPrefixAnomalyThreshold `json:"prefix-anomaly-threshold"`
	Uuid                   string                                        `json:"uuid"`
	UserTag                string                                        `json:"user-tag"`
	SamplingEnable         []DdosNetworkObjectIpListSamplingEnable       `json:"sampling-enable"`
	SrcPortList            []DdosNetworkObjectIpListSrcPortList          `json:"src-port-list"`
}

type DdosNetworkObjectIpListPrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectIpListSrcPortList struct {
	PortNum                       int                                                             `json:"port-num"`
	Protocol                      string                                                          `json:"protocol"`
	HostSrcPortAnomalyThreshold   DdosNetworkObjectIpListSrcPortListHostSrcPortAnomalyThreshold   `json:"host-src-port-anomaly-threshold"`
	SubnetSrcPortAnomalyThreshold DdosNetworkObjectIpListSrcPortListSubnetSrcPortAnomalyThreshold `json:"subnet-src-port-anomaly-threshold"`
	Uuid                          string                                                          `json:"uuid"`
	UserTag                       string                                                          `json:"user-tag"`
}

type DdosNetworkObjectIpListSrcPortListHostSrcPortAnomalyThreshold struct {
	HostSrcPortPktRate int `json:"host-src-port-pkt-rate"`
	HostSrcPortBitRate int `json:"host-src-port-bit-rate"`
}

type DdosNetworkObjectIpListSrcPortListSubnetSrcPortAnomalyThreshold struct {
	SubnetSrcPortPktRate int `json:"subnet-src-port-pkt-rate"`
	SubnetSrcPortBitRate int `json:"subnet-src-port-bit-rate"`
}

type DdosNetworkObjectIpv6List struct {
	SubnetIpv6Addr         string                                          `json:"subnet-ipv6-addr"`
	PrefixAnomalyThreshold DdosNetworkObjectIpv6ListPrefixAnomalyThreshold `json:"prefix-anomaly-threshold"`
	Uuid                   string                                          `json:"uuid"`
	UserTag                string                                          `json:"user-tag"`
	SamplingEnable         []DdosNetworkObjectIpv6ListSamplingEnable       `json:"sampling-enable"`
	SrcPortList            []DdosNetworkObjectIpv6ListSrcPortList          `json:"src-port-list"`
}

type DdosNetworkObjectIpv6ListPrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpv6ListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectIpv6ListSrcPortList struct {
	PortNum                       int                                                               `json:"port-num"`
	Protocol                      string                                                            `json:"protocol"`
	HostSrcPortAnomalyThreshold   DdosNetworkObjectIpv6ListSrcPortListHostSrcPortAnomalyThreshold   `json:"host-src-port-anomaly-threshold"`
	SubnetSrcPortAnomalyThreshold DdosNetworkObjectIpv6ListSrcPortListSubnetSrcPortAnomalyThreshold `json:"subnet-src-port-anomaly-threshold"`
	Uuid                          string                                                            `json:"uuid"`
	UserTag                       string                                                            `json:"user-tag"`
}

type DdosNetworkObjectIpv6ListSrcPortListHostSrcPortAnomalyThreshold struct {
	HostSrcPortPktRate int `json:"host-src-port-pkt-rate"`
	HostSrcPortBitRate int `json:"host-src-port-bit-rate"`
}

type DdosNetworkObjectIpv6ListSrcPortListSubnetSrcPortAnomalyThreshold struct {
	SubnetSrcPortPktRate int `json:"subnet-src-port-pkt-rate"`
	SubnetSrcPortBitRate int `json:"subnet-src-port-bit-rate"`
}

type DdosNetworkObjectNetworkObjectAnomalyThreshold struct {
	NetworkObjectPktRate int `json:"network-object-pkt-rate"`
	NetworkObjectBitRate int `json:"network-object-bit-rate"`
}

type DdosNetworkObjectNotification318 struct {
	Configuration string                                         `json:"configuration"`
	Notification  []DdosNetworkObjectNotificationNotification319 `json:"notification"`
	Uuid          string                                         `json:"uuid"`
}

type DdosNetworkObjectNotificationNotification319 struct {
	NotificationTemplateName string `json:"notification-template-name"`
}

type DdosNetworkObjectRelativeAutoBreakDownThreshold struct {
	NetworkPercentage int `json:"network-percentage"`
	Permil            int `json:"permil"`
}

type DdosNetworkObjectSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectServiceBreakDownThresholdLocal struct {
	SvcPercentage int `json:"svc-percentage"`
}

type DdosNetworkObjectStaticAutoBreakDownThreshold struct {
	NetworkPktRate int `json:"network-pkt-rate"`
}

type DdosNetworkObjectSubNetwork320 struct {
	SubNetworkV4List []DdosNetworkObjectSubNetworkSubNetworkV4List `json:"sub-network-v4-list"`
	SubNetworkV6List []DdosNetworkObjectSubNetworkSubNetworkV6List `json:"sub-network-v6-list"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4List struct {
	SubnetIpAddr               string                                                                `json:"subnet-ip-addr"`
	HostAnomalyThreshold       DdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold       `json:"host-anomaly-threshold"`
	SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`
	SubnetBreakdown            int                                                                   `json:"subnet-breakdown"`
	BreakdownSubnetThreshold   DdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold   `json:"breakdown-subnet-threshold"`
	Uuid                       string                                                                `json:"uuid"`
	SamplingEnable             []DdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable           `json:"sampling-enable"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4ListHostAnomalyThreshold struct {
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

type DdosNetworkObjectSubNetworkSubNetworkV4ListSubNetworkAnomalyThreshold struct {
	StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
	StaticSubNetworkBitRate int `json:"static-sub-network-bit-rate"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4ListBreakdownSubnetThreshold struct {
	BreakdownSubnetPktRate int `json:"breakdown-subnet-pkt-rate"`
	BreakdownSubnetBitRate int `json:"breakdown-subnet-bit-rate"`
}

type DdosNetworkObjectSubNetworkSubNetworkV4ListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6List struct {
	SubnetIpv6Addr             string                                                                `json:"subnet-ipv6-addr"`
	HostAnomalyThreshold       DdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold       `json:"host-anomaly-threshold"`
	SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`
	SubnetBreakdown            int                                                                   `json:"subnet-breakdown"`
	Uuid                       string                                                                `json:"uuid"`
	SamplingEnable             []DdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable           `json:"sampling-enable"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6ListHostAnomalyThreshold struct {
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

type DdosNetworkObjectSubNetworkSubNetworkV6ListSubNetworkAnomalyThreshold struct {
	StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
	StaticSubNetworkBitRate int `json:"static-sub-network-bit-rate"`
}

type DdosNetworkObjectSubNetworkSubNetworkV6ListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectTopkDestinations321 struct {
	Uuid string `json:"uuid"`
}

type DdosNetworkObjectTrustlist322 struct {
	V4ClassList string `json:"v4-class-list"`
	V6ClassList string `json:"v6-class-list"`
	Uuid        string `json:"uuid"`
}

func (p *DdosNetworkObject) GetId() string {
	return url.QueryEscape(p.Inst.ObjectName)
}

func (p *DdosNetworkObject) getPath() string {
	return "ddos/network-object"
}

func (p *DdosNetworkObject) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObject::Post")
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

func (p *DdosNetworkObject) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObject::Get")
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
func (p *DdosNetworkObject) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObject::Put")
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

func (p *DdosNetworkObject) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObject::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
