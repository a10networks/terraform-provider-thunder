package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosMultiPuTrafficStatsOper struct {
	Oper DdosMultiPuTrafficStatsOperOper `json:"oper"`
}
type DataDdosMultiPuTrafficStatsOper struct {
	DtDdosMultiPuTrafficStatsOper DdosMultiPuTrafficStatsOper `json:"multi-pu-traffic-stats"`
}

type DdosMultiPuTrafficStatsOperOper struct {
	Pu1CpuRate        int                                     `json:"pu1-cpu-rate"`
	Pu1ThroughputRate int                                     `json:"pu1-throughput-rate"`
	Pu2CpuRate        int                                     `json:"pu2-cpu-rate"`
	Pu2ThroughputRate int                                     `json:"pu2-throughput-rate"`
	PuList            []DdosMultiPuTrafficStatsOperOperPuList `json:"pu-list"`
}

type DdosMultiPuTrafficStatsOperOperPuList struct {
	PuIndex int                                            `json:"pu-index"`
	KbitTop []DdosMultiPuTrafficStatsOperOperPuListKbitTop `json:"kbit-top"`
	PktTop  []DdosMultiPuTrafficStatsOperOperPuListPktTop  `json:"pkt-top"`
	CpuTop  []DdosMultiPuTrafficStatsOperOperPuListCpuTop  `json:"cpu-top"`
}

type DdosMultiPuTrafficStatsOperOperPuListKbitTop struct {
	ZoneName string `json:"zone-name"`
	Pu1Kbit  int    `json:"pu1-kbit"`
	Pu2Kbit  int    `json:"pu2-kbit"`
}

type DdosMultiPuTrafficStatsOperOperPuListPktTop struct {
	ZoneName string `json:"zone-name"`
	Pu1Pkt   int    `json:"pu1-pkt"`
	Pu2Pkt   int    `json:"pu2-pkt"`
}

type DdosMultiPuTrafficStatsOperOperPuListCpuTop struct {
	ZoneName string `json:"zone-name"`
	Pu1Cpu   int    `json:"pu1-cpu"`
	Pu2Cpu   int    `json:"pu2-cpu"`
}

func (p *DdosMultiPuTrafficStatsOper) GetId() string {
	return "1"
}

func (p *DdosMultiPuTrafficStatsOper) getPath() string {
	return "ddos/multi-pu-traffic-stats/oper"
}

func (p *DdosMultiPuTrafficStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosMultiPuTrafficStatsOper, error) {
	logger.Println("DdosMultiPuTrafficStatsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosMultiPuTrafficStatsOper
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
