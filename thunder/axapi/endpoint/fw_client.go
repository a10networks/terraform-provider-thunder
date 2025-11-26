package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type FwClient struct {
	Inst struct {
		Name string `json:"name"`

		Type string `json:"type"`

		Uuid string `json:"uuid"`
	} `json:"client"`
}

func (p *FwClient) GetId() string {
	return "1"
}

func (p *FwClient) getPath() string {
	return "fw/client"
}

func (p *FwClient) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwClient::Post")
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

func (p *FwClient) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwClient::Get")
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
func (p *FwClient) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwClient::Put")
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

func (p *FwClient) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwClient::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
