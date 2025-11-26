package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ControllerProbeInfoOper struct {
	Oper ControllerProbeInfoOperOper `json:"oper"`
}
type DataControllerProbeInfoOper struct {
	DtControllerProbeInfoOper ControllerProbeInfoOper `json:"probe-info"`
}

type ControllerProbeInfoOperOper struct {
	ProbeStatus               string                                                 `json:"probe-status"`
	ControllerStatus          string                                                 `json:"controller-status"`
	DataShowtechExportLogList []ControllerProbeInfoOperOperDataShowtechExportLogList `json:"DATA-SHOWTECH-export-log-list"`
	DataVarlogExportLogList   []ControllerProbeInfoOperOperDataVarlogExportLogList   `json:"DATA-VARLOG-export-log-list"`
}

type ControllerProbeInfoOperOperDataShowtechExportLogList struct {
	ShowtechFilename  string `json:"showtech-filename"`
	ShowtechSize      int    `json:"showtech-size"`
	ShowtechTimestamp string `json:"showtech-timestamp"`
}

type ControllerProbeInfoOperOperDataVarlogExportLogList struct {
	VarlogFilename  string `json:"varlog-filename"`
	VarlogSize      int    `json:"varlog-size"`
	VarlogTimestamp string `json:"varlog-timestamp"`
}

func (p *ControllerProbeInfoOper) GetId() string {
	return "1"
}

func (p *ControllerProbeInfoOper) getPath() string {
	return "controller/probe-info/oper"
}

func (p *ControllerProbeInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataControllerProbeInfoOper, error) {
	logger.Println("ControllerProbeInfoOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataControllerProbeInfoOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
