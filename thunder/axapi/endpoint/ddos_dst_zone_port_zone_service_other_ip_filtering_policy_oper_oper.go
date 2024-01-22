

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper struct {
    DtDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/" +p.PortOther + "+" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper
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
