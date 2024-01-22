

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitoredEntitySessionsOper struct {
    
    Oper VisibilityMonitoredEntitySessionsOperOper `json:"oper"`

}
type DataVisibilityMonitoredEntitySessionsOper struct {
    DtVisibilityMonitoredEntitySessionsOper VisibilityMonitoredEntitySessionsOper `json:"sessions"`
}


type VisibilityMonitoredEntitySessionsOperOper struct {
    MonEntityList []VisibilityMonitoredEntitySessionsOperOperMonEntityList `json:"mon-entity-list"`
}


type VisibilityMonitoredEntitySessionsOperOperMonEntityList struct {
    EntityKey string `json:"entity-key"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    SessionList []VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList `json:"session-list"`
    SecEntityList []VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList `json:"sec-entity-list"`
}


type VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList struct {
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


type VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList struct {
    EntityKey string `json:"entity-key"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4Proto string `json:"l4-proto"`
    L4Port int `json:"l4-port"`
    SessionList []VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList `json:"session-list"`
}


type VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList struct {
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

func (p *VisibilityMonitoredEntitySessionsOper) GetId() string{
    return "1"
}

func (p *VisibilityMonitoredEntitySessionsOper) getPath() string{
    return "visibility/monitored-entity/sessions/oper"
}

func (p *VisibilityMonitoredEntitySessionsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonitoredEntitySessionsOper,error) {
logger.Println("VisibilityMonitoredEntitySessionsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonitoredEntitySessionsOper
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
