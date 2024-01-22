

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLiteGlobalStats struct {
    
    Stats Cgnv6DsLiteGlobalStatsStats `json:"stats"`

}
type DataCgnv6DsLiteGlobalStats struct {
    DtCgnv6DsLiteGlobalStats Cgnv6DsLiteGlobalStats `json:"global"`
}


type Cgnv6DsLiteGlobalStatsStats struct {
    Total_tcp_allocated int `json:"total_tcp_allocated"`
    Total_tcp_freed int `json:"total_tcp_freed"`
    Total_udp_allocated int `json:"total_udp_allocated"`
    Total_udp_freed int `json:"total_udp_freed"`
    Total_icmp_allocated int `json:"total_icmp_allocated"`
    Total_icmp_freed int `json:"total_icmp_freed"`
    Data_session_created int `json:"data_session_created"`
    Data_session_freed int `json:"data_session_freed"`
    User_quota_created int `json:"user_quota_created"`
    User_quota_put_in_del_q int `json:"user_quota_put_in_del_q"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    New_user_resource_unavailable int `json:"new_user_resource_unavailable"`
    Tcp_user_quota_exceeded int `json:"tcp_user_quota_exceeded"`
    Udp_user_quota_exceeded int `json:"udp_user_quota_exceeded"`
    Icmp_user_quota_exceeded int `json:"icmp_user_quota_exceeded"`
    Extended_quota_matched int `json:"extended_quota_matched"`
    Extended_quota_exceeded int `json:"extended_quota_exceeded"`
    Data_sesn_user_quota_exceeded int `json:"data_sesn_user_quota_exceeded"`
    Data_sesn_rate_user_quota_exceeded int `json:"data_sesn_rate_user_quota_exceeded"`
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    Fullcone_failure int `json:"fullcone_failure"`
    Hairpin int `json:"hairpin"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Endpoint_indep_map_match int `json:"endpoint_indep_map_match"`
    Endpoint_indep_filter_match int `json:"endpoint_indep_filter_match"`
    Inbound_filtered int `json:"inbound_filtered"`
    Eif_limit_exceeded int `json:"eif_limit_exceeded"`
    Total_tcp_overloaded int `json:"total_tcp_overloaded"`
    Total_udp_overloaded int `json:"total_udp_overloaded"`
    Port_overloading_smp_inserted_tcp int `json:"port_overloading_smp_inserted_tcp"`
    Port_overloading_smp_inserted_udp int `json:"port_overloading_smp_inserted_udp"`
    Port_overloading_smp_free_tcp int `json:"port_overloading_smp_free_tcp"`
    Port_overloading_smp_free_udp int `json:"port_overloading_smp_free_udp"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    No_radius_profile_match int `json:"no_radius_profile_match"`
    Truncated_packet int `json:"truncated_packet"`
    Lid_drop int `json:"lid_drop"`
    Lid_pass_through int `json:"lid_pass_through"`
    No_class_list_match int `json:"no_class_list_match"`
    Class_list_permit_drop int `json:"class_list_permit_drop"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
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
    Prefix_quota_mismatch int `json:"prefix_quota_mismatch"`
    Port_overloading_port_tcp_inserted int `json:"port_overloading_port_tcp_inserted"`
    Port_overloading_port_udp_inserted int `json:"port_overloading_port_udp_inserted"`
    Port_overloading_port_free_tcp int `json:"port_overloading_port_free_tcp"`
    Port_overloading_port_free_udp int `json:"port_overloading_port_free_udp"`
}

func (p *Cgnv6DsLiteGlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6DsLiteGlobalStats) getPath() string{
    return "cgnv6/ds-lite/global/stats"
}

func (p *Cgnv6DsLiteGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DsLiteGlobalStats,error) {
logger.Println("Cgnv6DsLiteGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DsLiteGlobalStats
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
