package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemDomainListSettingsOper struct {
	Oper SystemDomainListSettingsOperOper `json:"oper"`
}
type DataSystemDomainListSettingsOper struct {
	DtSystemDomainListSettingsOper SystemDomainListSettingsOper `json:"domain-list-settings"`
}

type SystemDomainListSettingsOperOper struct {
	TotalEntryNum     int `json:"total-entry-num"`
	InitedEntryNum    int `json:"inited-entry-num"`
	InProgEntryNum    int `json:"in-prog-entry-num"`
	ExceptionEntryNum int `json:"exception-entry-num"`
	SerialUpdated     int `json:"serial-updated"`
	XfrStartFail      int `json:"xfr-start-fail"`
	XfrTimeoutFail    int `json:"xfr-timeout-fail"`
	XfrParseFail      int `json:"xfr-parse-fail"`
}

func (p *SystemDomainListSettingsOper) GetId() string {
	return "1"
}

func (p *SystemDomainListSettingsOper) getPath() string {
	return "system/domain-list-settings/oper"
}

func (p *SystemDomainListSettingsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemDomainListSettingsOper, error) {
	logger.Println("SystemDomainListSettingsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSystemDomainListSettingsOper
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
