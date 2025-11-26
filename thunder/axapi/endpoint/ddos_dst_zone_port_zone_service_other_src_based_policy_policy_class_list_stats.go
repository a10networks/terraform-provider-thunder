package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats struct {
	ClassListName string `json:"class-list-name"`

	Stats DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats `json:"stats"`

	PortOther string

	Protocol string

	SrcBasedPolicyName string

	ZoneName string
}
type DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats struct {
	DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats `json:"policy-class-list"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats struct {
	Packet_received                 int `json:"packet_received"`
	Packet_dropped                  int `json:"packet_dropped"`
	Entry_learned                   int `json:"entry_learned"`
	Entry_count_overflow            int `json:"entry_count_overflow"`
	Exceed_drop_pkt_rate_clist      int `json:"exceed_drop_pkt_rate_clist"`
	Exceed_drop_conn_rate_clist     int `json:"exceed_drop_conn_rate_clist"`
	Exceed_drop_conn_limit_clist    int `json:"exceed_drop_conn_limit_clist"`
	Exceed_drop_kbit_rate_clist     int `json:"exceed_drop_kbit_rate_clist"`
	Exceed_drop_kbit_rate_clist_pkt int `json:"exceed_drop_kbit_rate_clist_pkt"`
	Exceed_drop_frag_rate_clist     int `json:"exceed_drop_frag_rate_clist"`
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port/zone-service-other/" + p.PortOther + "+" + p.Protocol + "/src-based-policy/" + p.SrcBasedPolicyName + "/policy-class-list/" + p.ClassListName + "/stats"
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats, error) {
	logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats
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
