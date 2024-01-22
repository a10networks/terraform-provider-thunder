

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonTopkSourcesOper struct {
    
    Oper VisibilityMonTopkSourcesOperOper `json:"oper"`

}
type DataVisibilityMonTopkSourcesOper struct {
    DtVisibilityMonTopkSourcesOper VisibilityMonTopkSourcesOper `json:"sources"`
}


type VisibilityMonTopkSourcesOperOper struct {
    MetricTopkList []VisibilityMonTopkSourcesOperOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonTopkSourcesOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonTopkSourcesOperOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonTopkSourcesOperOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}

func (p *VisibilityMonTopkSourcesOper) GetId() string{
    return "1"
}

func (p *VisibilityMonTopkSourcesOper) getPath() string{
    return "visibility/mon-topk/sources/oper"
}

func (p *VisibilityMonTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonTopkSourcesOper,error) {
logger.Println("VisibilityMonTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonTopkSourcesOper
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
