package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DebugController struct {
	Inst struct {
		Anomaly int `json:"anomaly"`

		AppSvcId string `json:"app-svc-id"`

		Error int `json:"error"`

		LogdAuditExport int `json:"logd-audit-export"`

		LogdSyslogExport int `json:"logd-syslog-export"`

		Metrics int `json:"metrics"`

		ObjectUuid string `json:"object-uuid"`

		PerConnection int `json:"per-connection"`

		PerRequest int `json:"per-request"`

		Registration int `json:"registration"`

		Uri string `json:"uri"`

		Uuid string `json:"uuid"`
	} `json:"controller"`
}

func (p *DebugController) GetId() string {
	return "1"
}

func (p *DebugController) getPath() string {
	return "debug/controller"
}

func (p *DebugController) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugController::Post")
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

func (p *DebugController) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugController::Get")
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
func (p *DebugController) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugController::Put")
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

func (p *DebugController) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugController::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
