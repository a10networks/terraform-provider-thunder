package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type HealthSourceNat struct {
	Inst struct {
		EnableVrrpAMode int `json:"enable-vrrp-a-mode"`

		Ethernet int `json:"ethernet"`

		Interface string `json:"interface"`

		SamplingEnable []HealthSourceNatSamplingEnable `json:"sampling-enable"`

		SmartNatPrecedence int `json:"smart-nat-precedence"`

		SmartNatVrid int `json:"smart-nat-vrid"`

		SourceNatPool string `json:"source-nat-pool"`

		SourceNatPoolV6 string `json:"source-nat-pool-v6"`

		Trunk int `json:"trunk"`

		Uuid string `json:"uuid"`

		Ve int `json:"ve"`
	} `json:"source-nat"`
}

type HealthSourceNatSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *HealthSourceNat) GetId() string {
	return "1"
}

func (p *HealthSourceNat) getPath() string {
	return "health/source-nat"
}

func (p *HealthSourceNat) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("HealthSourceNat::Post")
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

func (p *HealthSourceNat) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("HealthSourceNat::Get")
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
func (p *HealthSourceNat) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("HealthSourceNat::Put")
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

func (p *HealthSourceNat) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("HealthSourceNat::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
