

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangeIpFilteringPolicyOperOper struct {
    
    Oper DdosDstEntryPortRangeIpFilteringPolicyOperOperOper `json:"oper"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

}
type DataDdosDstEntryPortRangeIpFilteringPolicyOperOper struct {
    DtDdosDstEntryPortRangeIpFilteringPolicyOperOper DdosDstEntryPortRangeIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryPortRangeIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstEntryPortRangeIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryPortRangeIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstEntryPortRangeIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangeIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstEntryPortRangeIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangeIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstEntryPortRangeIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangeIpFilteringPolicyOperOper
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
