package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectIpv6SrcPort struct {
	Inst struct {
		HostSrcPortAnomalyThreshold DdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold `json:"host-src-port-anomaly-threshold"`

		PortNum int `json:"port-num"`

		Protocol string `json:"protocol"`

		SubnetSrcPortAnomalyThreshold DdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold `json:"subnet-src-port-anomaly-threshold"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		SubnetIpv6Addr string

		ObjectName string
	} `json:"src-port"`
}

type DdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold struct {
	HostSrcPortPktRate int `json:"host-src-port-pkt-rate"`
	HostSrcPortBitRate int `json:"host-src-port-bit-rate"`
}

type DdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold struct {
	SubnetSrcPortPktRate int `json:"subnet-src-port-pkt-rate"`
	SubnetSrcPortBitRate int `json:"subnet-src-port-bit-rate"`
}

func (p *DdosNetworkObjectIpv6SrcPort) GetId() string {
	return strconv.Itoa(p.Inst.PortNum) + "+" + p.Inst.Protocol
}

func (p *DdosNetworkObjectIpv6SrcPort) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/ipv6/" + p.Inst.SubnetIpv6Addr + "/src-port"
}

func (p *DdosNetworkObjectIpv6SrcPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6SrcPort::Post")
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

func (p *DdosNetworkObjectIpv6SrcPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6SrcPort::Get")
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
func (p *DdosNetworkObjectIpv6SrcPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6SrcPort::Put")
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

func (p *DdosNetworkObjectIpv6SrcPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6SrcPort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
