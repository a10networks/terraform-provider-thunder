package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SlbConnectionReuseStats struct {
	Stats SlbConnectionReuseStatsStats `json:"stats"`
}
type DataSlbConnectionReuseStats struct {
	DtSlbConnectionReuseStats SlbConnectionReuseStats `json:"connection-reuse"`
}

type SlbConnectionReuseStatsStats struct {
	Current_open                   int `json:"current_open"`
	Current_active                 int `json:"current_active"`
	Nbind                          int `json:"nbind"`
	Nunbind                        int `json:"nunbind"`
	Nestab                         int `json:"nestab"`
	Ntermi                         int `json:"ntermi"`
	Ntermi_err                     int `json:"ntermi_err"`
	Delay_unbind                   int `json:"delay_unbind"`
	Long_resp                      int `json:"long_resp"`
	Miss_resp                      int `json:"miss_resp"`
	Unbound_data_rcv               int `json:"unbound_data_rcv"`
	Pause_conn                     int `json:"pause_conn"`
	Pause_conn_fail                int `json:"pause_conn_fail"`
	Resume_conn                    int `json:"resume_conn"`
	Not_remove_from_rport          int `json:"not_remove_from_rport"`
	Zero_pconn_value               int `json:"zero_pconn_value"`
	Zero_pconn_bind_value          int `json:"zero_pconn_bind_value"`
	Current_http1_conn_in_the_pool int `json:"current_http1_conn_in_the_pool"`
	Current_http2_conn_in_the_pool int `json:"current_http2_conn_in_the_pool"`
}

func (p *SlbConnectionReuseStats) GetId() string {
	return "1"
}

func (p *SlbConnectionReuseStats) getPath() string {
	return "slb/connection-reuse/stats"
}

func (p *SlbConnectionReuseStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbConnectionReuseStats, error) {
	logger.Println("SlbConnectionReuseStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbConnectionReuseStats
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
