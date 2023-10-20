package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type GslbTemplateCsv struct {
	Inst struct {
		CsvName        string                          `json:"csv-name"`
		DelimChar      string                          `json:"delim-char" dval:","`
		DelimNum       int                             `json:"delim-num" dval:"44"`
		Ipv6Enable     int                             `json:"ipv6-enable"`
		MultipleFields []GslbTemplateCsvMultipleFields `json:"multiple-fields"`
		UserTag        string                          `json:"user-tag"`
		Uuid           string                          `json:"uuid"`
	} `json:"csv"`
}

type GslbTemplateCsvMultipleFields struct {
	Field   int    `json:"field"`
	CsvType string `json:"csv-type"`
}

func (p *GslbTemplateCsv) GetId() string {
	return p.Inst.CsvName
}

func (p *GslbTemplateCsv) getPath() string {
	return "gslb/template/csv"
}

func (p *GslbTemplateCsv) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbTemplateCsv::Post")
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

func (p *GslbTemplateCsv) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbTemplateCsv::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *GslbTemplateCsv) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbTemplateCsv::Put")
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

func (p *GslbTemplateCsv) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbTemplateCsv::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
