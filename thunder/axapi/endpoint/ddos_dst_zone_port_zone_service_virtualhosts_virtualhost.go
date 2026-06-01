package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortZoneServiceVirtualhostsVirtualhost struct {
	Inst struct {
		Deny int `json:"deny"`

		GlidCfg DdosDstZonePortZoneServiceVirtualhostsVirtualhostGlidCfg `json:"glid-cfg"`

		LevelList []DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelList `json:"level-list"`

		Servername []DdosDstZonePortZoneServiceVirtualhostsVirtualhostServername `json:"servername"`

		ServernameList string `json:"servername-list"`

		ServernameMatchAny int `json:"servername-match-any"`

		ServernameNoSni int `json:"servername-no-sni"`

		SourceTracking string `json:"source-tracking" dval:"follow"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Vhost string `json:"vhost"`

		ZoneName string

		PortNum string

		Protocol string
	} `json:"virtualhost"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelList struct {
	LevelNum       string                                                                 `json:"level-num"`
	SrcDefaultGlid string                                                                 `json:"src-default-glid"`
	GlidAction     string                                                                 `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelListZoneTemplate `json:"zone-template"`
	Uuid           string                                                                 `json:"uuid"`
	UserTag        string                                                                 `json:"user-tag"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelListZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostServername struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) GetId() string {
	return url.QueryEscape(p.Inst.Vhost)
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port/zone-service/" + p.Inst.PortNum + "+" + p.Inst.Protocol + "/virtualhosts/virtualhost"
}

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhost::Post")
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

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhost::Get")
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
func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhost::Put")
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

func (p *DdosDstZonePortZoneServiceVirtualhostsVirtualhost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceVirtualhostsVirtualhost::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
