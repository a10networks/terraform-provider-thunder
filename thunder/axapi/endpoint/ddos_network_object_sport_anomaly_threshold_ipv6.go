package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
	"strconv"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectSportAnomalyThresholdIpv6 struct {
	Inst struct {
		BitRate int `json:"bit-rate"`

		BitRatePercentage int `json:"bit-rate-percentage"`

		BitRatePercentageStr string `json:"bit-rate-percentage-str"`

		BitRateStr string `json:"bit-rate-str"`

		IpAddr string `json:"ip-addr"`

		IpSportBitRate int `json:"ip-sport-bit-rate"`

		IpSportBitRatePercentage int `json:"ip-sport-bit-rate-percentage"`

		IpSportBitRatePercentageStr string `json:"ip-sport-bit-rate-percentage-str"`

		IpSportBitRateStr string `json:"ip-sport-bit-rate-str"`

		IpSportPacketRate int `json:"ip-sport-packet-rate"`

		IpSportPacketRatePercentage int `json:"ip-sport-packet-rate-percentage"`

		IpSportPacketRatePercentageStr string `json:"ip-sport-packet-rate-percentage-str"`

		IpSportPacketRateStr string `json:"ip-sport-packet-rate-str"`

		PacketRate int `json:"packet-rate"`

		PacketRatePercentage int `json:"packet-rate-percentage"`

		PacketRatePercentageStr string `json:"packet-rate-percentage-str"`

		PacketRateStr string `json:"packet-rate-str"`

		Protocol string `json:"protocol"`

		SportNum int `json:"sport-num"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"ipv6"`
}

func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) GetId() string {
	return url.QueryEscape(p.Inst.IpAddr) + "+" + p.Inst.PacketRateStr + "+" + p.Inst.PacketRatePercentageStr + "+" + p.Inst.BitRateStr + "+" + p.Inst.BitRatePercentageStr + "+" + strconv.Itoa(p.Inst.SportNum) + "+" + p.Inst.Protocol + "+" + p.Inst.IpSportPacketRateStr + "+" + p.Inst.IpSportPacketRatePercentageStr + "+" + p.Inst.IpSportBitRateStr + "+" + p.Inst.IpSportBitRatePercentageStr
}

func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/sport-anomaly-threshold/ipv6"
}

func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdIpv6::Post")
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

func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdIpv6::Get")
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
func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdIpv6::Put")
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

func (p *DdosNetworkObjectSportAnomalyThresholdIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
