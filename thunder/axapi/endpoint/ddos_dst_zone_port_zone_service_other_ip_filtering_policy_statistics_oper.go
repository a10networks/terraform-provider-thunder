package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper `json:"oper"`

	Protocol string

	ZoneName string

	PortOther string
}
type DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port/zone-service-other/" + p.PortOther + "+" + p.Protocol + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper
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
