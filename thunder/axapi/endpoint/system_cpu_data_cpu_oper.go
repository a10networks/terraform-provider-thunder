package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type SystemCpuDataCpuOper struct {
	Inst struct {
		Oper SystemCpuDataCpuOperOper `json:"oper"`
	} `json:"oper"`
}

type SystemCpuDataCpuOperOper struct {
	NumberOfCpu     int                                `json:"number-of-cpu"`
	NumberOfDataCpu int                                `json:"number-of-data-cpu"`
	CpuUsage        []SystemCpuDataCpuOperOperCpuUsage `json:"cpu-usage"`
}

type SystemCpuDataCpuOperOperCpuUsage struct {
	CpuId   int    `json:"cpu-id"`
	Sec1    int    `json:"1-sec"`
	Sec5    int    `json:"5-sec"`
	Sec10   int    `json:"10-sec"`
	Sec30   int    `json:"30-sec"`
	Sec60   int    `json:"60-sec"`
	DcpuStr string `json:"dcpu-str"`
}

func (p *SystemCpuDataCpuOper) GetId() string {
	return "1"
}

func (p *SystemCpuDataCpuOper) getPath() string {
	return "system-cpu/data-cpu/oper"
}

func (p *SystemCpuDataCpuOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuDataCpuOper::Post")
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

func (p *SystemCpuDataCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuDataCpuOper::Get")
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

func (p *SystemCpuDataCpuOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuDataCpuOper::Put")
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

func (p *SystemCpuDataCpuOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemCpuDataCpuOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
