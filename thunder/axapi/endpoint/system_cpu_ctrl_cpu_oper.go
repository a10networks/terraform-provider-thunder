package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type SystemCpuCtrlCpuOper struct {
	Inst struct {
		Oper SystemCpuCtrlCpuOperOper `json:"oper"`
	} `json:"oper"`
}

type SystemCpuCtrlCpuOperOper struct {
	CurrentTime string                             `json:"current-time"`
	NumberOfCpu int                                `json:"number-of-cpu"`
	CpuUsage    []SystemCpuCtrlCpuOperOperCpuUsage `json:"cpu-usage"`
}

type SystemCpuCtrlCpuOperOperCpuUsage struct {
	CpuId int `json:"cpu-id"`
	Sec1  int `json:"1-sec`
	Sec5  int `json:"5-sec"`
	Sec10 int `json:"10-sec"`
	Sec30 int `json:"30-sec"`
	Sec60 int `json:"60-sec"`
}

func (p *SystemCpuCtrlCpuOper) GetId() string {
	return "1"
}

func (p *SystemCpuCtrlCpuOper) getPath() string {
	return "system-cpu/ctrl-cpu/oper"
}

func (p *SystemCpuCtrlCpuOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuCtrlCpuOper::Post")
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

func (p *SystemCpuCtrlCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuCtrlCpuOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SystemCpuCtrlCpuOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuCtrlCpuOper::Put")
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

func (p *SystemCpuCtrlCpuOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuCtrlCpuOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
