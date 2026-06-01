package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type GslbServiceMatchingRules struct {
	Inst struct {
		Disable int `json:"disable"`

		HitcountEnable int `json:"hitcount-enable"`

		RuleList []GslbServiceMatchingRulesRuleList `json:"rule-list"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Zone string `json:"zone"`
	} `json:"service-matching-rules"`
}

type GslbServiceMatchingRulesRuleList struct {
	SeqNum            int                                           `json:"seq-num"`
	DomainMatchType   string                                        `json:"domain-match-type"`
	DomainMatchString string                                        `json:"domain-match-string"`
	SrcIpv4           string                                        `json:"src-ipv4"`
	SrcIpv6           string                                        `json:"src-ipv6"`
	HealthState       []GslbServiceMatchingRulesRuleListHealthState `json:"health-state"`
	Service           string                                        `json:"service"`
	Uuid              string                                        `json:"uuid"`
	UserTag           string                                        `json:"user-tag"`
}

type GslbServiceMatchingRulesRuleListHealthState struct {
	GslbSite       string `json:"gslb-site"`
	SiteState      string `json:"site-state"`
	GslbServiceIp  string `json:"gslb-service-ip"`
	ServiceIpState string `json:"service-ip-state"`
	SlbServer      string `json:"slb-server"`
	SlbSvrState    string `json:"slb-svr-state"`
}

func (p *GslbServiceMatchingRules) GetId() string {
	return url.QueryEscape(p.Inst.Zone)
}

func (p *GslbServiceMatchingRules) getPath() string {
	return "gslb/service-matching-rules"
}

func (p *GslbServiceMatchingRules) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRules::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *GslbServiceMatchingRules) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRules::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *GslbServiceMatchingRules) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRules::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *GslbServiceMatchingRules) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRules::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
