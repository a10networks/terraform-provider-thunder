package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P3_70
type PkiDelete struct {
	Inst struct {
		Ca         string `json:"ca,omitempty"`
		CertName   string `json:"cert-name,omitempty"`
		Crl        string `json:"crl,omitempty"`
		Csr        string `json:"csr,omitempty"`
		PrivateKey string `json:"private-key,omitempty"`
	} `json:"delete"`
}

func (p *PkiDelete) GetId() string {
	return "1"
}

func (p *PkiDelete) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("PkiDelete::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "pki/delete", payloadBytes, headers, logger)
	return err
}
