package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectTrustlist struct {
	Inst struct {
		Uuid string `json:"uuid"`

		V4ClassList string `json:"v4-class-list"`

		V6ClassList string `json:"v6-class-list"`

		ObjectName string
	} `json:"trustlist"`
}

func (p *DdosNetworkObjectTrustlist) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectTrustlist) getPath() string {
	return "ddos/network-object/" + p.Inst.ObjectName + "/trustlist"
}

func (p *DdosNetworkObjectTrustlist) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTrustlist::Post")
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

func (p *DdosNetworkObjectTrustlist) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTrustlist::Get")
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
func (p *DdosNetworkObjectTrustlist) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTrustlist::Put")
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

func (p *DdosNetworkObjectTrustlist) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosNetworkObjectTrustlist::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
