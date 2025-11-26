package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper struct {
	ClassListName string `json:"class-list-name"`

	Oper DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper `json:"oper"`

	Protocol string

	SrcBasedPolicyName string

	ZoneName string

	PortNum string
}
type DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper struct {
	DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper `json:"policy-class-list"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper struct {
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

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port/zone-service/" + p.PortNum + "+" + p.Protocol + "/src-based-policy/" + p.SrcBasedPolicyName + "/policy-class-list/" + p.ClassListName + "/oper"
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper, error) {
	logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper
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
