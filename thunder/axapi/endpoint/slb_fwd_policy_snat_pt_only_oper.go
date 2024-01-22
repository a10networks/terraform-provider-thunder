

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFwdPolicySnatPtOnlyOper struct {
    
    Oper SlbFwdPolicySnatPtOnlyOperOper `json:"oper"`

}
type DataSlbFwdPolicySnatPtOnlyOper struct {
    DtSlbFwdPolicySnatPtOnlyOper SlbFwdPolicySnatPtOnlyOper `json:"fwd-policy-snat-pt-only"`
}


type SlbFwdPolicySnatPtOnlyOperOper struct {
    GroupList []SlbFwdPolicySnatPtOnlyOperOperGroupList `json:"group-list"`
    DetailInfo int `json:"detail-info"`
}


type SlbFwdPolicySnatPtOnlyOperOperGroupList struct {
    Group string `json:"group"`
    PortUsage int `json:"port-usage"`
    TotalUsed int `json:"total-used"`
    TotalFreed int `json:"total-freed"`
    Failed int `json:"failed"`
}

func (p *SlbFwdPolicySnatPtOnlyOper) GetId() string{
    return "1"
}

func (p *SlbFwdPolicySnatPtOnlyOper) getPath() string{
    return "slb/fwd-policy-snat-pt-only/oper"
}

func (p *SlbFwdPolicySnatPtOnlyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFwdPolicySnatPtOnlyOper,error) {
logger.Println("SlbFwdPolicySnatPtOnlyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFwdPolicySnatPtOnlyOper
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
