

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryIpProtoIpFilteringPolicyOperOper struct {
    
    Oper DdosDstEntryIpProtoIpFilteringPolicyOperOperOper `json:"oper"`
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryIpProtoIpFilteringPolicyOperOper struct {
    DtDdosDstEntryIpProtoIpFilteringPolicyOperOper DdosDstEntryIpProtoIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryIpProtoIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstEntryIpProtoIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryIpProtoIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/ip-proto/" +p.PortNum + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstEntryIpProtoIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryIpProtoIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstEntryIpProtoIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryIpProtoIpFilteringPolicyOperOper
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
