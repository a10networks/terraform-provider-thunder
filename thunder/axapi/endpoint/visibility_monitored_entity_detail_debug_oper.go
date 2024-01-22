

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntityDetailDebugOper struct {
    
    Oper VisibilityMonitoredEntityDetailDebugOperOper `json:"oper"`

}
type DataVisibilityMonitoredEntityDetailDebugOper struct {
    DtVisibilityMonitoredEntityDetailDebugOper VisibilityMonitoredEntityDetailDebugOper `json:"debug"`
}


type VisibilityMonitoredEntityDetailDebugOperOper struct {
    PrimaryKeys int `json:"primary-keys"`
    AllKeys int `json:"all-keys"`
    MonEntityList []VisibilityMonitoredEntityDetailDebugOperOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntityDetailDebugOperOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList `json:"entity-metric-list"`
    SecEntityList []VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}


type VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    Mode string `json:"mode"`
    HaState string `json:"ha-state"`
    EntityMetricList []VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList `json:"entity-metric-list"`
}


type VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList struct {
    MetricName string `json:"metric-name"`
    Current string `json:"current"`
    Min string `json:"min"`
    Max string `json:"max"`
    Mean string `json:"mean"`
    Threshold string `json:"threshold"`
    StdDev string `json:"std-dev"`
    Anomaly string `json:"anomaly"`
}

func (p *VisibilityMonitoredEntityDetailDebugOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntityDetailDebugOper) getPath() string{
    return "visibility/monitored-entity/detail/debug/oper"
}

func (p *VisibilityMonitoredEntityDetailDebugOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntityDetailDebugOper,error) {
logger.Println("VisibilityMonitoredEntityDetailDebugOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntityDetailDebugOper
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
