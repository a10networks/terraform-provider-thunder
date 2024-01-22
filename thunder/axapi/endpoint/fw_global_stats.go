

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwGlobalStats struct {
    
    Stats FwGlobalStatsStats `json:"stats"`

}
type DataFwGlobalStats struct {
    DtFwGlobalStats FwGlobalStats `json:"global"`
}


type FwGlobalStatsStats struct {
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Data_session_created int `json:"data_session_created"`
    Data_session_created_local int `json:"data_session_created_local"`
    Data_session_freed int `json:"data_session_freed"`
    Data_session_freed_local int `json:"data_session_freed_local"`
    Dyn_blist_sess_created int `json:"dyn_blist_sess_created"`
    Dyn_blist_sess_freed int `json:"dyn_blist_sess_freed"`
    Dyn_blist_pkt_drop int `json:"dyn_blist_pkt_drop"`
    Active_fullcone_session int `json:"active_fullcone_session"`
    LimitEntryCreated int `json:"limit-entry-created"`
    LimitEntryMarkedDeleted int `json:"limit-entry-marked-deleted"`
    UndeterminedRuleCounter int `json:"undetermined-rule-counter"`
    Non_syn_pkt_fwd_allowed int `json:"non_syn_pkt_fwd_allowed"`
    Fwd_ingress_packets_tcp int `json:"fwd_ingress_packets_tcp"`
    Fwd_egress_packets_tcp int `json:"fwd_egress_packets_tcp"`
    Rev_ingress_packets_tcp int `json:"rev_ingress_packets_tcp"`
    Rev_egress_packets_tcp int `json:"rev_egress_packets_tcp"`
    Fwd_ingress_bytes_tcp int `json:"fwd_ingress_bytes_tcp"`
    Fwd_egress_bytes_tcp int `json:"fwd_egress_bytes_tcp"`
    Rev_ingress_bytes_tcp int `json:"rev_ingress_bytes_tcp"`
    Rev_egress_bytes_tcp int `json:"rev_egress_bytes_tcp"`
    Fwd_ingress_packets_udp int `json:"fwd_ingress_packets_udp"`
    Fwd_egress_packets_udp int `json:"fwd_egress_packets_udp"`
    Rev_ingress_packets_udp int `json:"rev_ingress_packets_udp"`
    Rev_egress_packets_udp int `json:"rev_egress_packets_udp"`
    Fwd_ingress_bytes_udp int `json:"fwd_ingress_bytes_udp"`
    Fwd_egress_bytes_udp int `json:"fwd_egress_bytes_udp"`
    Rev_ingress_bytes_udp int `json:"rev_ingress_bytes_udp"`
    Rev_egress_bytes_udp int `json:"rev_egress_bytes_udp"`
    Fwd_ingress_packets_icmp int `json:"fwd_ingress_packets_icmp"`
    Fwd_egress_packets_icmp int `json:"fwd_egress_packets_icmp"`
    Rev_ingress_packets_icmp int `json:"rev_ingress_packets_icmp"`
    Rev_egress_packets_icmp int `json:"rev_egress_packets_icmp"`
    Fwd_ingress_bytes_icmp int `json:"fwd_ingress_bytes_icmp"`
    Fwd_egress_bytes_icmp int `json:"fwd_egress_bytes_icmp"`
    Rev_ingress_bytes_icmp int `json:"rev_ingress_bytes_icmp"`
    Rev_egress_bytes_icmp int `json:"rev_egress_bytes_icmp"`
    Fwd_ingress_packets_others int `json:"fwd_ingress_packets_others"`
    Fwd_egress_packets_others int `json:"fwd_egress_packets_others"`
    Rev_ingress_packets_others int `json:"rev_ingress_packets_others"`
    Rev_egress_packets_others int `json:"rev_egress_packets_others"`
    Fwd_ingress_bytes_others int `json:"fwd_ingress_bytes_others"`
    Fwd_egress_bytes_others int `json:"fwd_egress_bytes_others"`
    Rev_ingress_bytes_others int `json:"rev_ingress_bytes_others"`
    Rev_egress_bytes_others int `json:"rev_egress_bytes_others"`
    Fwd_ingress_pkt_size_range1 int `json:"fwd_ingress_pkt_size_range1"`
    Fwd_ingress_pkt_size_range2 int `json:"fwd_ingress_pkt_size_range2"`
    Fwd_ingress_pkt_size_range3 int `json:"fwd_ingress_pkt_size_range3"`
    Fwd_ingress_pkt_size_range4 int `json:"fwd_ingress_pkt_size_range4"`
    Fwd_egress_pkt_size_range1 int `json:"fwd_egress_pkt_size_range1"`
    Fwd_egress_pkt_size_range2 int `json:"fwd_egress_pkt_size_range2"`
    Fwd_egress_pkt_size_range3 int `json:"fwd_egress_pkt_size_range3"`
    Fwd_egress_pkt_size_range4 int `json:"fwd_egress_pkt_size_range4"`
    Rev_ingress_pkt_size_range1 int `json:"rev_ingress_pkt_size_range1"`
    Rev_ingress_pkt_size_range2 int `json:"rev_ingress_pkt_size_range2"`
    Rev_ingress_pkt_size_range3 int `json:"rev_ingress_pkt_size_range3"`
    Rev_ingress_pkt_size_range4 int `json:"rev_ingress_pkt_size_range4"`
    Rev_egress_pkt_size_range1 int `json:"rev_egress_pkt_size_range1"`
    Rev_egress_pkt_size_range2 int `json:"rev_egress_pkt_size_range2"`
    Rev_egress_pkt_size_range3 int `json:"rev_egress_pkt_size_range3"`
    Rev_egress_pkt_size_range4 int `json:"rev_egress_pkt_size_range4"`
}

func (p *FwGlobalStats) GetId() string{
    return "1"
}

func (p *FwGlobalStats) getPath() string{
    return "fw/global/stats"
}

func (p *FwGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwGlobalStats,error) {
logger.Println("FwGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwGlobalStats
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
