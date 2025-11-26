package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ImportPeriodicWsdl struct {
	Inst struct {
		Period int `json:"period"`

		RemoteFile string `json:"remote-file"`

		UseMgmtPort int `json:"use-mgmt-port"`

		Uuid string `json:"uuid"`

		Wsdl string `json:"wsdl"`
	} `json:"wsdl"`
}

func (p *ImportPeriodicWsdl) GetId() string {
	return p.Inst.Wsdl
}

func (p *ImportPeriodicWsdl) getPath() string {
	return "import-periodic/wsdl"
}

func (p *ImportPeriodicWsdl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ImportPeriodicWsdl::Post")
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

func (p *ImportPeriodicWsdl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ImportPeriodicWsdl::Get")
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
func (p *ImportPeriodicWsdl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ImportPeriodicWsdl::Put")
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

func (p *ImportPeriodicWsdl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ImportPeriodicWsdl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
