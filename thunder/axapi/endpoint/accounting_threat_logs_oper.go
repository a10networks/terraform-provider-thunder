package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type AccountingThreatLogsOper struct {
	Oper AccountingThreatLogsOperOper `json:"oper"`
}
type DataAccountingThreatLogsOper struct {
	DtAccountingThreatLogsOper AccountingThreatLogsOper `json:"threat-logs"`
}

type AccountingThreatLogsOperOper struct {
	Status string `json:"status"`
	Result int    `json:"result"`
	Msg    string `json:"msg"`
}

func (p *AccountingThreatLogsOper) GetId() string {
	return "1"
}

func (p *AccountingThreatLogsOper) getPath() string {
	return "accounting/threat-logs/oper"
}

func (p *AccountingThreatLogsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAccountingThreatLogsOper, error) {
	logger.Println("AccountingThreatLogsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataAccountingThreatLogsOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
