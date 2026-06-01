package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type HealthSourceNatStats struct {
	Stats HealthSourceNatStatsStats `json:"stats"`
}
type DataHealthSourceNatStats struct {
	DtHealthSourceNatStats HealthSourceNatStats `json:"source-nat"`
}

type HealthSourceNatStatsStats struct {
	Act_recv_from_sby           int `json:"act_recv_from_sby"`
	Act_send_to_sby             int `json:"act_send_to_sby"`
	Sby_recv_from_act           int `json:"sby_recv_from_act"`
	Sby_send_to_act             int `json:"sby_send_to_act"`
	Sby_recv_from_act_err       int `json:"sby_recv_from_act_err"`
	Recv_from_kernel            int `json:"recv_from_kernel"`
	Send_to_kernel              int `json:"send_to_kernel"`
	Send_to_kernel_err          int `json:"send_to_kernel_err"`
	Sby_no_peer                 int `json:"sby_no_peer"`
	Dcmsg_err                   int `json:"dcmsg_err"`
	No_slb_object               int `json:"no_slb_object"`
	Smart_nat_init_port_err     int `json:"smart_nat_init_port_err"`
	Smart_nat_init_inst_err     int `json:"smart_nat_init_inst_err"`
	Smart_nat_rserver_route_err int `json:"smart_nat_rserver_route_err"`
	Smart_nat_rserver_ip_err    int `json:"smart_nat_rserver_ip_err"`
	Nat_resource_err            int `json:"nat_resource_err"`
	Frag_err                    int `json:"frag_err"`
}

func (p *HealthSourceNatStats) GetId() string {
	return "1"
}

func (p *HealthSourceNatStats) getPath() string {
	return "health/source-nat/stats"
}

func (p *HealthSourceNatStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataHealthSourceNatStats, error) {
	logger.Println("HealthSourceNatStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataHealthSourceNatStats
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
