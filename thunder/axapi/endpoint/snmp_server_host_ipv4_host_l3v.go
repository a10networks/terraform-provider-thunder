package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P4_81
type SnmpServerHostIpv4HostL3v struct {
	Inst struct {
		Ipv4Addr  string `json:"ipv4-addr"`
	} `json:"ipv4-host"`
}

func (p *SnmpServerHostIpv4HostL3v) GetId() string {
	return p.Inst.Ipv4Addr
}

func (p *SnmpServerHostIpv4HostL3v) getPath() string {
	return "snmp-server/host/ipv4-host"
}

func (p *SnmpServerHostIpv4HostL3v) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerHostIpv4HostL3v::Post")
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

func (p *SnmpServerHostIpv4HostL3v) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerHostIpv4HostL3v::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SnmpServerHostIpv4HostL3v) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerHostIpv4HostL3v::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *SnmpServerHostIpv4HostL3v) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerHostIpv4HostL3v::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
