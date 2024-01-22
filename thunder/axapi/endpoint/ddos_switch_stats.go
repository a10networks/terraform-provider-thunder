

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSwitchStats struct {
    
    Stats DdosSwitchStatsStats `json:"stats"`

}
type DataDdosSwitchStats struct {
    DtDdosSwitchStats DdosSwitchStats `json:"switch"`
}


type DdosSwitchStatsStats struct {
    Ip_rcvd int `json:"ip_rcvd"`
    Ip_sent int `json:"ip_sent"`
    Ipv6_rcvd int `json:"ipv6_rcvd"`
    Ipv6_sent int `json:"ipv6_sent"`
    In_stateless_pkt int `json:"in_stateless_pkt"`
    Noroute int `json:"noroute"`
    Pkt_not_for_ddos int `json:"pkt_not_for_ddos"`
    Ingress_bytes int `json:"ingress_bytes"`
    Egress_bytes int `json:"egress_bytes"`
    Ingress_packets int `json:"ingress_packets"`
    Egress_packets int `json:"egress_packets"`
    Src_ip_bypass int `json:"src_ip_bypass"`
    Dst_ip_bypass int `json:"dst_ip_bypass"`
    Dst_blackhole_inject int `json:"dst_blackhole_inject"`
    Dst_blackhole_withdraw int `json:"dst_blackhole_withdraw"`
    Mpls_rcvd int `json:"mpls_rcvd"`
    Mpls_drop int `json:"mpls_drop"`
    Mpls_malformed int `json:"mpls_malformed"`
    Jumbo_frag_drop_by_filter int `json:"jumbo_frag_drop_by_filter"`
    Jumbo_frag_drop_before_slb int `json:"jumbo_frag_drop_before_slb"`
    Jumbo_outgoing_mtu_exceed_drop int `json:"jumbo_outgoing_mtu_exceed_drop"`
}

func (p *DdosSwitchStats) GetId() string{
    return "1"
}

func (p *DdosSwitchStats) getPath() string{
    return "ddos/switch/stats"
}

func (p *DdosSwitchStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSwitchStats,error) {
logger.Println("DdosSwitchStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosSwitchStats
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
