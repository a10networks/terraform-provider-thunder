package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SysUtStateNextStateCaseActionTcpFlags struct {
	Inst struct {
		Ack int `json:"ack"`

		Cwr int `json:"cwr"`

		Ece int `json:"ece"`

		Fin int `json:"fin"`

		Psh int `json:"psh"`

		Rst int `json:"rst"`

		Syn int `json:"syn"`

		Urg int `json:"urg"`

		Uuid string `json:"uuid"`

		State_name string

		Direction string

		CaseNumber string

		Next_state_name string
	} `json:"flags"`
}

func (p *SysUtStateNextStateCaseActionTcpFlags) GetId() string {
	return "1"
}

func (p *SysUtStateNextStateCaseActionTcpFlags) getPath() string {
	return "sys-ut/state/" + p.Inst.State_name + "/next-state/" + p.Inst.Next_state_name + "/case/" + p.Inst.CaseNumber + "/action/" + p.Inst.Direction + "/tcp/flags"
}

func (p *SysUtStateNextStateCaseActionTcpFlags) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpFlags::Post")
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

func (p *SysUtStateNextStateCaseActionTcpFlags) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpFlags::Get")
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
func (p *SysUtStateNextStateCaseActionTcpFlags) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpFlags::Put")
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

func (p *SysUtStateNextStateCaseActionTcpFlags) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionTcpFlags::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
