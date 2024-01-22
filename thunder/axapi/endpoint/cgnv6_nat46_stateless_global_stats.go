

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat46StatelessGlobalStats struct {
    
    Stats Cgnv6Nat46StatelessGlobalStatsStats `json:"stats"`

}
type DataCgnv6Nat46StatelessGlobalStats struct {
    DtCgnv6Nat46StatelessGlobalStats Cgnv6Nat46StatelessGlobalStats `json:"global"`
}


type Cgnv6Nat46StatelessGlobalStatsStats struct {
    Outbound_ipv4_received int `json:"outbound_ipv4_received"`
    Outbound_ipv4_drop int `json:"outbound_ipv4_drop"`
    Outbound_ipv4_fragment_received int `json:"outbound_ipv4_fragment_received"`
    Outbound_ipv6_unreachable int `json:"outbound_ipv6_unreachable"`
    Outbound_ipv6_fragmented int `json:"outbound_ipv6_fragmented"`
    Inbound_ipv6_received int `json:"inbound_ipv6_received"`
    Inbound_ipv6_drop int `json:"inbound_ipv6_drop"`
    Inbound_ipv6_fragment_received int `json:"inbound_ipv6_fragment_received"`
    Inbound_ipv4_unreachable int `json:"inbound_ipv4_unreachable"`
    Inbound_ipv4_fragmented int `json:"inbound_ipv4_fragmented"`
    Packet_too_big int `json:"packet_too_big"`
    Fragment_error int `json:"fragment_error"`
    Icmpv6_to_icmp int `json:"icmpv6_to_icmp"`
    Icmpv6_to_icmp_error int `json:"icmpv6_to_icmp_error"`
    Icmp_to_icmpv6 int `json:"icmp_to_icmpv6"`
    Icmp_to_icmpv6_error int `json:"icmp_to_icmpv6_error"`
    Ha_standby int `json:"ha_standby"`
    Other_error int `json:"other_error"`
    Conn_count int `json:"conn_count"`
}

func (p *Cgnv6Nat46StatelessGlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6Nat46StatelessGlobalStats) getPath() string{
    return "cgnv6/nat46-stateless/global/stats"
}

func (p *Cgnv6Nat46StatelessGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat46StatelessGlobalStats,error) {
logger.Println("Cgnv6Nat46StatelessGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat46StatelessGlobalStats
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
