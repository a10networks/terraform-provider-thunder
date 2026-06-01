package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectIndicatorsToMonitor struct {
	Inst struct {
		Enable int `json:"enable"`

		MonitorBitRate int `json:"monitor-bit-rate"`

		MonitorFinRate int `json:"monitor-fin-rate"`

		MonitorFlowCount int `json:"monitor-flow-count"`

		MonitorIcmpPktRate int `json:"monitor-icmp-pkt-rate"`

		MonitorPktRate int `json:"monitor-pkt-rate"`

		MonitorRevBitRate int `json:"monitor-rev-bit-rate"`

		MonitorRevPktRate int `json:"monitor-rev-pkt-rate"`

		MonitorRstRate int `json:"monitor-rst-rate"`

		MonitorSynRate int `json:"monitor-syn-rate"`

		MonitorTcpPktRate int `json:"monitor-tcp-pkt-rate"`

		MonitorUdpPktRate int `json:"monitor-udp-pkt-rate"`

		MonitorUndiscoveredPktRate int `json:"monitor-undiscovered-pkt-rate"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"indicators-to-monitor"`
}

func (p *DdosNetworkObjectIndicatorsToMonitor) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectIndicatorsToMonitor) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/indicators-to-monitor"
}

func (p *DdosNetworkObjectIndicatorsToMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIndicatorsToMonitor::Post")
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

func (p *DdosNetworkObjectIndicatorsToMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIndicatorsToMonitor::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosNetworkObjectIndicatorsToMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIndicatorsToMonitor::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *DdosNetworkObjectIndicatorsToMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIndicatorsToMonitor::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
