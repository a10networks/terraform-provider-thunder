

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityMonTopkOper struct {
    
    Oper VisibilityMonitoredEntityMonTopkOperOper `json:"oper"`
    Sources VisibilityMonitoredEntityMonTopkOperSources `json:"sources"`

}
type DataVisibilityMonitoredEntityMonTopkOper struct {
    DtVisibilityMonitoredEntityMonTopkOper VisibilityMonitoredEntityMonTopkOper `json:"mon-topk"`
}


type VisibilityMonitoredEntityMonTopkOperOper struct {
    MetricTopkList []VisibilityMonitoredEntityMonTopkOperOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityMonTopkOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntityMonTopkOperSources struct {
    Oper VisibilityMonitoredEntityMonTopkOperSourcesOper `json:"oper"`
}


type VisibilityMonitoredEntityMonTopkOperSourcesOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}

func (p *VisibilityMonitoredEntityMonTopkOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityMonTopkOper) getPath() string{
    return "visibility/monitored-entity/mon-topk/oper"
}

func (p *VisibilityMonitoredEntityMonTopkOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntityMonTopkOper,error) {
logger.Println("VisibilityMonitoredEntityMonTopkOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntityMonTopkOper
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
