package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type PkiDelete struct {
	Inst struct {
		Ca string `json:"ca"`

		CertName string `json:"cert-name"`

		Crl string `json:"crl"`

		Csr string `json:"csr"`

		PrivateKey string `json:"private-key"`
	} `json:"delete"`
}

func (p *PkiDelete) GetId() string {
	return "1"
}

func (p *PkiDelete) getPath() string {
	return "pki/delete"
}

func (p *PkiDelete) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("PkiDelete::Post")
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

func (p *PkiDelete) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("PkiDelete::Get")
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
func (p *PkiDelete) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("PkiDelete::Put")
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

func (p *PkiDelete) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("PkiDelete::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
