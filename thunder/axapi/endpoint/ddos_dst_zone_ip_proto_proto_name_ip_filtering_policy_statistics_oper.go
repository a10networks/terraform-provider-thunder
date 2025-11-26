package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper `json:"oper"`

	Protocol string

	ZoneName string
}
type DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/ip-proto/proto-name/" + p.Protocol + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper
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
