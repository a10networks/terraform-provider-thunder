

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwTopKRulesOper struct {
    
    Oper FwTopKRulesOperOper `json:"oper"`

}
type DataFwTopKRulesOper struct {
    DtFwTopKRulesOper FwTopKRulesOper `json:"top-k-rules"`
}


type FwTopKRulesOperOper struct {
    MetricTopkList []FwTopKRulesOperOperMetricTopkList `json:"metric-topk-list"`
}


type FwTopKRulesOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []FwTopKRulesOperOperMetricTopkListTopkList `json:"topk-list"`
}


type FwTopKRulesOperOperMetricTopkListTopkList struct {
    RuleName string `json:"rule-name"`
    MetricValue int `json:"metric-value"`
}

func (p *FwTopKRulesOper) GetId() string{
    return "1"
}

func (p *FwTopKRulesOper) getPath() string{
    return "fw/top-k-rules/oper"
}

func (p *FwTopKRulesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwTopKRulesOper,error) {
logger.Println("FwTopKRulesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwTopKRulesOper
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
