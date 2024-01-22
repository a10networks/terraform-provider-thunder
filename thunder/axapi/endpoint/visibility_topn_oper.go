

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnOper struct {
    
    Oper VisibilityTopnOperOper `json:"oper"`

}
type DataVisibilityTopnOper struct {
    DtVisibilityTopnOper VisibilityTopnOper `json:"topn"`
}


type VisibilityTopnOperOper struct {
    Class string `json:"class"`
    Metric string `json:"metric"`
    MemoryUsage int `json:"memory-usage"`
    TopnDuration string `json:"topn-duration"`
    MetricTopnList []VisibilityTopnOperOperMetricTopnList `json:"metric-topn-list"`
    TotalMemory string `json:"total-memory"`
    UsedMemory string `json:"used-memory"`
}


type VisibilityTopnOperOperMetricTopnList struct {
    ObjectName string `json:"object-name"`
    ObjectVal int `json:"object-val"`
}

func (p *VisibilityTopnOper) GetId() string{
    return "1"
}

func (p *VisibilityTopnOper) getPath() string{
    return "visibility/topn/oper"
}

func (p *VisibilityTopnOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityTopnOper,error) {
logger.Println("VisibilityTopnOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityTopnOper
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
