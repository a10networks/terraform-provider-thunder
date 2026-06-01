package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate struct {
	Inst struct {
		Uuid string `json:"uuid"`

		Value int `json:"value"`

		Network_object_template_name string
	} `json:"packet-rate"`
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) getPath() string {
	return "ddos/network-object-template/" + p.Inst.Network_object_template_name + "/sport-anomaly-threshold/packet-rate"
}

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate::Post")
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

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate::Get")
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
func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate::Put")
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

func (p *DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
