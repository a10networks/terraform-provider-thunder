package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type Cgnv6GlobalDomainList struct {
	Inst struct {
		AaaaQuery string `json:"aaaa-query" dval:"enable"`

		FailInterval int `json:"fail-interval"`

		Interval int `json:"interval"`

		Uuid string `json:"uuid"`
	} `json:"domain-list"`
}

func (p *Cgnv6GlobalDomainList) GetId() string {
	return "1"
}

func (p *Cgnv6GlobalDomainList) getPath() string {
	return "cgnv6/global/domain-list"
}

func (p *Cgnv6GlobalDomainList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6GlobalDomainList::Post")
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

func (p *Cgnv6GlobalDomainList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6GlobalDomainList::Get")
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
func (p *Cgnv6GlobalDomainList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6GlobalDomainList::Put")
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

func (p *Cgnv6GlobalDomainList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6GlobalDomainList::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
