package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel struct {
	Inst struct {
		GlidAction string `json:"glid-action"`

		LevelNum string `json:"level-num"`

		SrcDefaultGlid string `json:"src-default-glid"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		ZoneTemplate DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate `json:"zone-template"`

		ZoneName string

		Protocol string

		Vhost string

		PortNum string
	} `json:"level"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) GetId() string {
	return p.Inst.LevelNum
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port/zone-service/" + p.Inst.PortNum + "+" + p.Inst.Protocol + "/virtualhosts/virtualhost/" + p.Inst.Vhost + "/level"
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel::Post")
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

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel::Put")
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

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
