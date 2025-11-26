package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper `json:"oper"`

	ZoneName string

	ProtocolNum string
}
type DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/ip-proto/proto-number/" + p.ProtocolNum + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper
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
