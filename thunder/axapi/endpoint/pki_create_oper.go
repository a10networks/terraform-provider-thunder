package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type PkiCreateOper struct {
	Bits string `json:"bits" dval:"1024"`

	CertType string `json:"cert-type" dval:"rsa"`

	CommonName string `json:"common-name"`

	Country string `json:"country"`

	CsrGenerate int `json:"csr-generate"`

	Digest string `json:"digest" dval:"sha1"`

	Division string `json:"division"`

	Email string `json:"email"`

	Filename string `json:"filename"`

	Locality string `json:"locality"`

	Organization string `json:"organization"`

	Rootca int `json:"rootca"`

	Secured int `json:"secured"`

	StateProvince string `json:"state-province"`

	V3Request int `json:"v3-request"`

	ValidDays int `json:"valid-days" dval:"730"`
}
type DataPkiCreateOper struct {
	DtPkiCreateOper PkiCreateOper `json:"create-oper"`
}

func (p *PkiCreateOper) GetId() string {
	return "1"
}

func (p *PkiCreateOper) getPath() string {
	return "pki/create-oper"
}

func (p *PkiCreateOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPkiCreateOper, error) {
	logger.Println("PkiCreateOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataPkiCreateOper
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
