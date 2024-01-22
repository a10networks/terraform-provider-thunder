

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntitySecondaryMonTopkSourcesOper struct {
    
    Oper VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper `json:"oper"`

}
type DataVisibilityMonitoredEntitySecondaryMonTopkSourcesOper struct {
    DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper VisibilityMonitoredEntitySecondaryMonTopkSourcesOper `json:"sources"`
}


type VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkSourcesOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkSourcesOper) getPath() string{
    return "visibility/monitored-entity/secondary/mon-topk/sources/oper"
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntitySecondaryMonTopkSourcesOper,error) {
logger.Println("VisibilityMonitoredEntitySecondaryMonTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntitySecondaryMonTopkSourcesOper
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
