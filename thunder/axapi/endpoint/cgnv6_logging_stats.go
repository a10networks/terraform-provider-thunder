

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LoggingStats struct {
    
    Stats Cgnv6LoggingStatsStats `json:"stats"`

}
type DataCgnv6LoggingStats struct {
    DtCgnv6LoggingStats Cgnv6LoggingStats `json:"logging"`
}


type Cgnv6LoggingStatsStats struct {
    TcpSessionCreated int `json:"tcp-session-created"`
    TcpSessionDeleted int `json:"tcp-session-deleted"`
    TcpPortAllocated int `json:"tcp-port-allocated"`
    TcpPortFreed int `json:"tcp-port-freed"`
    TcpPortBatchAllocated int `json:"tcp-port-batch-allocated"`
    TcpPortBatchFreed int `json:"tcp-port-batch-freed"`
    UdpSessionCreated int `json:"udp-session-created"`
    UdpSessionDeleted int `json:"udp-session-deleted"`
    UdpPortAllocated int `json:"udp-port-allocated"`
    UdpPortFreed int `json:"udp-port-freed"`
    UdpPortBatchAllocated int `json:"udp-port-batch-allocated"`
    UdpPortBatchFreed int `json:"udp-port-batch-freed"`
    IcmpSessionCreated int `json:"icmp-session-created"`
    IcmpSessionDeleted int `json:"icmp-session-deleted"`
    IcmpResourceAllocated int `json:"icmp-resource-allocated"`
    IcmpResourceFreed int `json:"icmp-resource-freed"`
    Icmpv6SessionCreated int `json:"icmpv6-session-created"`
    Icmpv6SessionDeleted int `json:"icmpv6-session-deleted"`
    Icmpv6ResourceAllocated int `json:"icmpv6-resource-allocated"`
    Icmpv6ResourceFreed int `json:"icmpv6-resource-freed"`
    GreSessionCreated int `json:"gre-session-created"`
    GreSessionDeleted int `json:"gre-session-deleted"`
    GreResourceAllocated int `json:"gre-resource-allocated"`
    GreResourceFreed int `json:"gre-resource-freed"`
    EspSessionCreated int `json:"esp-session-created"`
    EspSessionDeleted int `json:"esp-session-deleted"`
    EspResourceAllocated int `json:"esp-resource-allocated"`
    EspResourceFreed int `json:"esp-resource-freed"`
    FixedNatUserPorts int `json:"fixed-nat-user-ports"`
    FixedNatDisableConfigLogged int `json:"fixed-nat-disable-config-logged"`
    FixedNatDisableConfigLogsSent int `json:"fixed-nat-disable-config-logs-sent"`
    FixedNatPeriodicConfigLogsSent int `json:"fixed-nat-periodic-config-logs-sent"`
    FixedNatPeriodicConfigLogged int `json:"fixed-nat-periodic-config-logged"`
    FixedNatInterimUpdated int `json:"fixed-nat-interim-updated"`
    EnhancedUserLog int `json:"enhanced-user-log"`
    LogSent int `json:"log-sent"`
    LogDropped int `json:"log-dropped"`
    ConnTcpEstablished int `json:"conn-tcp-established"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    TcpPortOverloadingAllocated int `json:"tcp-port-overloading-allocated"`
    TcpPortOverloadingFreed int `json:"tcp-port-overloading-freed"`
    UdpPortOverloadingAllocated int `json:"udp-port-overloading-allocated"`
    UdpPortOverloadingFreed int `json:"udp-port-overloading-freed"`
    HttpRequestLogged int `json:"http-request-logged"`
    ReducedLogsByDestination int `json:"reduced-logs-by-destination"`
    InterimUpdateScheduled int `json:"interim-update-scheduled"`
    TcpPortBatchInterimUpdated int `json:"tcp-port-batch-interim-updated"`
    UdpPortBatchInterimUpdated int `json:"udp-port-batch-interim-updated"`
    IddosL3EntryCreate int `json:"iddos-l3-entry-create"`
    IddosL3EntryDelete int `json:"iddos-l3-entry-delete"`
    IddosL4EntryCreate int `json:"iddos-l4-entry-create"`
    IddosL4EntryDelete int `json:"iddos-l4-entry-delete"`
}

func (p *Cgnv6LoggingStats) GetId() string{
    return "1"
}

func (p *Cgnv6LoggingStats) getPath() string{
    return "cgnv6/logging/stats"
}

func (p *Cgnv6LoggingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LoggingStats,error) {
logger.Println("Cgnv6LoggingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LoggingStats
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
