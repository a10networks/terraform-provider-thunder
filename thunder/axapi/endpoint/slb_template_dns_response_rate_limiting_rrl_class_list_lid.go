package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type SlbTemplateDnsResponseRateLimitingRrlClassListLid struct {
	Inst struct {
		LidAction string `json:"lid-action" dval:"rate-limit"`

		LidEnableLog int `json:"lid-enable-log"`

		LidMatchSubnet string `json:"lid-match-subnet" dval:"255.255.255.255"`

		LidMatchSubnetV6 int `json:"lid-match-subnet-v6" dval:"128"`

		LidNxResponseRate int `json:"lid-nx-response-rate" dval:"5"`

		LidResponseRate int `json:"lid-response-rate" dval:"5"`

		LidSlipRate int `json:"lid-slip-rate"`

		LidSrcIpOnly int `json:"lid-src-ip-only"`

		LidTcRate int `json:"lid-tc-rate"`

		LidWindow int `json:"lid-window" dval:"1"`

		Lidnum int `json:"lidnum"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Dns_name string

		Rrl_class_list_name string
	} `json:"lid"`
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) GetId() string {
	return strconv.Itoa(p.Inst.Lidnum)
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) getPath() string {
	return "slb/template/dns/" + p.Inst.Dns_name + "/response-rate-limiting/rrl-class-list/" + p.Inst.Rrl_class_list_name + "/lid"
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassListLid::Post")
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

func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassListLid::Get")
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
func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassListLid::Put")
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

func (p *SlbTemplateDnsResponseRateLimitingRrlClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassListLid::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
