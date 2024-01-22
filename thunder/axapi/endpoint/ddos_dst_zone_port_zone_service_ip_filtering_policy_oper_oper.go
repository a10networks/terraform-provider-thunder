

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceIpFilteringPolicyOperOper struct {
    DtDdosDstZonePortZoneServiceIpFilteringPolicyOperOper DdosDstZonePortZoneServiceIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZonePortZoneServiceIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZonePortZoneServiceIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceIpFilteringPolicyOperOper
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
