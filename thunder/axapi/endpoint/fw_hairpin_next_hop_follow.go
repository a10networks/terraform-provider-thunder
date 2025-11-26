package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type FwHairpinNextHopFollow struct {
	Inst struct {
		DisableLocalHairpin int `json:"disable-local-hairpin"`

		Ipv4 string `json:"ipv4"`

		Ipv6 string `json:"ipv6"`

		Uuid string `json:"uuid"`
	} `json:"next-hop-follow"`
}

func (p *FwHairpinNextHopFollow) GetId() string {
	return "1"
}

func (p *FwHairpinNextHopFollow) getPath() string {
	return "fw/hairpin/next-hop-follow"
}

func (p *FwHairpinNextHopFollow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwHairpinNextHopFollow::Post")
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

func (p *FwHairpinNextHopFollow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwHairpinNextHopFollow::Get")
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
func (p *FwHairpinNextHopFollow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwHairpinNextHopFollow::Put")
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

func (p *FwHairpinNextHopFollow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwHairpinNextHopFollow::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
