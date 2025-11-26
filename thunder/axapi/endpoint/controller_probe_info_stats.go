package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ControllerProbeInfoStats struct {
	Stats ControllerProbeInfoStatsStats `json:"stats"`
}
type DataControllerProbeInfoStats struct {
	DtControllerProbeInfoStats ControllerProbeInfoStats `json:"probe-info"`
}

type ControllerProbeInfoStatsStats struct {
	DataShowtechSent    int `json:"data-showtech-sent"`
	DataShowtechFailed  int `json:"data-showtech-failed"`
	DataVarlogSent      int `json:"data-varlog-sent"`
	DataVarlogFailed    int `json:"data-varlog-failed"`
	SshConnectionFailed int `json:"ssh-connection-failed"`
}

func (p *ControllerProbeInfoStats) GetId() string {
	return "1"
}

func (p *ControllerProbeInfoStats) getPath() string {
	return "controller/probe-info/stats"
}

func (p *ControllerProbeInfoStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataControllerProbeInfoStats, error) {
	logger.Println("ControllerProbeInfoStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataControllerProbeInfoStats
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
