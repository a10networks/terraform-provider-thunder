

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypeIpFilteringPolicyOperOper struct {
    
    Oper DdosDstEntryL4TypeIpFilteringPolicyOperOperOper `json:"oper"`
    Protocol string 
    DstEntryName string 

}
type DataDdosDstEntryL4TypeIpFilteringPolicyOperOper struct {
    DtDdosDstEntryL4TypeIpFilteringPolicyOperOper DdosDstEntryL4TypeIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryL4TypeIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstEntryL4TypeIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypeIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/l4-type/" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstEntryL4TypeIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryL4TypeIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstEntryL4TypeIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryL4TypeIpFilteringPolicyOperOper
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
