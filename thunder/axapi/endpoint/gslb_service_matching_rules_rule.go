package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type GslbServiceMatchingRulesRule struct {
	Inst struct {
		DomainMatchString string `json:"domain-match-string"`

		DomainMatchType string `json:"domain-match-type"`

		HealthState []GslbServiceMatchingRulesRuleHealthState `json:"health-state"`

		SeqNum int `json:"seq-num"`

		Service string `json:"service"`

		SrcIpv4 string `json:"src-ipv4"`

		SrcIpv6 string `json:"src-ipv6"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Zone string
	} `json:"rule"`
}

type GslbServiceMatchingRulesRuleHealthState struct {
	GslbSite       string `json:"gslb-site"`
	SiteState      string `json:"site-state"`
	GslbServiceIp  string `json:"gslb-service-ip"`
	ServiceIpState string `json:"service-ip-state"`
	SlbServer      string `json:"slb-server"`
	SlbSvrState    string `json:"slb-svr-state"`
}

func (p *GslbServiceMatchingRulesRule) GetId() string {
	return strconv.Itoa(p.Inst.SeqNum)
}

func (p *GslbServiceMatchingRulesRule) getPath() string {
	return "gslb/service-matching-rules/" + p.Inst.Zone + "/rule"
}

func (p *GslbServiceMatchingRulesRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRulesRule::Post")
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

func (p *GslbServiceMatchingRulesRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRulesRule::Get")
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
func (p *GslbServiceMatchingRulesRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRulesRule::Put")
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

func (p *GslbServiceMatchingRulesRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbServiceMatchingRulesRule::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
