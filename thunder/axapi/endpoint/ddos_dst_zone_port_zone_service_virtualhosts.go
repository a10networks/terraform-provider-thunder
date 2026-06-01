package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortZoneServiceVirtualhosts struct {
	Inst struct {
		SourceTrackingAll int `json:"source-tracking-all"`

		Uuid string `json:"uuid"`

		VhostsConfig string `json:"vhosts-config"`

		VirtualhostList []DdosDstZonePortZoneServiceVirtualhostsVirtualhostList `json:"virtualhost-list"`

		ZoneName string

		PortNum string

		Protocol string
	} `json:"virtualhosts"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostList struct {
	Vhost              string                                                            `json:"vhost"`
	Servername         []DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername `json:"servername"`
	ServernameList     string                                                            `json:"servername-list"`
	ServernameMatchAny int                                                               `json:"servername-match-any"`
	ServernameNoSni    int                                                               `json:"servername-no-sni"`
	SourceTracking     string                                                            `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg      `json:"glid-cfg"`
	Deny               int                                                               `json:"deny"`
	Uuid               string                                                            `json:"uuid"`
	UserTag            string                                                            `json:"user-tag"`
	LevelList          []DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList  `json:"level-list"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList struct {
	LevelNum       string                                                                     `json:"level-num"`
	SrcDefaultGlid string                                                                     `json:"src-default-glid"`
	GlidAction     string                                                                     `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate `json:"zone-template"`
	Uuid           string                                                                     `json:"uuid"`
	UserTag        string                                                                     `json:"user-tag"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortZoneServiceVirtualhosts) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceVirtualhosts) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port/zone-service/" + p.Inst.PortNum + "+" + p.Inst.Protocol + "/virtualhosts"
}

func (p *DdosDstZonePortZoneServiceVirtualhosts) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhosts::Post")
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

func (p *DdosDstZonePortZoneServiceVirtualhosts) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhosts::Get")
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
func (p *DdosDstZonePortZoneServiceVirtualhosts) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhosts::Put")
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

func (p *DdosDstZonePortZoneServiceVirtualhosts) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhosts::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
