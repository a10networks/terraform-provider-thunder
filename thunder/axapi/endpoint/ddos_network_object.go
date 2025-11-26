package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosNetworkObject struct {
	Inst struct {
		AnomalyDetectionTrigger string `json:"anomaly-detection-trigger"`

		EnableTopK []DdosNetworkObjectEnableTopK `json:"enable-top-k"`

		FloodingMultiplier int `json:"flooding-multiplier"`

		HistogramMode string `json:"histogram-mode"`

		HostAnomalyThreshold DdosNetworkObjectHostAnomalyThreshold `json:"host-anomaly-threshold"`

		HostSportDiscovery string `json:"host-sport-discovery"`

		IndicatorsToMonitor DdosNetworkObjectIndicatorsToMonitor312 `json:"indicators-to-monitor"`

		IpList []DdosNetworkObjectIpList `json:"ip-list"`

		Ipv6List []DdosNetworkObjectIpv6List `json:"ipv6-list"`

		NetworkObjectAnomalyThreshold DdosNetworkObjectNetworkObjectAnomalyThreshold `json:"network-object-anomaly-threshold"`

		NetworkObjectTemplate string `json:"network-object-template"`

		Notification DdosNetworkObjectNotification313 `json:"notification"`

		ObjectName string `json:"object-name"`

		OperationalMode string `json:"operational-mode"`

		RelativeAutoBreakDownThreshold DdosNetworkObjectRelativeAutoBreakDownThreshold `json:"relative-auto-break-down-threshold"`

		SamplingEnable []DdosNetworkObjectSamplingEnable `json:"sampling-enable"`

		ServiceBreakDownThresholdLocal DdosNetworkObjectServiceBreakDownThresholdLocal `json:"service-break-down-threshold-local"`

		ServiceDiscovery string `json:"service-discovery"`

		SportAnomalyDetection string `json:"sport-anomaly-detection"`

		SportAnomalyThreshold DdosNetworkObjectSportAnomalyThreshold315 `json:"sport-anomaly-threshold"`

		SportDiscoveryThreshold DdosNetworkObjectSportDiscoveryThreshold `json:"sport-discovery-threshold"`

		SportList []DdosNetworkObjectSportList `json:"sport-list"`

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

type DdosNetworkObjectIndicatorsToMonitor312 struct {
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
}

type DdosNetworkObjectIpListPrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectIpv6List struct {
	SubnetIpv6Addr         string                                          `json:"subnet-ipv6-addr"`
	PrefixAnomalyThreshold DdosNetworkObjectIpv6ListPrefixAnomalyThreshold `json:"prefix-anomaly-threshold"`
	Uuid                   string                                          `json:"uuid"`
	UserTag                string                                          `json:"user-tag"`
	SamplingEnable         []DdosNetworkObjectIpv6ListSamplingEnable       `json:"sampling-enable"`
}

type DdosNetworkObjectIpv6ListPrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpv6ListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectNetworkObjectAnomalyThreshold struct {
	NetworkObjectPktRate int `json:"network-object-pkt-rate"`
	NetworkObjectBitRate int `json:"network-object-bit-rate"`
}

type DdosNetworkObjectNotification313 struct {
	Configuration string                                         `json:"configuration"`
	Notification  []DdosNetworkObjectNotificationNotification314 `json:"notification"`
	Uuid          string                                         `json:"uuid"`
}

type DdosNetworkObjectNotificationNotification314 struct {
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

type DdosNetworkObjectSportAnomalyThreshold315 struct {
	PacketRate           DdosNetworkObjectSportAnomalyThresholdPacketRate316           `json:"packet-rate"`
	PacketRatePercentage DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317 `json:"packet-rate-percentage"`
	BitRate              DdosNetworkObjectSportAnomalyThresholdBitRate318              `json:"bit-rate"`
	BitRatePercentage    DdosNetworkObjectSportAnomalyThresholdBitRatePercentage319    `json:"bit-rate-percentage"`
	IpList               []DdosNetworkObjectSportAnomalyThresholdIpList                `json:"ip-list"`
	Ipv6List             []DdosNetworkObjectSportAnomalyThresholdIpv6List              `json:"ipv6-list"`
	SportList            []DdosNetworkObjectSportAnomalyThresholdSportList             `json:"sport-list"`
}

type DdosNetworkObjectSportAnomalyThresholdPacketRate316 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage317 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdBitRate318 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdBitRatePercentage319 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdIpList struct {
	IpAddr                         string `json:"ip-addr"`
	PacketRateStr                  string `json:"packet-rate-str"`
	PacketRatePercentageStr        string `json:"packet-rate-percentage-str"`
	BitRateStr                     string `json:"bit-rate-str"`
	BitRatePercentageStr           string `json:"bit-rate-percentage-str"`
	PacketRate                     int    `json:"packet-rate"`
	PacketRatePercentage           int    `json:"packet-rate-percentage"`
	BitRate                        int    `json:"bit-rate"`
	BitRatePercentage              int    `json:"bit-rate-percentage"`
	SportNum                       int    `json:"sport-num"`
	Protocol                       string `json:"protocol"`
	IpSportPacketRateStr           string `json:"ip-sport-packet-rate-str"`
	IpSportPacketRatePercentageStr string `json:"ip-sport-packet-rate-percentage-str"`
	IpSportBitRateStr              string `json:"ip-sport-bit-rate-str"`
	IpSportBitRatePercentageStr    string `json:"ip-sport-bit-rate-percentage-str"`
	IpSportPacketRate              int    `json:"ip-sport-packet-rate"`
	IpSportPacketRatePercentage    int    `json:"ip-sport-packet-rate-percentage"`
	IpSportBitRate                 int    `json:"ip-sport-bit-rate"`
	IpSportBitRatePercentage       int    `json:"ip-sport-bit-rate-percentage"`
	Uuid                           string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdIpv6List struct {
	IpAddr                         string `json:"ip-addr"`
	PacketRateStr                  string `json:"packet-rate-str"`
	PacketRatePercentageStr        string `json:"packet-rate-percentage-str"`
	BitRateStr                     string `json:"bit-rate-str"`
	BitRatePercentageStr           string `json:"bit-rate-percentage-str"`
	PacketRate                     int    `json:"packet-rate"`
	PacketRatePercentage           int    `json:"packet-rate-percentage"`
	BitRate                        int    `json:"bit-rate"`
	BitRatePercentage              int    `json:"bit-rate-percentage"`
	SportNum                       int    `json:"sport-num"`
	Protocol                       string `json:"protocol"`
	IpSportPacketRateStr           string `json:"ip-sport-packet-rate-str"`
	IpSportPacketRatePercentageStr string `json:"ip-sport-packet-rate-percentage-str"`
	IpSportBitRateStr              string `json:"ip-sport-bit-rate-str"`
	IpSportBitRatePercentageStr    string `json:"ip-sport-bit-rate-percentage-str"`
	IpSportPacketRate              int    `json:"ip-sport-packet-rate"`
	IpSportPacketRatePercentage    int    `json:"ip-sport-packet-rate-percentage"`
	IpSportBitRate                 int    `json:"ip-sport-bit-rate"`
	IpSportBitRatePercentage       int    `json:"ip-sport-bit-rate-percentage"`
	Uuid                           string `json:"uuid"`
}

type DdosNetworkObjectSportAnomalyThresholdSportList struct {
	SportNum                int    `json:"sport-num"`
	Protocol                string `json:"protocol"`
	PacketRateStr           string `json:"packet-rate-str"`
	PacketRatePercentageStr string `json:"packet-rate-percentage-str"`
	BitRateStr              string `json:"bit-rate-str"`
	BitRatePercentageStr    string `json:"bit-rate-percentage-str"`
	PacketRate              int    `json:"packet-rate"`
	PacketRatePercentage    int    `json:"packet-rate-percentage"`
	BitRate                 int    `json:"bit-rate"`
	BitRatePercentage       int    `json:"bit-rate-percentage"`
	Uuid                    string `json:"uuid"`
}

type DdosNetworkObjectSportDiscoveryThreshold struct {
	SportHeavyHitterPercentage      int `json:"sport-heavy-hitter-percentage" dval:"10"`
	SportDiscoveryBitRatePercentage int `json:"sport-discovery-bit-rate-percentage"`
}

type DdosNetworkObjectSportList struct {
	PortNum  int    `json:"port-num"`
	Protocol string `json:"protocol"`
	Uuid     string `json:"uuid"`
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
