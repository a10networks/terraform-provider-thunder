package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsLabelCountFilter struct {
	Inst struct {
		DropLogEnable int `json:"drop-log-enable"`

		LabelCountFilterAction string `json:"label-count-filter-action" dval:"drop"`

		MaxFqdnLabelCount int `json:"max-fqdn-label-count"`

		MinFqdnLabelCount int `json:"min-fqdn-label-count"`

		Uuid string `json:"uuid"`

		Dns_name string
	} `json:"label-count-filter"`
}

func (p *SlbTemplateDnsLabelCountFilter) GetId() string {
	return "1"
}

func (p *SlbTemplateDnsLabelCountFilter) getPath() string {
	return "slb/template/dns/" + p.Inst.Dns_name + "/label-count-filter"
}

func (p *SlbTemplateDnsLabelCountFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelCountFilter::Post")
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

func (p *SlbTemplateDnsLabelCountFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelCountFilter::Get")
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
func (p *SlbTemplateDnsLabelCountFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelCountFilter::Put")
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

func (p *SlbTemplateDnsLabelCountFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsLabelCountFilter::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
