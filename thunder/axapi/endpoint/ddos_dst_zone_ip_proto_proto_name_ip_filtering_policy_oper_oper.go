

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOperOper `json:"oper"`
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper struct {
    DtDdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyOperOper
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
