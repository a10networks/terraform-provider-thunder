package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ControllerTunnelStatsOper struct {
	Oper ControllerTunnelStatsOperOper `json:"oper"`
}
type DataControllerTunnelStatsOper struct {
	DtControllerTunnelStatsOper ControllerTunnelStatsOper `json:"tunnel-stats"`
}

type ControllerTunnelStatsOperOper struct {
	Status         string `json:"status"`
	BytesSent      int    `json:"bytes-sent"`
	BytesRecieved  int    `json:"bytes-recieved"`
	NumberOfErrors int    `json:"number-of-errors"`
	Uptime         string `json:"uptime"`
	ErrorMessage   string `json:"error-message"`
}

func (p *ControllerTunnelStatsOper) GetId() string {
	return "1"
}

func (p *ControllerTunnelStatsOper) getPath() string {
	return "controller/tunnel-stats/oper"
}

func (p *ControllerTunnelStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataControllerTunnelStatsOper, error) {
	logger.Println("ControllerTunnelStatsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataControllerTunnelStatsOper
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
