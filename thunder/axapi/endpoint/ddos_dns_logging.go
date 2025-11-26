package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosDnsLogging struct {
	Inst struct {
		Disable int `json:"disable"`

		DnsLoggingProtocol string `json:"dns-logging-protocol"`

		DnsLoggingRequestSection string `json:"dns-logging-request-section"`

		DnsLoggingType string `json:"dns-logging-type"`

		Name string `json:"name"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"dns-logging"`
}

func (p *DdosDnsLogging) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDnsLogging) getPath() string {
	return "ddos/dns-logging"
}

func (p *DdosDnsLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsLogging::Post")
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

func (p *DdosDnsLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsLogging::Get")
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
func (p *DdosDnsLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsLogging::Put")
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

func (p *DdosDnsLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsLogging::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
