package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type AccountingThreatLogs struct {
	Inst struct {
		Check int `json:"check"`

		Days int `json:"days" dval:"30"`

		Uuid string `json:"uuid"`
	} `json:"threat-logs"`
}

func (p *AccountingThreatLogs) GetId() string {
	return "1"
}

func (p *AccountingThreatLogs) getPath() string {
	return "accounting/threat-logs"
}

func (p *AccountingThreatLogs) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("AccountingThreatLogs::Post")
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

func (p *AccountingThreatLogs) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AccountingThreatLogs::Get")
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
func (p *AccountingThreatLogs) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("AccountingThreatLogs::Put")
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

func (p *AccountingThreatLogs) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AccountingThreatLogs::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
