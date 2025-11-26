package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SlbTemplatePolicyForwardPolicySourceStats struct {
	Name string `json:"name"`

	Stats SlbTemplatePolicyForwardPolicySourceStatsStats `json:"stats"`

	Policy_name string
}
type DataSlbTemplatePolicyForwardPolicySourceStats struct {
	DtSlbTemplatePolicyForwardPolicySourceStats SlbTemplatePolicyForwardPolicySourceStats `json:"source"`
}

type SlbTemplatePolicyForwardPolicySourceStatsStats struct {
	Hits                     int `json:"hits"`
	DestinationMatchNotFound int `json:"destination-match-not-found"`
	NoHostInfo               int `json:"no-host-info"`
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) GetId() string {
	return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) getPath() string {

	return "slb/template/policy/" + p.Policy_name + "/forward-policy/source/" + p.Name + "/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceStats, error) {
	logger.Println("SlbTemplatePolicyForwardPolicySourceStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbTemplatePolicyForwardPolicySourceStats
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
