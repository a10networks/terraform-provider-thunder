package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P4_70
type PkiCreateOper struct {
	Inst struct {
		Bits          string `json:"bits"`
		CertType      string `json:"cert-type" dval:"rsa"`
		CommonName    string `json:"common-name"`
		Country       string `json:"country"`
		CsrGenerate   int    `json:"csr-generate"`
		Digest        string `json:"digest" dval:"sha1"`
		Division      string `json:"division"`
		Email         string `json:"email"`
		Filename      string `json:"filename"`
		Locality      string `json:"locality"`
		Organization  string `json:"organization"`
		Rootca        int    `json:"rootca"`
		Secured       int    `json:"secured"`
		StateProvince string `json:"state-province"`
		V3Request     int    `json:"v3-request"`
		ValidDays     int    `json:"valid-days"`
	} `json:"create-oper"`
}

func (p *PkiCreateOper) GetId() string {
	return "1"
}

func (p *PkiCreateOper) getPath() string {
	return "pki/create-oper"
}

func (p *PkiCreateOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("PkiCreateOper::Post")
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

func (p *PkiCreateOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("PkiCreateOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *PkiCreateOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("PkiCreateOper::Put")
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

func (p *PkiCreateOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("PkiCreateOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
