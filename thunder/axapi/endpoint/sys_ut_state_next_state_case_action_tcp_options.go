package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SysUtStateNextStateCaseActionTcpOptions struct {
	Inst struct {
		Mss int `json:"mss"`

		Nop int `json:"nop"`

		SackType string `json:"sack-type"`

		TimeStampEnable int `json:"time-stamp-enable"`

		Uuid string `json:"uuid"`

		Wscale int `json:"wscale"`

		Next_state_name string

		Direction string

		CaseNumber string

		State_name string
	} `json:"options"`
}

func (p *SysUtStateNextStateCaseActionTcpOptions) GetId() string {
	return "1"
}

func (p *SysUtStateNextStateCaseActionTcpOptions) getPath() string {
	return "sys-ut/state/" + p.Inst.State_name + "/next-state/" + p.Inst.Next_state_name + "/case/" + p.Inst.CaseNumber + "/action/" + p.Inst.Direction + "/tcp/options"
}

func (p *SysUtStateNextStateCaseActionTcpOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpOptions::Post")
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

func (p *SysUtStateNextStateCaseActionTcpOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpOptions::Get")
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
func (p *SysUtStateNextStateCaseActionTcpOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpOptions::Put")
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

func (p *SysUtStateNextStateCaseActionTcpOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpOptions::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
