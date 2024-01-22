

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityDetailOper struct {
    
    Debug VisibilityMonitoredEntityDetailOperDebug `json:"debug"`
    Oper VisibilityMonitoredEntityDetailOperOper `json:"oper"`

}
type DataVisibilityMonitoredEntityDetailOper struct {
    DtVisibilityMonitoredEntityDetailOper VisibilityMonitoredEntityDetailOper `json:"detail"`
}


type VisibilityMonitoredEntityDetailOperDebug struct {
    Oper VisibilityMonitoredEntityDetailOperDebugOper `json:"oper"`
}


type VisibilityMonitoredEntityDetailOperDebugOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityDetailOperDebugOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityDetailOperDebugOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList `json:"entity-metric-list"`
    SecEntityList []VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList `json:"entity-metric-list"`
}


type VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityDetailOperOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityDetailOperOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityDetailOperOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList `json:"entity-metric-list"`
    SecEntityList []VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Threshold string `json:"threshold"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList `json:"entity-metric-list"`
}


type VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Threshold string `json:"threshold"`
    Anomaly string `json:"anomaly"`
}

func (p *VisibilityMonitoredEntityDetailOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityDetailOper) getPath() string{
    return "visibility/monitored-entity/detail/oper"
}

func (p *VisibilityMonitoredEntityDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntityDetailOper,error) {
logger.Println("VisibilityMonitoredEntityDetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntityDetailOper
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
