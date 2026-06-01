package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectTemplate struct {
	Inst struct {
		AnomalyDetectionTrigger string `json:"anomaly-detection-trigger" dval:"all"`

		FloodingMultiplier int `json:"flooding-multiplier" dval:"2"`

		HistogramMode string `json:"histogram-mode" dval:"observe"`

		HostAnomalyThreshold DdosNetworkObjectTemplateHostAnomalyThreshold `json:"host-anomaly-threshold"`

		IndicatorsToMonitor DdosNetworkObjectTemplateIndicatorsToMonitor311 `json:"indicators-to-monitor"`

		Name string `json:"name"`

		NetworkObjectAnomalyThreshold DdosNetworkObjectTemplateNetworkObjectAnomalyThreshold `json:"network-object-anomaly-threshold"`

		OperationalMode string `json:"operational-mode" dval:"learning"`

		ServiceBreakDownThresholdLocal DdosNetworkObjectTemplateServiceBreakDownThresholdLocal `json:"service-break-down-threshold-local"`

		ServiceDiscovery string `json:"service-discovery"`

		SportAnomalyThreshold DdosNetworkObjectTemplateSportAnomalyThreshold312 `json:"sport-anomaly-threshold"`

		ThresholdSensitivity string `json:"threshold-sensitivity" dval:"OFF"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"network-object-template"`
}

type DdosNetworkObjectTemplateHostAnomalyThreshold struct {
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

type DdosNetworkObjectTemplateIndicatorsToMonitor311 struct {
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

type DdosNetworkObjectTemplateNetworkObjectAnomalyThreshold struct {
	NetworkObjectPktRate int `json:"network-object-pkt-rate"`
	NetworkObjectBitRate int `json:"network-object-bit-rate"`
}

type DdosNetworkObjectTemplateServiceBreakDownThresholdLocal struct {
	SvcPercentage int `json:"svc-percentage"`
}

type DdosNetworkObjectTemplateSportAnomalyThreshold312 struct {
	PacketRate           DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate313           `json:"packet-rate"`
	PacketRatePercentage DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage314 `json:"packet-rate-percentage"`
	BitRate              DdosNetworkObjectTemplateSportAnomalyThresholdBitRate315              `json:"bit-rate"`
	BitRatePercentage    DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage316    `json:"bit-rate-percentage"`
}

type DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate313 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage314 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectTemplateSportAnomalyThresholdBitRate315 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

type DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage316 struct {
	Value int    `json:"value"`
	Uuid  string `json:"uuid"`
}

func (p *DdosNetworkObjectTemplate) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *DdosNetworkObjectTemplate) getPath() string {
	return "ddos/network-object-template"
}

func (p *DdosNetworkObjectTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplate::Post")
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

func (p *DdosNetworkObjectTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplate::Get")
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
func (p *DdosNetworkObjectTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplate::Put")
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

func (p *DdosNetworkObjectTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
