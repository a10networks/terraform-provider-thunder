package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DebugCaptcha struct {
	Inst struct {
		Dumy int `json:"dumy"`

		Uuid string `json:"uuid"`
	} `json:"captcha"`
}

func (p *DebugCaptcha) GetId() string {
	return "1"
}

func (p *DebugCaptcha) getPath() string {
	return "debug/captcha"
}

func (p *DebugCaptcha) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugCaptcha::Post")
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

func (p *DebugCaptcha) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugCaptcha::Get")
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
func (p *DebugCaptcha) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugCaptcha::Put")
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

func (p *DebugCaptcha) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugCaptcha::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
