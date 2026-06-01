package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectIpSrcPort struct {
	Inst struct {
		HostSrcPortAnomalyThreshold DdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold `json:"host-src-port-anomaly-threshold"`

		PortNum int `json:"port-num"`

		Protocol string `json:"protocol"`

		SubnetSrcPortAnomalyThreshold DdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold `json:"subnet-src-port-anomaly-threshold"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		SubnetIpAddr string

		ObjectName string
	} `json:"src-port"`
}

type DdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold struct {
	HostSrcPortPktRate int `json:"host-src-port-pkt-rate"`
	HostSrcPortBitRate int `json:"host-src-port-bit-rate"`
}

type DdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold struct {
	SubnetSrcPortPktRate int `json:"subnet-src-port-pkt-rate"`
	SubnetSrcPortBitRate int `json:"subnet-src-port-bit-rate"`
}

func (p *DdosNetworkObjectIpSrcPort) GetId() string {
	return strconv.Itoa(p.Inst.PortNum) + "+" + p.Inst.Protocol
}

func (p *DdosNetworkObjectIpSrcPort) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/ip/" + p.Inst.SubnetIpAddr + "/src-port"
}

func (p *DdosNetworkObjectIpSrcPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpSrcPort::Post")
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

func (p *DdosNetworkObjectIpSrcPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpSrcPort::Get")
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
func (p *DdosNetworkObjectIpSrcPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpSrcPort::Put")
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

func (p *DdosNetworkObjectIpSrcPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpSrcPort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
