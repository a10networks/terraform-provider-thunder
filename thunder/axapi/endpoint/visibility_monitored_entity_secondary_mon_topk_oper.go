

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntitySecondaryMonTopkOper struct {
    
    Oper VisibilityMonitoredEntitySecondaryMonTopkOperOper `json:"oper"`
    Sources VisibilityMonitoredEntitySecondaryMonTopkOperSources `json:"sources"`

}
type DataVisibilityMonitoredEntitySecondaryMonTopkOper struct {
    DtVisibilityMonitoredEntitySecondaryMonTopkOper VisibilityMonitoredEntitySecondaryMonTopkOper `json:"mon-topk"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperSources struct {
    Oper VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper `json:"oper"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkOper) getPath() string{
    return "visibility/monitored-entity/secondary/mon-topk/oper"
}

func (p *VisibilityMonitoredEntitySecondaryMonTopkOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntitySecondaryMonTopkOper,error) {
logger.Println("VisibilityMonitoredEntitySecondaryMonTopkOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntitySecondaryMonTopkOper
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
