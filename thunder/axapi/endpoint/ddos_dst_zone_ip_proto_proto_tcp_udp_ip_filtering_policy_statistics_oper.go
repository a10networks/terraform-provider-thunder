package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper struct {
	Oper DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOperOper `json:"oper"`

	Protocol string

	ZoneName string
}
type DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper struct {
	DtDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper `json:"ip-filtering-policy-statistics"`
}

type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOperOper struct {
	RuleList []DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOperOperRuleList `json:"rule-list"`
}

type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOperOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/ip-proto/proto-tcp-udp/" + p.Protocol + "/ip-filtering-policy-statistics/oper"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper, error) {
	logger.Println("DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyStatisticsOper
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
