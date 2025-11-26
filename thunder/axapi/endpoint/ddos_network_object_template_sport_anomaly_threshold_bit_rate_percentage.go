package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage struct {
	Inst struct {
		Uuid string `json:"uuid"`

		Value int `json:"value"`

		Network_object_template_name string
	} `json:"bit-rate-percentage"`
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) getPath() string {
	return "ddos/network-object-template/" + p.Inst.Network_object_template_name + "/sport-anomaly-threshold/bit-rate-percentage"
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage::Post")
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

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage::Get")
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
func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage::Put")
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

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
