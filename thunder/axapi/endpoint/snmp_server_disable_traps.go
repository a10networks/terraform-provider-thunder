package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P4_81
type SnmpServerDisableTraps struct {
	Inst struct {
		All       int    `json:"all"`
		Gslb      int    `json:"gslb"`
		Slb       int    `json:"slb"`
		SlbChange int    `json:"slb-change"`
		Snmp      int    `json:"snmp"`
		Uuid      string `json:"uuid"`
		VrrpA     int    `json:"vrrp-a"`
	} `json:"traps"`
}

func (p *SnmpServerDisableTraps) GetId() string {
	return "1"
}

func (p *SnmpServerDisableTraps) getPath() string {
	return "snmp-server/disable/traps"
}

func (p *SnmpServerDisableTraps) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerDisableTraps::Post")
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

func (p *SnmpServerDisableTraps) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerDisableTraps::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SnmpServerDisableTraps) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerDisableTraps::Put")
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

func (p *SnmpServerDisableTraps) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerDisableTraps::Delete")
	headers := axapi.GenRequestHeader(authToken)
	//this endpoint does not support DELETE method
	_, _, err := axapi.SendPut(host, p.getPath(), "", []byte("{\"traps\":{\"all\":0}}"), headers, logger)
	return err
}
