

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepStats struct {
    
    Id1 int `json:"id1"`
    Stats OverlayTunnelVtepStatsStats `json:"stats"`

}
type DataOverlayTunnelVtepStats struct {
    DtOverlayTunnelVtepStats OverlayTunnelVtepStats `json:"vtep"`
}


type OverlayTunnelVtepStatsStats struct {
    Cfg_err_count int `json:"cfg_err_count"`
    Flooded_pkt_count int `json:"flooded_pkt_count"`
    Encap_unresolved_count int `json:"encap_unresolved_count"`
    Unknown_encap_rx_pkt int `json:"unknown_encap_rx_pkt"`
    Unknown_encap_tx_pkt int `json:"unknown_encap_tx_pkt"`
    Arp_req_sent int `json:"arp_req_sent"`
    Vtep_host_learned int `json:"vtep_host_learned"`
    Vtep_host_learn_error int `json:"vtep_host_learn_error"`
    Invalid_lif_rx int `json:"invalid_lif_rx"`
    Invalid_lif_tx int `json:"invalid_lif_tx"`
    Unknown_vtep_tx int `json:"unknown_vtep_tx"`
    Unknown_vtep_rx int `json:"unknown_vtep_rx"`
    Unhandled_pkt_rx int `json:"unhandled_pkt_rx"`
    Unhandled_pkt_tx int `json:"unhandled_pkt_tx"`
    Total_pkts_rx int `json:"total_pkts_rx"`
    Total_bytes_rx int `json:"total_bytes_rx"`
    Unicast_pkt_rx int `json:"unicast_pkt_rx"`
    Bcast_pkt_rx int `json:"bcast_pkt_rx"`
    Mcast_pkt_rx int `json:"mcast_pkt_rx"`
    Dropped_pkt_rx int `json:"dropped_pkt_rx"`
    Encap_miss_pkts_rx int `json:"encap_miss_pkts_rx"`
    Bad_chksum_pks_rx int `json:"bad_chksum_pks_rx"`
    Requeue_pkts_in int `json:"requeue_pkts_in"`
    Pkts_out int `json:"pkts_out"`
    Total_bytes_tx int `json:"total_bytes_tx"`
    Unicast_pkt_tx int `json:"unicast_pkt_tx"`
    Bcast_pkt_tx int `json:"bcast_pkt_tx"`
    Mcast_pkt_tx int `json:"mcast_pkt_tx"`
    Dropped_pkts_tx int `json:"dropped_pkts_tx"`
    Large_pkts_rx int `json:"large_pkts_rx"`
    Dot1q_pkts_rx int `json:"dot1q_pkts_rx"`
    Frag_pkts_tx int `json:"frag_pkts_tx"`
    Reassembled_pkts_rx int `json:"reassembled_pkts_rx"`
    Bad_inner_ipv4_len_rx int `json:"bad_inner_ipv4_len_rx"`
    Bad_inner_ipv6_len_rx int `json:"bad_inner_ipv6_len_rx"`
    Frag_drop_pkts_tx int `json:"frag_drop_pkts_tx"`
    Lif_un_init_rx int `json:"lif_un_init_rx"`
}

func (p *OverlayTunnelVtepStats) GetId() string{
    return "1"
}

func (p *OverlayTunnelVtepStats) getPath() string{
    
    return "overlay-tunnel/vtep/" +strconv.Itoa(p.Id1)+"/stats"
}

func (p *OverlayTunnelVtepStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataOverlayTunnelVtepStats,error) {
logger.Println("OverlayTunnelVtepStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataOverlayTunnelVtepStats
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
