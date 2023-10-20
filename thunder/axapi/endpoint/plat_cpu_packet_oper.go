package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type PlatCpuPacketOper struct {
	Oper PlatCpuPacketOperOper `json:"oper"`
}

type PlatCpuPackett struct {
	DataPlatCpuPacket PlatCpuPacketOper `json:"plat-cpu-packet"`
}

type PlatCpuPacketOperOper struct {
	PktStats []PlatCpuPacketOperOperPktStats `json:"pkt-stats"`
}

type PlatCpuPacketOperOperPktStats struct {
	CpuNum  int `json:"cpu-num"`
	PktSent int `json:"pkt-sent"`
	PktRcvd int `json:"pkt-rcvd"`
	PktDrop int `json:"pkt-drop"`
}

func (p *PlatCpuPacketOper) GetId() string {
	return "1"
}

func (p *PlatCpuPacketOper) getPath() string {
	return "plat-cpu-packet/oper"
}

func (p *PlatCpuPacketOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (PlatCpuPackett, error) {
	logger.Println("PlatCpuPacketOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var pcpu PlatCpuPackett
	if err == nil {
		err = json.Unmarshal(axResp, &pcpu)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return pcpu, err
}
