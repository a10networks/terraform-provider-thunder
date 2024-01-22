

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6IcmpStats struct {
    
    Stats Cgnv6IcmpStatsStats `json:"stats"`

}
type DataCgnv6IcmpStats struct {
    DtCgnv6IcmpStats Cgnv6IcmpStats `json:"icmp"`
}


type Cgnv6IcmpStatsStats struct {
    IcmpUnknownType int `json:"icmp-unknown-type"`
    IcmpNoPortInfo int `json:"icmp-no-port-info"`
    IcmpNoSessionDrop int `json:"icmp-no-session-drop"`
    Icmpv6UnknownType int `json:"icmpv6-unknown-type"`
    Icmpv6NoPortInfo int `json:"icmpv6-no-port-info"`
    Icmpv6NoSessionDrop int `json:"icmpv6-no-session-drop"`
    IcmpToIcmp int `json:"icmp-to-icmp"`
    IcmpToIcmpv6 int `json:"icmp-to-icmpv6"`
    Icmpv6ToIcmp int `json:"icmpv6-to-icmp"`
    Icmpv6ToIcmpv6 int `json:"icmpv6-to-icmpv6"`
    IcmpBadType int `json:"icmp-bad-type"`
    Icmpv6BadType int `json:"icmpv6-bad-type"`
    KnownDrop64 int `json:"known-drop64"`
    UnknownDrop64 int `json:"unknown-drop64"`
    MidpointHop64 int `json:"midpoint-hop64"`
    KnownDrop46 int `json:"known-drop46"`
    UnknownDrop46 int `json:"unknown-drop46"`
    NoPrefixForIpv446 int `json:"no-prefix-for-ipv446"`
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
}

func (p *Cgnv6IcmpStats) GetId() string{
    return "1"
}

func (p *Cgnv6IcmpStats) getPath() string{
    return "cgnv6/icmp/stats"
}

func (p *Cgnv6IcmpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6IcmpStats,error) {
logger.Println("Cgnv6IcmpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6IcmpStats
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
