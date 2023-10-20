package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type InterfaceAvailableEthListOper struct {
	Oper InterfaceAvailableEthListOperOper `json:"oper"`
}

type InterfaceAvailableEthListt struct {
	InterfaceAvailableEthList InterfaceAvailableEthListOper `json:"available-eth-list"`
}

type InterfaceAvailableEthListOperOper struct {
	Tot_num_of_ports int                                       `json:"tot_num_of_ports"`
	IfList           []InterfaceAvailableEthListOperOperIfList `json:"if-list"`
}

type InterfaceAvailableEthListOperOperIfList struct {
	IfType   string `json:"IF-Type"`
	IfNum    int    `json:"IF-Num"`
	IfStatus string `json:"IF-Status"`
	State    string `json:"state"`
}

func (p *InterfaceAvailableEthListOper) GetId() string {
	return "1"
}

func (p *InterfaceAvailableEthListOper) getPath() string {
	return "interface/available-eth-list/oper"
}

func (p *InterfaceAvailableEthListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (InterfaceAvailableEthListt, error) {
	logger.Println("InterfaceAvailableEthListOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var availableList InterfaceAvailableEthListt

	if err == nil {
		err = json.Unmarshal(axResp, &availableList)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return availableList, err
}
