package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectIpv6 struct {
	Inst struct {
		PrefixAnomalyThreshold DdosNetworkObjectIpv6PrefixAnomalyThreshold `json:"prefix-anomaly-threshold"`

		SamplingEnable []DdosNetworkObjectIpv6SamplingEnable `json:"sampling-enable"`

		SubnetIpv6Addr string `json:"subnet-ipv6-addr"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"ipv6"`
}

type DdosNetworkObjectIpv6PrefixAnomalyThreshold struct {
	PrefixPktRate int `json:"prefix-pkt-rate"`
	PrefixBitRate int `json:"prefix-bit-rate"`
}

type DdosNetworkObjectIpv6SamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *DdosNetworkObjectIpv6) GetId() string {
	return url.QueryEscape(p.Inst.SubnetIpv6Addr)
}

func (p *DdosNetworkObjectIpv6) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/ipv6"
}

func (p *DdosNetworkObjectIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6::Post")
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

func (p *DdosNetworkObjectIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6::Get")
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
func (p *DdosNetworkObjectIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6::Put")
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

func (p *DdosNetworkObjectIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
