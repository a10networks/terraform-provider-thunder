package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper struct {
	ClassListName string `json:"class-list-name"`

	Oper DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOperOper `json:"oper"`

	PortOther string

	Protocol string

	SrcBasedPolicyName string

	ZoneName string
}
type DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper struct {
	DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper `json:"policy-class-list"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOperOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port/zone-service-other/" + p.PortOther + "+" + p.Protocol + "/src-based-policy/" + p.SrcBasedPolicyName + "/policy-class-list/" + p.ClassListName + "/oper"
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper, error) {
	logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
