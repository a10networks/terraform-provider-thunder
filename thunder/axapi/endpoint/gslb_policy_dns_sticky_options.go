package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type GslbPolicyDnsStickyOptions struct {
	Inst struct {
		EdnsClientSubnet int `json:"edns-client-subnet"`

		OnlyEcs int `json:"only-ecs"`

		Uuid string `json:"uuid"`

		Policy_name string
	} `json:"sticky-options"`
}

func (p *GslbPolicyDnsStickyOptions) GetId() string {
	return "1"
}

func (p *GslbPolicyDnsStickyOptions) getPath() string {
	return "gslb/policy/" + p.Inst.Policy_name + "/dns/sticky-options"
}

func (p *GslbPolicyDnsStickyOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyDnsStickyOptions::Post")
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

func (p *GslbPolicyDnsStickyOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyDnsStickyOptions::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *GslbPolicyDnsStickyOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyDnsStickyOptions::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *GslbPolicyDnsStickyOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyDnsStickyOptions::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
