

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamResourceUsageOper struct {
    
    Oper AamResourceUsageOperOper `json:"oper"`

}
type DataAamResourceUsageOper struct {
    DtAamResourceUsageOper AamResourceUsageOper `json:"resource-usage"`
}


type AamResourceUsageOperOper struct {
    IdpLimitCurrent int `json:"idp-limit-current"`
    IdpLimitDefault int `json:"idp-limit-default"`
    IdpLimitMinimum int `json:"idp-limit-minimum"`
    IdpLimitMaximum int `json:"idp-limit-maximum"`
}

func (p *AamResourceUsageOper) GetId() string{
    return "1"
}

func (p *AamResourceUsageOper) getPath() string{
    return "aam/resource-usage/oper"
}

func (p *AamResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamResourceUsageOper,error) {
logger.Println("AamResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamResourceUsageOper
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
