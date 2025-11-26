package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SlbTemplateDnsRecursiveDnsResolutionOper struct {
	Oper SlbTemplateDnsRecursiveDnsResolutionOperOper `json:"oper"`

	Dns_name string
}
type DataSlbTemplateDnsRecursiveDnsResolutionOper struct {
	DtSlbTemplateDnsRecursiveDnsResolutionOper SlbTemplateDnsRecursiveDnsResolutionOper `json:"recursive-dns-resolution"`
}

type SlbTemplateDnsRecursiveDnsResolutionOperOper struct {
	GwhcStatus      string `json:"gwhc-status"`
	GwhcUpRetries   int    `json:"gwhc-up-retries"`
	GwhcDownRetries int    `json:"gwhc-down-retries"`
}

func (p *SlbTemplateDnsRecursiveDnsResolutionOper) GetId() string {
	return "1"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionOper) getPath() string {

	return "slb/template/dns/" + p.Dns_name + "/recursive-dns-resolution/oper"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateDnsRecursiveDnsResolutionOper, error) {
	logger.Println("SlbTemplateDnsRecursiveDnsResolutionOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbTemplateDnsRecursiveDnsResolutionOper
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
