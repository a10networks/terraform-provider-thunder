package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type SlbTemplateDnsCategoryLookup struct {
	Inst struct {
		CategoryName string `json:"category-name"`

		Drop int `json:"drop"`

		Permit int `json:"permit"`

		Respond int `json:"respond"`

		RespondCnameStr string `json:"respond-cname-str"`

		RespondIpAddr string `json:"respond-ip-addr"`

		RespondIpv6Addr string `json:"respond-ipv6-addr"`

		RespondNxdomain int `json:"respond-nxdomain"`

		ResponseTtl int `json:"response-ttl" dval:"300"`

		Uuid string `json:"uuid"`

		Dns_name string
	} `json:"category-lookup"`
}

func (p *SlbTemplateDnsCategoryLookup) GetId() string {
	return url.QueryEscape(p.Inst.CategoryName)
}

func (p *SlbTemplateDnsCategoryLookup) getPath() string {
	return "slb/template/dns/" + p.Inst.Dns_name + "/category-lookup"
}

func (p *SlbTemplateDnsCategoryLookup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsCategoryLookup::Post")
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

func (p *SlbTemplateDnsCategoryLookup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsCategoryLookup::Get")
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
func (p *SlbTemplateDnsCategoryLookup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsCategoryLookup::Put")
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

func (p *SlbTemplateDnsCategoryLookup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsCategoryLookup::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
