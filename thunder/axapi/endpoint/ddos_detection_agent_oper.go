package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDetectionAgentOper struct {
	AgentName string `json:"agent-name"`

	Oper DdosDetectionAgentOperOper `json:"oper"`
}
type DataDdosDetectionAgentOper struct {
	DtDdosDetectionAgentOper DdosDetectionAgentOper `json:"agent"`
}

type DdosDetectionAgentOperOper struct {
	Brand       string                                  `json:"brand"`
	SamplerList []DdosDetectionAgentOperOperSamplerList `json:"sampler-list"`
}

type DdosDetectionAgentOperOperSamplerList struct {
	SamplerId         int `json:"sampler-id"`
	SampleMode        int `json:"sample-mode"`
	SamplingAlgorithm int `json:"sampling-algorithm"`
	SamplingRate      int `json:"sampling-rate"`
	ActiveTimeout     int `json:"active-timeout"`
	InactiveTimeout   int `json:"inactive-timeout"`
}

func (p *DdosDetectionAgentOper) GetId() string {
	return "1"
}

func (p *DdosDetectionAgentOper) getPath() string {
	return "ddos/detection/agent/" + p.AgentName + "/oper"
}

func (p *DdosDetectionAgentOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDetectionAgentOper, error) {
	logger.Println("DdosDetectionAgentOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDetectionAgentOper
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
