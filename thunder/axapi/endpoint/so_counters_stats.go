

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SoCountersStats struct {
    
    Stats SoCountersStatsStats `json:"stats"`

}
type DataSoCountersStats struct {
    DtSoCountersStats SoCountersStats `json:"so-counters"`
}


type SoCountersStatsStats struct {
    So_pkts_rcvd int `json:"so_pkts_rcvd"`
    So_redirect_pkts_sent int `json:"so_redirect_pkts_sent"`
    So_pkts_dropped int `json:"so_pkts_dropped"`
    So_redirected_pkts_rcvd int `json:"so_redirected_pkts_rcvd"`
    So_fw_shadow_session_created int `json:"so_fw_shadow_session_created"`
    So_pkts_traffic_map_not_found_drop int `json:"so_pkts_traffic_map_not_found_drop"`
    So_slb_pkts_redirect_conn_aged_out int `json:"so_slb_pkts_redirect_conn_aged_out"`
    So_pkts_scaleout_not_active_drop int `json:"so_pkts_scaleout_not_active_drop"`
    So_pkts_slb_nat_reserve_fail int `json:"so_pkts_slb_nat_reserve_fail"`
    So_pkts_slb_nat_release_fail int `json:"so_pkts_slb_nat_release_fail"`
    So_pkts_dest_mac_mismatch_drop int `json:"so_pkts_dest_mac_mismatch_drop"`
    So_pkts_l2redirect_dest_mac_zero_drop int `json:"so_pkts_l2redirect_dest_mac_zero_drop"`
    So_pkts_l2redirect_interface_not_up int `json:"so_pkts_l2redirect_interface_not_up"`
    So_pkts_l2redirect_invalid_redirect_info_error int `json:"so_pkts_l2redirect_invalid_redirect_info_error"`
    So_pkts_l3_redirect_encap_error_drop int `json:"so_pkts_l3_redirect_encap_error_drop"`
    So_pkts_l3_redirect_inner_mac_zero_drop int `json:"so_pkts_l3_redirect_inner_mac_zero_drop"`
    So_pkts_l3_redirect_decap_vlan_sanity_drop int `json:"so_pkts_l3_redirect_decap_vlan_sanity_drop"`
    So_pkts_l3_redirect_decap_non_ipv4_vxlan_drop int `json:"so_pkts_l3_redirect_decap_non_ipv4_vxlan_drop"`
    So_pkts_l3_redirect_decap_rx_encap_params_drop int `json:"so_pkts_l3_redirect_decap_rx_encap_params_drop"`
    So_pkts_l3_redirect_table_error int `json:"so_pkts_l3_redirect_table_error"`
    So_pkts_l3_redirect_rcvd_in_l2_mode_drop int `json:"so_pkts_l3_redirect_rcvd_in_l2_mode_drop"`
    So_pkts_l3_redirect_fragmentation_error int `json:"so_pkts_l3_redirect_fragmentation_error"`
    So_pkts_l3_redirect_table_no_entry_found int `json:"so_pkts_l3_redirect_table_no_entry_found"`
    So_pkts_l3_redirect_invalid_dev_dir int `json:"so_pkts_l3_redirect_invalid_dev_dir"`
    So_pkts_l3_redirect_chassis_dest_mac_error int `json:"so_pkts_l3_redirect_chassis_dest_mac_error"`
    So_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop int `json:"so_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop"`
    So_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop int `json:"so_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop"`
    So_pkts_l3_redirect_too_large_pkts_in_drop int `json:"so_pkts_l3_redirect_too_large_pkts_in_drop"`
    So_pkts_l2redirect_vlan_retrieval_error int `json:"so_pkts_l2redirect_vlan_retrieval_error"`
    So_pkts_l2redirect_port_retrieval_error int `json:"so_pkts_l2redirect_port_retrieval_error"`
    So_slb_shadow_session_created int `json:"so_slb_shadow_session_created"`
}

func (p *SoCountersStats) GetId() string{
    return "1"
}

func (p *SoCountersStats) getPath() string{
    return "so-counters/stats"
}

func (p *SoCountersStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSoCountersStats,error) {
logger.Println("SoCountersStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSoCountersStats
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
