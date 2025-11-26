package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortRangeVirtualhosts struct {
	Inst struct {
		SourceTrackingAll int `json:"source-tracking-all"`

		Uuid string `json:"uuid"`

		VhostsConfig string `json:"vhosts-config"`

		VirtualhostList []DdosDstZonePortRangeVirtualhostsVirtualhostList `json:"virtualhost-list"`

		Protocol string

		ZoneName string

		PortRangeEnd string

		PortRangeStart string
	} `json:"virtualhosts"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostList struct {
	Vhost              string                                                      `json:"vhost"`
	Servername         []DdosDstZonePortRangeVirtualhostsVirtualhostListServername `json:"servername"`
	ServernameList     string                                                      `json:"servername-list"`
	ServernameMatchAny int                                                         `json:"servername-match-any"`
	SourceTracking     string                                                      `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg      `json:"glid-cfg"`
	Deny               int                                                         `json:"deny"`
	Uuid               string                                                      `json:"uuid"`
	UserTag            string                                                      `json:"user-tag"`
	LevelList          []DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList  `json:"level-list"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListServername struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList struct {
	LevelNum       string                                                               `json:"level-num"`
	SrcDefaultGlid string                                                               `json:"src-default-glid"`
	GlidAction     string                                                               `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate `json:"zone-template"`
	Uuid           string                                                               `json:"uuid"`
	UserTag        string                                                               `json:"user-tag"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortRangeVirtualhosts) GetId() string {
	return "1"
}

func (p *DdosDstZonePortRangeVirtualhosts) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port-range/" + p.Inst.PortRangeStart + "+" + p.Inst.PortRangeEnd + "+" + p.Inst.Protocol + "/virtualhosts"
}

func (p *DdosDstZonePortRangeVirtualhosts) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhosts::Post")
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

func (p *DdosDstZonePortRangeVirtualhosts) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhosts::Get")
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
func (p *DdosDstZonePortRangeVirtualhosts) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhosts::Put")
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

func (p *DdosDstZonePortRangeVirtualhosts) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRangeVirtualhosts::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
