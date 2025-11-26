package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZoneWebGui struct {
	Inst struct {
		ActivatedAfterLearning int `json:"activated-after-learning"`

		CreateTime string `json:"create-time"`

		Learning DdosDstZoneWebGuiLearning270 `json:"learning"`

		ModifyTime string `json:"modify-time"`

		Protection DdosDstZoneWebGuiProtection271 `json:"protection"`

		Sensitivity string `json:"sensitivity" dval:"3"`

		Status string `json:"status" dval:"newly"`

		Uuid string `json:"uuid"`

		ZoneName string
	} `json:"web-gui"`
}

type DdosDstZoneWebGuiLearning270 struct {
	Duration     string `json:"duration" dval:"6hour"`
	StartingTime string `json:"starting-time"`
	Uuid         string `json:"uuid"`
}

type DdosDstZoneWebGuiProtection271 struct {
	Port          DdosDstZoneWebGuiProtectionPort272            `json:"port"`
	IpProto       DdosDstZoneWebGuiProtectionIpProto275         `json:"ip-proto"`
	PortRangeList []DdosDstZoneWebGuiProtectionPortRangeList277 `json:"port-range-list"`
}

type DdosDstZoneWebGuiProtectionPort272 struct {
	ZoneServiceList      []DdosDstZoneWebGuiProtectionPortZoneServiceList273      `json:"zone-service-list"`
	ZoneServiceOtherList []DdosDstZoneWebGuiProtectionPortZoneServiceOtherList274 `json:"zone-service-other-list"`
}

type DdosDstZoneWebGuiProtectionPortZoneServiceList273 struct {
	PortNum  int    `json:"port-num"`
	Protocol string `json:"protocol"`
	Pbe      string `json:"pbe"`
	Uuid     string `json:"uuid"`
}

type DdosDstZoneWebGuiProtectionPortZoneServiceOtherList274 struct {
	PortOther string `json:"port-other"`
	Protocol  string `json:"protocol"`
	Pbe       string `json:"pbe"`
	Uuid      string `json:"uuid"`
}

type DdosDstZoneWebGuiProtectionIpProto275 struct {
	ProtoNameList []DdosDstZoneWebGuiProtectionIpProtoProtoNameList276 `json:"proto-name-list"`
}

type DdosDstZoneWebGuiProtectionIpProtoProtoNameList276 struct {
	Protocol string `json:"protocol"`
	Pbe      string `json:"pbe"`
	Uuid     string `json:"uuid"`
	UserTag  string `json:"user-tag"`
}

type DdosDstZoneWebGuiProtectionPortRangeList277 struct {
	PortRangeStart int    `json:"port-range-start"`
	PortRangeEnd   int    `json:"port-range-end"`
	Protocol       string `json:"protocol"`
	Pbe            string `json:"pbe"`
	Uuid           string `json:"uuid"`
	UserTag        string `json:"user-tag"`
}

func (p *DdosDstZoneWebGui) GetId() string {
	return "1"
}

func (p *DdosDstZoneWebGui) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/web-gui"
}

func (p *DdosDstZoneWebGui) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneWebGui::Post")
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

func (p *DdosDstZoneWebGui) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneWebGui::Get")
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
func (p *DdosDstZoneWebGui) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneWebGui::Put")
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

func (p *DdosDstZoneWebGui) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneWebGui::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
