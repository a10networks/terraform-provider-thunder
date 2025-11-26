package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectSportAnomalyThresholdSport struct {
	Inst struct {
		BitRate int `json:"bit-rate"`

		BitRatePercentage int `json:"bit-rate-percentage"`

		BitRatePercentageStr string `json:"bit-rate-percentage-str"`

		BitRateStr string `json:"bit-rate-str"`

		PacketRate int `json:"packet-rate"`

		PacketRatePercentage int `json:"packet-rate-percentage"`

		PacketRatePercentageStr string `json:"packet-rate-percentage-str"`

		PacketRateStr string `json:"packet-rate-str"`

		Protocol string `json:"protocol"`

		SportNum int `json:"sport-num"`

		Uuid string `json:"uuid"`

		ObjectName string
	} `json:"sport"`
}

func (p *DdosNetworkObjectSportAnomalyThresholdSport) GetId() string {
	return strconv.Itoa(p.Inst.SportNum) + "+" + p.Inst.Protocol + "+" + p.Inst.PacketRateStr + "+" + p.Inst.PacketRatePercentageStr + "+" + p.Inst.BitRateStr + "+" + p.Inst.BitRatePercentageStr
}

func (p *DdosNetworkObjectSportAnomalyThresholdSport) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/sport-anomaly-threshold/sport"
}

func (p *DdosNetworkObjectSportAnomalyThresholdSport) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdSport::Post")
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

func (p *DdosNetworkObjectSportAnomalyThresholdSport) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdSport::Get")
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
func (p *DdosNetworkObjectSportAnomalyThresholdSport) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdSport::Put")
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

func (p *DdosNetworkObjectSportAnomalyThresholdSport) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectSportAnomalyThresholdSport::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
