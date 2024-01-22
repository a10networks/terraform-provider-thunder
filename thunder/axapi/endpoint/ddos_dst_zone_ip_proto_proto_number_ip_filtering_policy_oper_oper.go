

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOperOper `json:"oper"`
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper struct {
    DtDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyOperOper
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
