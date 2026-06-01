package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper `json:"oper"`

	ZoneName string

	PortNum string

	Protocol string
}
type DataDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port/zone-service/" + p.PortNum + "+" + p.Protocol + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper
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
