package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type RuleSetRulesByZoneStats struct {
	Stats RuleSetRulesByZoneStatsStats `json:"stats"`

	Rule_set_name string
}
type DataRuleSetRulesByZoneStats struct {
	DtRuleSetRulesByZoneStats RuleSetRulesByZoneStats `json:"rules-by-zone"`
}

type RuleSetRulesByZoneStatsStats struct {
	Dummy int `json:"dummy"`
}

func (p *RuleSetRulesByZoneStats) GetId() string {
	return "1"
}

func (p *RuleSetRulesByZoneStats) getPath() string {

	return "rule-set/" + p.Rule_set_name + "/rules-by-zone/stats"
}

func (p *RuleSetRulesByZoneStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetRulesByZoneStats, error) {
	logger.Println("RuleSetRulesByZoneStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataRuleSetRulesByZoneStats
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
