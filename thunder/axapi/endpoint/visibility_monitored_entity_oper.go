

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityOper struct {
    
    Detail VisibilityMonitoredEntityOperDetail `json:"detail"`
    MonTopk VisibilityMonitoredEntityOperMonTopk `json:"mon-topk"`
    Oper VisibilityMonitoredEntityOperOper `json:"oper"`
    Secondary VisibilityMonitoredEntityOperSecondary `json:"secondary"`
    Sessions VisibilityMonitoredEntityOperSessions `json:"sessions"`

}
type DataVisibilityMonitoredEntityOper struct {
    DtVisibilityMonitoredEntityOper VisibilityMonitoredEntityOper `json:"monitored-entity"`
}


type VisibilityMonitoredEntityOperDetail struct {
    Oper VisibilityMonitoredEntityOperDetailOper `json:"oper"`
    Debug VisibilityMonitoredEntityOperDetailDebug `json:"debug"`
}


type VisibilityMonitoredEntityOperDetailOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityOperDetailOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityOperDetailOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList `json:"entity-metric-list"`
    SecEntityList []VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Threshold string `json:"threshold"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList `json:"entity-metric-list"`
}


type VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Threshold string `json:"threshold"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityOperDetailDebug struct {
    Oper VisibilityMonitoredEntityOperDetailDebugOper `json:"oper"`
}


type VisibilityMonitoredEntityOperDetailDebugOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityOperDetailDebugOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityOperDetailDebugOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList `json:"entity-metric-list"`
    SecEntityList []VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList `json:"entity-metric-list"`
}


type VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityOperMonTopk struct {
    Oper VisibilityMonitoredEntityOperMonTopkOper `json:"oper"`
    Sources VisibilityMonitoredEntityOperMonTopkSources `json:"sources"`
}


type VisibilityMonitoredEntityOperMonTopkOper struct {
    MetricTopkList []VisibilityMonitoredEntityOperMonTopkOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityOperMonTopkOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntityOperMonTopkSources struct {
    Oper VisibilityMonitoredEntityOperMonTopkSourcesOper `json:"oper"`
}


type VisibilityMonitoredEntityOperMonTopkSourcesOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntityOperOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityOperOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityOperOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    SecEntityList []VisibilityMonitoredEntityOperOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityOperOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
}


type VisibilityMonitoredEntityOperSecondary struct {
    Oper VisibilityMonitoredEntityOperSecondaryOper `json:"oper"`
    MonTopk VisibilityMonitoredEntityOperSecondaryMonTopk `json:"mon-topk"`
}


type VisibilityMonitoredEntityOperSecondaryOper struct {
}


type VisibilityMonitoredEntityOperSecondaryMonTopk struct {
    Oper VisibilityMonitoredEntityOperSecondaryMonTopkOper `json:"oper"`
    Sources VisibilityMonitoredEntityOperSecondaryMonTopkSources `json:"sources"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkSources struct {
    Oper VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper `json:"oper"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    MetricTopkList []VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList `json:"metric-topk-list"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList struct {
    MetricName string `json:"metric-name"`
    TopkList []VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList `json:"topk-list"`
}


type VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList struct {
    IpAddr string `json:"ip-addr"`
    MetricValue string `json:"metric-value"`
}


type VisibilityMonitoredEntityOperSessions struct {
    Oper VisibilityMonitoredEntityOperSessionsOper `json:"oper"`
}


type VisibilityMonitoredEntityOperSessionsOper struct {
    MonEntityList []VisibilityMonitoredEntityOperSessionsOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityOperSessionsOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    SessionList []VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList `json:"session-list"`
    SecEntityList []VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList struct {
    Proto string `json:"proto"`
    FwdSrcIp string `json:"fwd-src-ip"`
    FwdSrcPort int `json:"fwd-src-port"`
    FwdDstIp string `json:"fwd-dst-ip"`
    FwdDstPort int `json:"fwd-dst-port"`
    RevSrcIp string `json:"rev-src-ip"`
    RevSrcPort int `json:"rev-src-port"`
    RevDstIp string `json:"rev-dst-ip"`
    RevDstPort int `json:"rev-dst-port"`
}


type VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    SessionList []VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList `json:"session-list"`
}


type VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList struct {
    Proto string `json:"proto"`
    FwdSrcIp string `json:"fwd-src-ip"`
    FwdSrcPort int `json:"fwd-src-port"`
    FwdDstIp string `json:"fwd-dst-ip"`
    FwdDstPort int `json:"fwd-dst-port"`
    RevSrcIp string `json:"rev-src-ip"`
    RevSrcPort int `json:"rev-src-port"`
    RevDstIp string `json:"rev-dst-ip"`
    RevDstPort int `json:"rev-dst-port"`
}

func (p *VisibilityMonitoredEntityOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityOper) getPath() string{
    return "visibility/monitored-entity/oper"
}

func (p *VisibilityMonitoredEntityOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntityOper,error) {
logger.Println("VisibilityMonitoredEntityOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntityOper
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
