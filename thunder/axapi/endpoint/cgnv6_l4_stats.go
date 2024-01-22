

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6L4Stats struct {
    
    Stats Cgnv6L4StatsStats `json:"stats"`

}
type DataCgnv6L4Stats struct {
    DtCgnv6L4Stats Cgnv6L4Stats `json:"l4"`
}


type Cgnv6L4StatsStats struct {
    NoFwdRoute int `json:"no-fwd-route"`
    NoRevRoute int `json:"no-rev-route"`
    OutOfSessionMemory int `json:"out-of-session-memory"`
    TcpRstSent int `json:"tcp-rst-sent"`
    IpipIcmpReplySent int `json:"ipip-icmp-reply-sent"`
    IcmpFilteredSent int `json:"icmp-filtered-sent"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    IcmpReplyNoSessionDrop int `json:"icmp-reply-no-session-drop"`
    IpipTruncated int `json:"ipip-truncated"`
    IpSrcInvalidUnicast int `json:"ip-src-invalid-unicast"`
    IpDstInvalidUnicast int `json:"ip-dst-invalid-unicast"`
    Ipv6SrcInvalidUnicast int `json:"ipv6-src-invalid-unicast"`
    Ipv6DstInvalidUnicast int `json:"ipv6-dst-invalid-unicast"`
    Rate_drop_reset_unkn int `json:"rate_drop_reset_unkn"`
}

func (p *Cgnv6L4Stats) GetId() string{
    return "1"
}

func (p *Cgnv6L4Stats) getPath() string{
    return "cgnv6/l4/stats"
}

func (p *Cgnv6L4Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6L4Stats,error) {
logger.Println("Cgnv6L4Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6L4Stats
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
