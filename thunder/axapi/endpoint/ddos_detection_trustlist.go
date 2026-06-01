package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDetectionTrustlist struct {
	Inst struct {
		Uuid string `json:"uuid"`

		V4ClassList string `json:"v4-class-list"`

		V6ClassList string `json:"v6-class-list"`
	} `json:"trustlist"`
}

func (p *DdosDetectionTrustlist) GetId() string {
	return "1"
}

func (p *DdosDetectionTrustlist) getPath() string {
	return "ddos/detection/trustlist"
}

func (p *DdosDetectionTrustlist) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionTrustlist::Post")
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

func (p *DdosDetectionTrustlist) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionTrustlist::Get")
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
func (p *DdosDetectionTrustlist) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionTrustlist::Put")
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

func (p *DdosDetectionTrustlist) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionTrustlist::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
