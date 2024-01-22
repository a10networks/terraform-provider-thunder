

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityMonTopkSourcesOper struct {
    
    Oper VisibilityMonitoredEntityMonTopkSourcesOperOper `json:"oper"`

}
type DataVisibilityMonitoredEntityMonTopkSourcesOper struct {
    DtVisibilityMonitoredEntityMonTopkSourcesOper VisibilityMonitoredEntityMonTopkSourcesOper `json:"sources"`
}


type VisibilityMonitoredEntityMonTopkSourcesOperOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}

func (p *VisibilityMonitoredEntityMonTopkSourcesOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityMonTopkSourcesOper) getPath() string{
    return "visibility/monitored-entity/mon-topk/sources/oper"
}

func (p *VisibilityMonitoredEntityMonTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntityMonTopkSourcesOper,error) {
logger.Println("VisibilityMonitoredEntityMonTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntityMonTopkSourcesOper
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
