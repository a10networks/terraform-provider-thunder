

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosBriefStats struct {
    
    Stats DdosBriefStatsStats `json:"stats"`

}
type DataDdosBriefStats struct {
    DtDdosBriefStats DdosBriefStats `json:"brief"`
}


type DdosBriefStatsStats struct {
    Ip_rcv int `json:"ip_rcv"`
    Ip_sent int `json:"ip_sent"`
    Ipv6_rcv int `json:"ipv6_rcv"`
    Ipv6_sent int `json:"ipv6_sent"`
    Out_no_route int `json:"out_no_route"`
    Not_for_ddos int `json:"not_for_ddos"`
    Instateless int `json:"instateless"`
    Intcp int `json:"intcp"`
    Inudp int `json:"inudp"`
    Inicmp int `json:"inicmp"`
    Inother int `json:"inother"`
    V4_sess_create int `json:"v4_sess_create"`
    V6_sess_create int `json:"v6_sess_create"`
    Tcp_sess_create int `json:"tcp_sess_create"`
    Udp_sess_create int `json:"udp_sess_create"`
    Sess_aged_out int `json:"sess_aged_out"`
    Tcp_total_drop int `json:"tcp_total_drop"`
    Tcp_dst_drop int `json:"tcp_dst_drop"`
    Tcp_src_drop int `json:"tcp_src_drop"`
    Tcp_src_dst_drop int `json:"tcp_src_dst_drop"`
    Udp_total_drop int `json:"udp_total_drop"`
    Udp_dst_drop int `json:"udp_dst_drop"`
    Udp_src_drop int `json:"udp_src_drop"`
    Udp_src_dst_drop int `json:"udp_src_dst_drop"`
    Icmp_total_drop int `json:"icmp_total_drop"`
    Icmp_dst_drop int `json:"icmp_dst_drop"`
    Icmp_src_drop int `json:"icmp_src_drop"`
    Icmp_src_dst_drop int `json:"icmp_src_dst_drop"`
    Other_total_drop int `json:"other_total_drop"`
    Other_dst_drop int `json:"other_dst_drop"`
    Other_src_drop int `json:"other_src_drop"`
    Other_src_dst_drop int `json:"other_src_dst_drop"`
    Frag_rcvd int `json:"frag_rcvd"`
    Frag_drop int `json:"frag_drop"`
    Dst_port_undef_drop int `json:"dst_port_undef_drop"`
    Dst_port_exceed_drop_any int `json:"dst_port_exceed_drop_any"`
    Dst_ipproto_bl int `json:"dst_ipproto_bl"`
    Dst_port_bl int `json:"dst_port_bl"`
    Dst_sport_bl int `json:"dst_sport_bl"`
    Dst_sport_exceed_drop_any int `json:"dst_sport_exceed_drop_any"`
    Dst_ipproto_rcvd int `json:"dst_ipproto_rcvd"`
    Dst_ipproto_drop int `json:"dst_ipproto_drop"`
    Dst_ipproto_exceed_drop_any int `json:"dst_ipproto_exceed_drop_any"`
    Src_ip_bypass int `json:"src_ip_bypass"`
    Dst_ingress_bytes int `json:"dst_ingress_bytes"`
    Dst_egress_bytes int `json:"dst_egress_bytes"`
    Dst_ingress_packets int `json:"dst_ingress_packets"`
    Dst_egress_packets int `json:"dst_egress_packets"`
    Dst_ip_bypass int `json:"dst_ip_bypass"`
    Dst_blackhole_inject int `json:"dst_blackhole_inject"`
    Dst_blackhole_withdraw int `json:"dst_blackhole_withdraw"`
    Tcp_total_bytes_rcv int `json:"tcp_total_bytes_rcv"`
    Tcp_total_bytes_drop int `json:"tcp_total_bytes_drop"`
    Udp_total_bytes_rcv int `json:"udp_total_bytes_rcv"`
    Udp_total_bytes_drop int `json:"udp_total_bytes_drop"`
    Icmp_total_bytes_rcv int `json:"icmp_total_bytes_rcv"`
    Icmp_total_bytes_drop int `json:"icmp_total_bytes_drop"`
    Other_total_bytes_rcv int `json:"other_total_bytes_rcv"`
    Other_total_bytes_drop int `json:"other_total_bytes_drop"`
    Udp_any_exceed int `json:"udp_any_exceed"`
    Tcp_any_exceed int `json:"tcp_any_exceed"`
    Icmp_any_exceed int `json:"icmp_any_exceed"`
    Other_any_exceed int `json:"other_any_exceed"`
    Tcp_drop_bl int `json:"tcp_drop_bl"`
    Udp_drop_bl int `json:"udp_drop_bl"`
    Icmp_drop_bl int `json:"icmp_drop_bl"`
    Other_drop_bl int `json:"other_drop_bl"`
    Glid_action_encap_send_immed int `json:"glid_action_encap_send_immed"`
    Glid_action_encap_send_delay int `json:"glid_action_encap_send_delay"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Dst_hw_drop_rule_inserted int `json:"dst_hw_drop_rule_inserted"`
    Dst_hw_drop_rule_removed int `json:"dst_hw_drop_rule_removed"`
    Src_hw_drop_rule_inserted int `json:"src_hw_drop_rule_inserted"`
    Src_hw_drop_rule_removed int `json:"src_hw_drop_rule_removed"`
}

func (p *DdosBriefStats) GetId() string{
    return "1"
}

func (p *DdosBriefStats) getPath() string{
    return "ddos/brief/stats"
}

func (p *DdosBriefStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosBriefStats,error) {
logger.Println("DdosBriefStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosBriefStats
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
