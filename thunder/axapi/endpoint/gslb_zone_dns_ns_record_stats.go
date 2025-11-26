package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type GslbZoneDnsNsRecordStats struct {
	NsName string `json:"ns-name"`

	Stats GslbZoneDnsNsRecordStatsStats `json:"stats"`

	Zone_name string
}
type DataGslbZoneDnsNsRecordStats struct {
	DtGslbZoneDnsNsRecordStats GslbZoneDnsNsRecordStats `json:"dns-ns-record"`
}

type GslbZoneDnsNsRecordStatsStats struct {
	Hits int `json:"hits"`
}

func (p *GslbZoneDnsNsRecordStats) GetId() string {
	return "1"
}

func (p *GslbZoneDnsNsRecordStats) getPath() string {

	return "gslb/zone/" + p.Zone_name + "/dns-ns-record/" + p.NsName + "/stats"
}

func (p *GslbZoneDnsNsRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneDnsNsRecordStats, error) {
	logger.Println("GslbZoneDnsNsRecordStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbZoneDnsNsRecordStats
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
