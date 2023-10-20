package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type SystemCpuCtrlCpuOper struct {
	Oper SystemCpuCtrlCpuOperOper `json:"oper"`
}

type SystemCpuCtrlCpus struct {
	DataSystemCpuCtrlCpu SystemCpuCtrlCpuOper `json:"ctrl-cpu"`
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

func (p *SystemCpuCtrlCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (SystemCpuCtrlCpus, error) {
	logger.Println("SystemCpuCtrlCpuOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)

	var systemCpuCtrl SystemCpuCtrlCpus
	if err == nil {
		err = json.Unmarshal(axResp, &systemCpuCtrl)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return systemCpuCtrl, err
}
