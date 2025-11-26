package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SysUtStateNextStateCaseActionL2 struct {
	Inst struct {
		Ethertype int `json:"ethertype"`

		MacList []SysUtStateNextStateCaseActionL2MacList `json:"mac-list"`

		Protocol string `json:"protocol" dval:"ipv4"`

		Uuid string `json:"uuid"`

		Value int `json:"value"`

		Vlan int `json:"vlan"`

		State_name string

		Direction string

		CaseNumber string

		Next_state_name string
	} `json:"l2"`
}

type SysUtStateNextStateCaseActionL2MacList struct {
	SrcDst        string `json:"src-dst"`
	AddressType   string `json:"address-type"`
	VirtualServer string `json:"virtual-server"`
	NatPool       string `json:"nat-pool"`
	Ethernet      int    `json:"ethernet"`
	Ve            int    `json:"ve"`
	Trunk         int    `json:"trunk"`
	Value         string `json:"value"`
	Uuid          string `json:"uuid"`
}

func (p *SysUtStateNextStateCaseActionL2) GetId() string {
	return "1"
}

func (p *SysUtStateNextStateCaseActionL2) getPath() string {
	return "sys-ut/state/" + p.Inst.State_name + "/next-state/" + p.Inst.Next_state_name + "/case/" + p.Inst.CaseNumber + "/action/" + p.Inst.Direction + "/l2"
}

func (p *SysUtStateNextStateCaseActionL2) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionL2::Post")
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

func (p *SysUtStateNextStateCaseActionL2) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionL2::Get")
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
func (p *SysUtStateNextStateCaseActionL2) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionL2::Put")
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

func (p *SysUtStateNextStateCaseActionL2) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseActionL2::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
