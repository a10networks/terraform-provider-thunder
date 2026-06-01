package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsLabelLengthFilter struct {
	Inst struct {
		DropLogEnable int `json:"drop-log-enable"`

		FqdnLabelLength []SlbTemplateDnsLabelLengthFilterFqdnLabelLength `json:"fqdn-label-length"`

		LabelLengthFilterAction string `json:"label-length-filter-action" dval:"drop"`

		Uuid string `json:"uuid"`

		Dns_name string
	} `json:"label-length-filter"`
}

type SlbTemplateDnsLabelLengthFilterFqdnLabelLength struct {
	Length int `json:"length"`
	Suffix int `json:"suffix"`
}

func (p *SlbTemplateDnsLabelLengthFilter) GetId() string {
	return "1"
}

func (p *SlbTemplateDnsLabelLengthFilter) getPath() string {
	return "slb/template/dns/" + p.Inst.Dns_name + "/label-length-filter"
}

func (p *SlbTemplateDnsLabelLengthFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelLengthFilter::Post")
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

func (p *SlbTemplateDnsLabelLengthFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelLengthFilter::Get")
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
func (p *SlbTemplateDnsLabelLengthFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelLengthFilter::Put")
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

func (p *SlbTemplateDnsLabelLengthFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelLengthFilter::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
