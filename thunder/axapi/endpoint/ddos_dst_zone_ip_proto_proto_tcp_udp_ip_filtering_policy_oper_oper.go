

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper `json:"oper"`
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper struct {
    DtDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-tcp-udp/" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
