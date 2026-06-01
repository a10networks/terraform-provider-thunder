package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper struct {
	ClassListName string `json:"class-list-name"`

	Oper DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOperOper `json:"oper"`

	ZoneName string

	ProtocolNum string

	SrcBasedPolicyName string
}
type DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper struct {
	DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper `json:"policy-class-list"`
}

type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOperOper struct {
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

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/ip-proto/proto-number/" + p.ProtocolNum + "/src-based-policy/" + p.SrcBasedPolicyName + "/policy-class-list/" + p.ClassListName + "/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper, error) {
	logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListOper
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
