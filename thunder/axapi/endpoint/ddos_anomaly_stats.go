

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosAnomalyStats struct {
    
    Stats DdosAnomalyStatsStats `json:"stats"`

}
type DataDdosAnomalyStats struct {
    DtDdosAnomalyStats DdosAnomalyStats `json:"anomaly"`
}


type DdosAnomalyStatsStats struct {
    Land_attack int `json:"land_attack"`
    Empty_frag int `json:"empty_frag"`
    Micro_frag int `json:"micro_frag"`
    Ipv4_opt int `json:"ipv4_opt"`
    Ip_frag int `json:"ip_frag"`
    Bad_ip_hdr_len int `json:"bad_ip_hdr_len"`
    Bad_ip_flags int `json:"bad_ip_flags"`
    Bad_ip_ttl int `json:"bad_ip_ttl"`
    No_ip_payload int `json:"no_ip_payload"`
    Oversize_ip_pl int `json:"oversize_ip_pl"`
    Bad_ip_pl_len int `json:"bad_ip_pl_len"`
    Bad_ip_frag_off int `json:"bad_ip_frag_off"`
    Bad_ip_csum int `json:"bad_ip_csum"`
    Icmp_pod int `json:"icmp_pod"`
    Tcp_bad_urg_off int `json:"tcp_bad_urg_off"`
    Tcp_short_hdr int `json:"tcp_short_hdr"`
    Tcp_bad_ip_len int `json:"tcp_bad_ip_len"`
    Tcp_null_flags int `json:"tcp_null_flags"`
    Tcp_null_scan int `json:"tcp_null_scan"`
    Tcp_syn_fin int `json:"tcp_syn_fin"`
    Tcp_xmas_flags int `json:"tcp_xmas_flags"`
    Tcp_xmas_scan int `json:"tcp_xmas_scan"`
    Tcp_syn_frag int `json:"tcp_syn_frag"`
    Tcp_frag_header int `json:"tcp_frag_header"`
    Tcp_bad_csum int `json:"tcp_bad_csum"`
    Udp_short_hdr int `json:"udp_short_hdr"`
    Udp_short_leng int `json:"udp_short_leng"`
    Udp_kb_frag int `json:"udp_kb_frag"`
    Udp_port_lb int `json:"udp_port_lb"`
    Udp_bad_csum int `json:"udp_bad_csum"`
    Runt_ip_hdr int `json:"runt_ip_hdr"`
    Runt_tcpudp_hdr int `json:"runt_tcpudp_hdr"`
    Tun_mismatch int `json:"tun_mismatch"`
    Tcp_opt_overflow int `json:"tcp_opt_overflow"`
}

func (p *DdosAnomalyStats) GetId() string{
    return "1"
}

func (p *DdosAnomalyStats) getPath() string{
    return "ddos/anomaly/stats"
}

func (p *DdosAnomalyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosAnomalyStats,error) {
logger.Println("DdosAnomalyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosAnomalyStats
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
