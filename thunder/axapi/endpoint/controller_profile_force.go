package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ControllerProfileForce struct {
	Inst struct {
		Deregister int `json:"deregister"`
	} `json:"force"`
}

func (p *ControllerProfileForce) GetId() string {
	return "1"
}

func (p *ControllerProfileForce) getPath() string {
	return "controller/profile/force"
}

func (p *ControllerProfileForce) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfileForce::Post")
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

func (p *ControllerProfileForce) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfileForce::Get")
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
func (p *ControllerProfileForce) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfileForce::Put")
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

func (p *ControllerProfileForce) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfileForce::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
