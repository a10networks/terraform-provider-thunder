package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceManagement2Stats struct {
	Stats InterfaceManagement2StatsStats `json:"stats"`
}
type DataInterfaceManagement2Stats struct {
	DtInterfaceManagement2Stats InterfaceManagement2Stats `json:"management2"`
}

type InterfaceManagement2StatsStats struct {
	Packets_input          int `json:"packets_input"`
	Bytes_input            int `json:"bytes_input"`
	Received_broadcasts    int `json:"received_broadcasts"`
	Received_multicasts    int `json:"received_multicasts"`
	Received_unicasts      int `json:"received_unicasts"`
	Input_errors           int `json:"input_errors"`
	Crc                    int `json:"crc"`
	Frame                  int `json:"frame"`
	Input_err_short        int `json:"input_err_short"`
	Input_err_long         int `json:"input_err_long"`
	Packets_output         int `json:"packets_output"`
	Bytes_output           int `json:"bytes_output"`
	Transmitted_broadcasts int `json:"transmitted_broadcasts"`
	Transmitted_multicasts int `json:"transmitted_multicasts"`
	Transmitted_unicasts   int `json:"transmitted_unicasts"`
	Output_errors          int `json:"output_errors"`
	Collisions             int `json:"collisions"`
}

func (p *InterfaceManagement2Stats) GetId() string {
	return "1"
}

func (p *InterfaceManagement2Stats) getPath() string {
	return "interface/management2/stats"
}

func (p *InterfaceManagement2Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceManagement2Stats, error) {
	logger.Println("InterfaceManagement2Stats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataInterfaceManagement2Stats
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
