package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DebugRmthttp struct {
	Inst struct {
		Dummy int `json:"dummy"`

		Uuid string `json:"uuid"`
	} `json:"rmthttp"`
}

func (p *DebugRmthttp) GetId() string {
	return "1"
}

func (p *DebugRmthttp) getPath() string {
	return "debug/rmthttp"
}

func (p *DebugRmthttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugRmthttp::Post")
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

func (p *DebugRmthttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugRmthttp::Get")
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
func (p *DebugRmthttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DebugRmthttp::Put")
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

func (p *DebugRmthttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DebugRmthttp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
