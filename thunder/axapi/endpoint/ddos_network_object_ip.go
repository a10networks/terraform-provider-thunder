package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectIp struct {
	Inst struct {
		PrefixAnomalyThreshold DdosNetworkObjectIpPrefixAnomalyThreshold `json:"prefix-anomaly-threshold"`

		SamplingEnable []DdosNetworkObjectIpSamplingEnable `json:"sampling-enable"`

		SrcPortList []DdosNetworkObjectIpSrcPortList `json:"src-port-list"`

		SubnetIpAddr string `json:"subnet-ip-addr"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"ip"`
}

type DdosNetworkObjectIpPrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosNetworkObjectIpSrcPortList struct {
	PortNum                       int                                                         `json:"port-num"`
	Protocol                      string                                                      `json:"protocol"`
	HostSrcPortAnomalyThreshold   DdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold   `json:"host-src-port-anomaly-threshold"`
	SubnetSrcPortAnomalyThreshold DdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold `json:"subnet-src-port-anomaly-threshold"`
	Uuid                          string                                                      `json:"uuid"`
	UserTag                       string                                                      `json:"user-tag"`
}

type DdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold struct {
	HostSrcPortPktRate int `json:"host-src-port-pkt-rate"`
	HostSrcPortBitRate int `json:"host-src-port-bit-rate"`
}

type DdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold struct {
	SubnetSrcPortPktRate int `json:"subnet-src-port-pkt-rate"`
	SubnetSrcPortBitRate int `json:"subnet-src-port-bit-rate"`
}

func (p *DdosNetworkObjectIp) GetId() string {
	return url.QueryEscape(p.Inst.SubnetIpAddr)
}

func (p *DdosNetworkObjectIp) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/ip"
}

func (p *DdosNetworkObjectIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIp::Post")
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

func (p *DdosNetworkObjectIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIp::Get")
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
func (p *DdosNetworkObjectIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIp::Put")
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

func (p *DdosNetworkObjectIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
