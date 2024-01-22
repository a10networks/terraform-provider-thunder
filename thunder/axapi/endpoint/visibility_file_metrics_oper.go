

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFileMetricsOper struct {
    
    Oper VisibilityFileMetricsOperOper `json:"oper"`

}
type DataVisibilityFileMetricsOper struct {
    DtVisibilityFileMetricsOper VisibilityFileMetricsOper `json:"metrics"`
}


type VisibilityFileMetricsOperOper struct {
    MonitorType string `json:"monitor-type"`
    PriType string `json:"pri-type"`
    PriIpv4Addr string `json:"pri-ipv4-addr"`
    PriIpv6Addr string `json:"pri-ipv6-addr"`
    PriL4Proto int `json:"pri-l4-proto"`
    PriL4Port int `json:"pri-l4-port"`
    SecType string `json:"sec-type"`
    SecIpv4Addr string `json:"sec-ipv4-addr"`
    SecIpv6Addr string `json:"sec-ipv6-addr"`
    SecL4Proto int `json:"sec-l4-proto"`
    SecL4Port int `json:"sec-l4-port"`
    FileName string `json:"file-name"`
    ProcMetricList []VisibilityFileMetricsOperOperProcMetricList `json:"proc-metric-list"`
}


type VisibilityFileMetricsOperOperProcMetricList struct {
    MetricName string `json:"metric-name"`
    MetricAttrList []VisibilityFileMetricsOperOperProcMetricListMetricAttrList `json:"metric-attr-list"`
}


type VisibilityFileMetricsOperOperProcMetricListMetricAttrList struct {
    MetricAttrName string `json:"metric-attr-name"`
    MetricAttrValue string `json:"metric-attr-value"`
}

func (p *VisibilityFileMetricsOper) GetId() string{
    return "1"
}

func (p *VisibilityFileMetricsOper) getPath() string{
    return "visibility/file/metrics/oper"
}

func (p *VisibilityFileMetricsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFileMetricsOper,error) {
logger.Println("VisibilityFileMetricsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFileMetricsOper
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
