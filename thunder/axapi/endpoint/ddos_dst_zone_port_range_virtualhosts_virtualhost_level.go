package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortRangeVirtualhostsVirtualhostLevel struct {
	Inst struct {
		GlidAction string `json:"glid-action"`

		LevelNum string `json:"level-num"`

		SrcDefaultGlid string `json:"src-default-glid"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		ZoneTemplate DdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate `json:"zone-template"`

		ZoneName string

		PortRangeStart string

		Protocol string

		Vhost string

		PortRangeEnd string
	} `json:"level"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) GetId() string {
	return p.Inst.LevelNum
}

func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port-range/" + p.Inst.PortRangeStart + "+" + p.Inst.PortRangeEnd + "+" + p.Inst.Protocol + "/virtualhosts/virtualhost/" + p.Inst.Vhost + "/level"
}

func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhostsVirtualhostLevel::Post")
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

func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhostsVirtualhostLevel::Get")
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
func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhostsVirtualhostLevel::Put")
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

func (p *DdosDstZonePortRangeVirtualhostsVirtualhostLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhostsVirtualhostLevel::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
