package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type GslbPolicyConnectionCountBySite struct {
	Inst struct {
		ConnectionCountEnable int `json:"connection-count-enable"`

		Uuid string `json:"uuid"`

		Policy_name string
	} `json:"connection-count-by-site"`
}

func (p *GslbPolicyConnectionCountBySite) GetId() string {
	return "1"
}

func (p *GslbPolicyConnectionCountBySite) getPath() string {
	return "gslb/policy/" + p.Inst.Policy_name + "/connection-count-by-site"
}

func (p *GslbPolicyConnectionCountBySite) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyConnectionCountBySite::Post")
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

func (p *GslbPolicyConnectionCountBySite) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyConnectionCountBySite::Get")
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
func (p *GslbPolicyConnectionCountBySite) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyConnectionCountBySite::Put")
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

func (p *GslbPolicyConnectionCountBySite) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicyConnectionCountBySite::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
