package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZoneSrcIpFilteringStats struct {
	Stats DdosDstZoneSrcIpFilteringStatsStats `json:"stats"`

	ZoneName string
}
type DataDdosDstZoneSrcIpFilteringStats struct {
	DtDdosDstZoneSrcIpFilteringStats DdosDstZoneSrcIpFilteringStats `json:"src-ip-filtering"`
}

type DdosDstZoneSrcIpFilteringStatsStats struct {
	ClassList1Match int `json:"class-list-1-match"`
	ClassList2Match int `json:"class-list-2-match"`
	ClassList3Match int `json:"class-list-3-match"`
	ClassList4Match int `json:"class-list-4-match"`
	ClassList5Match int `json:"class-list-5-match"`
	ClassList6Match int `json:"class-list-6-match"`
	ClassList7Match int `json:"class-list-7-match"`
	ClassList8Match int `json:"class-list-8-match"`
}

func (p *DdosDstZoneSrcIpFilteringStats) GetId() string {
	return "1"
}

func (p *DdosDstZoneSrcIpFilteringStats) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/src-ip-filtering/stats"
}

func (p *DdosDstZoneSrcIpFilteringStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcIpFilteringStats, error) {
	logger.Println("DdosDstZoneSrcIpFilteringStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneSrcIpFilteringStats
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
