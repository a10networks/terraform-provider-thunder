package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type FlowspecSourcePort struct {
	Inst struct {
		PortAttribute string `json:"port-attribute"`

		PortNum int `json:"port-num"`

		PortNumEnd int `json:"port-num-end"`

		Uuid string `json:"uuid"`

		Flowspec_name string
	} `json:"source-port"`
}

func (p *FlowspecSourcePort) GetId() string {
	return p.Inst.PortAttribute + "+" + strconv.Itoa(p.Inst.PortNum) + "+" + strconv.Itoa(p.Inst.PortNumEnd)
}

func (p *FlowspecSourcePort) getPath() string {
	return "flowspec/" + p.Inst.Flowspec_name + "/source-port"
}

func (p *FlowspecSourcePort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FlowspecSourcePort::Post")
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

func (p *FlowspecSourcePort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FlowspecSourcePort::Get")
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
func (p *FlowspecSourcePort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FlowspecSourcePort::Put")
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

func (p *FlowspecSourcePort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FlowspecSourcePort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
