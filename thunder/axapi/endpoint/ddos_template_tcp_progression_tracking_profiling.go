package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosTemplateTcpProgressionTrackingProfiling struct {
	Inst struct {
		ProfilingConnectionLifeModel int `json:"profiling-connection-life-model"`

		ProfilingRequestResponseModel int `json:"profiling-request-response-model"`

		ProfilingTimeWindowModel int `json:"profiling-time-window-model"`

		Uuid string `json:"uuid"`

		Tcp_name string
	} `json:"profiling"`
}

func (p *DdosTemplateTcpProgressionTrackingProfiling) GetId() string {
	return "1"
}

func (p *DdosTemplateTcpProgressionTrackingProfiling) getPath() string {
	return "ddos/template/tcp/" + p.Inst.Tcp_name + "/progression-tracking/profiling"
}

func (p *DdosTemplateTcpProgressionTrackingProfiling) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingProfiling::Post")
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

func (p *DdosTemplateTcpProgressionTrackingProfiling) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingProfiling::Get")
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
func (p *DdosTemplateTcpProgressionTrackingProfiling) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingProfiling::Put")
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

func (p *DdosTemplateTcpProgressionTrackingProfiling) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingProfiling::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
