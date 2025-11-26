package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortRangeIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZonePortRangeIpFilteringPolicyStatisticsOperOper `json:"oper"`

	Protocol string

	ZoneName string

	PortRangeEnd string

	PortRangeStart string
}
type DataDdosDstZonePortRangeIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZonePortRangeIpFilteringPolicyStatisticsOper DdosDstZonePortRangeIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZonePortRangeIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZonePortRangeIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZonePortRangeIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZonePortRangeIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortRangeIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port-range/" + p.PortRangeStart + "+" + p.PortRangeEnd + "+" + p.Protocol + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZonePortRangeIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZonePortRangeIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortRangeIpFilteringPolicyStatisticsOper
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
