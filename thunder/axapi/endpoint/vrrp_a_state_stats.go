

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAStateStats struct {
    
    Stats VrrpAStateStatsStats `json:"stats"`

}
type DataVrrpAStateStats struct {
    DtVrrpAStateStats VrrpAStateStats `json:"state"`
}


type VrrpAStateStatsStats struct {
    Sync_pkt_tx_counter int `json:"sync_pkt_tx_counter"`
    Sync_pkt_rcv_counter int `json:"sync_pkt_rcv_counter"`
    Sync_rx_create_counter int `json:"sync_rx_create_counter"`
    Sync_rx_del_counter int `json:"sync_rx_del_counter"`
    Sync_rx_update_age_counter int `json:"sync_rx_update_age_counter"`
    Sync_tx_create_counter int `json:"sync_tx_create_counter"`
    Sync_tx_del_counter int `json:"sync_tx_del_counter"`
    Sync_tx_update_age_counter int `json:"sync_tx_update_age_counter"`
    Sync_rx_persist_create_counter int `json:"sync_rx_persist_create_counter"`
    Sync_rx_persist_del_counter int `json:"sync_rx_persist_del_counter"`
    Sync_rx_persist_update_age_counter int `json:"sync_rx_persist_update_age_counter"`
    Sync_tx_persist_create_counter int `json:"sync_tx_persist_create_counter"`
    Sync_tx_persist_del_counter int `json:"sync_tx_persist_del_counter"`
    Sync_tx_persist_update_age_counter int `json:"sync_tx_persist_update_age_counter"`
    Query_pkt_tx_counter int `json:"query_pkt_tx_counter"`
    Query_pkt_rcv_counter int `json:"query_pkt_rcv_counter"`
    Sync_tx_smp_radius_table_counter int `json:"sync_tx_smp_radius_table_counter"`
    Sync_rx_smp_radius_table_counter int `json:"sync_rx_smp_radius_table_counter"`
    Query_tx_max_packed int `json:"query_tx_max_packed"`
    Query_tx_min_packed int `json:"query_tx_min_packed"`
    Query_pkt_invalid_idx_counter int `json:"query_pkt_invalid_idx_counter"`
    Query_tx_get_buff_failed int `json:"query_tx_get_buff_failed"`
    Query_rx_zero_info_counter int `json:"query_rx_zero_info_counter"`
    Query_rx_full_info_counter int `json:"query_rx_full_info_counter"`
    Query_rx_unk_counter int `json:"query_rx_unk_counter"`
    Sync_pkt_invalid_idx_counter int `json:"sync_pkt_invalid_idx_counter"`
    Sync_tx_get_buff_failed int `json:"sync_tx_get_buff_failed"`
    Sync_tx_total_info_counter int `json:"sync_tx_total_info_counter"`
    Sync_tx_create_ext_bit_counter int `json:"sync_tx_create_ext_bit_counter"`
    Sync_tx_update_seqnos_counter int `json:"sync_tx_update_seqnos_counter"`
    Sync_tx_min_packed int `json:"sync_tx_min_packed"`
    Sync_tx_max_packed int `json:"sync_tx_max_packed"`
    Sync_rx_len_invalid int `json:"sync_rx_len_invalid"`
    Sync_persist_rx_len_invalid int `json:"sync_persist_rx_len_invalid"`
    Sync_persist_rx_proto_not_supported int `json:"sync_persist_rx_proto_not_supported"`
    Sync_persist_rx_type_invalid int `json:"sync_persist_rx_type_invalid"`
    Sync_persist_rx_cannot_process_mandatory int `json:"sync_persist_rx_cannot_process_mandatory"`
    Sync_persist_rx_ext_bit_process_error int `json:"sync_persist_rx_ext_bit_process_error"`
    Sync_persist_rx_no_such_vport int `json:"sync_persist_rx_no_such_vport"`
    Sync_persist_rx_vporttype_not_supported int `json:"sync_persist_rx_vporttype_not_supported"`
    Sync_persist_rx_no_such_rport int `json:"sync_persist_rx_no_such_rport"`
    Sync_persist_rx_no_such_sg_group int `json:"sync_persist_rx_no_such_sg_group"`
    Sync_persist_rx_no_sg_group_info int `json:"sync_persist_rx_no_sg_group_info"`
    Sync_persist_rx_conn_get_failed int `json:"sync_persist_rx_conn_get_failed"`
    Sync_rx_no_such_vport int `json:"sync_rx_no_such_vport"`
    Sync_rx_no_such_rport int `json:"sync_rx_no_such_rport"`
    Sync_rx_cannot_process_mandatory int `json:"sync_rx_cannot_process_mandatory"`
    Sync_rx_ext_bit_process_error int `json:"sync_rx_ext_bit_process_error"`
    Sync_rx_create_ext_bit_counter int `json:"sync_rx_create_ext_bit_counter"`
    Sync_rx_conn_exists int `json:"sync_rx_conn_exists"`
    Sync_rx_conn_get_failed int `json:"sync_rx_conn_get_failed"`
    Sync_rx_proto_not_supported int `json:"sync_rx_proto_not_supported"`
    Sync_rx_no_dst_for_vport_inline int `json:"sync_rx_no_dst_for_vport_inline"`
    Sync_rx_no_such_nat_pool int `json:"sync_rx_no_such_nat_pool"`
    Sync_rx_no_such_sg_node int `json:"sync_rx_no_such_sg_node"`
    Sync_rx_del_no_such_session int `json:"sync_rx_del_no_such_session"`
    Sync_rx_type_invalid int `json:"sync_rx_type_invalid"`
    Sync_rx_zero_info_counter int `json:"sync_rx_zero_info_counter"`
    Sync_rx_dcmsg_counter int `json:"sync_rx_dcmsg_counter"`
    Sync_rx_total_info_counter int `json:"sync_rx_total_info_counter"`
    Sync_rx_update_seqnos_counter int `json:"sync_rx_update_seqnos_counter"`
    Sync_rx_unk_counter int `json:"sync_rx_unk_counter"`
    Sync_rx_apptype_not_supported int `json:"sync_rx_apptype_not_supported"`
    Sync_query_dcmsg_counter int `json:"sync_query_dcmsg_counter"`
    Sync_get_buff_failed_rt int `json:"sync_get_buff_failed_rt"`
    Sync_get_buff_failed_port int `json:"sync_get_buff_failed_port"`
    Sync_rx_lsn_create_sby int `json:"sync_rx_lsn_create_sby"`
    Sync_rx_nat_create_sby int `json:"sync_rx_nat_create_sby"`
    Sync_rx_nat_alloc_sby int `json:"sync_rx_nat_alloc_sby"`
    Sync_rx_insert_tuple int `json:"sync_rx_insert_tuple"`
    Sync_rx_sfw int `json:"sync_rx_sfw"`
    Sync_rx_create_static_sby int `json:"sync_rx_create_static_sby"`
    Sync_rx_ext_pptp int `json:"sync_rx_ext_pptp"`
    Sync_rx_ext_rtsp int `json:"sync_rx_ext_rtsp"`
    Sync_rx_reserve_ha int `json:"sync_rx_reserve_ha"`
    Sync_rx_seq_deltas int `json:"sync_rx_seq_deltas"`
    Sync_rx_ftp_control int `json:"sync_rx_ftp_control"`
    Sync_rx_ext_lsn_acl int `json:"sync_rx_ext_lsn_acl"`
    Sync_rx_ext_lsn_ac_idle_timeout int `json:"sync_rx_ext_lsn_ac_idle_timeout"`
    Sync_rx_ext_sip_alg int `json:"sync_rx_ext_sip_alg"`
    Sync_rx_ext_h323_alg int `json:"sync_rx_ext_h323_alg"`
    Sync_rx_ext_nat_mac int `json:"sync_rx_ext_nat_mac"`
    Sync_tx_lsn_fullcone int `json:"sync_tx_lsn_fullcone"`
    Sync_rx_lsn_fullcone int `json:"sync_rx_lsn_fullcone"`
    Sync_err_lsn_fullcone int `json:"sync_err_lsn_fullcone"`
    Sync_tx_update_sctp_conn_addr int `json:"sync_tx_update_sctp_conn_addr"`
    Sync_rx_update_sctp_conn_addr int `json:"sync_rx_update_sctp_conn_addr"`
    Sync_rx_ext_nat_alg_tcp_info int `json:"sync_rx_ext_nat_alg_tcp_info"`
    Sync_rx_ext_dcfw_rule_id int `json:"sync_rx_ext_dcfw_rule_id"`
    Sync_rx_ext_dcfw_log int `json:"sync_rx_ext_dcfw_log"`
    Sync_rx_estab_counter int `json:"sync_rx_estab_counter"`
    Sync_tx_estab_counter int `json:"sync_tx_estab_counter"`
    Sync_rx_zone_failure_counter int `json:"sync_rx_zone_failure_counter"`
    Sync_rx_ext_fw_http_logging int `json:"sync_rx_ext_fw_http_logging"`
    Sync_rx_ext_dcfw_rule_idle_timeout int `json:"sync_rx_ext_dcfw_rule_idle_timeout"`
    Sync_rx_ext_fw_gtp_info int `json:"sync_rx_ext_fw_gtp_info"`
    Sync_rx_not_expect_sync_pkt int `json:"sync_rx_not_expect_sync_pkt"`
    Sync_rx_ext_fw_apps int `json:"sync_rx_ext_fw_apps"`
    Sync_tx_mon_entity int `json:"sync_tx_mon_entity"`
    Sync_rx_mon_entity int `json:"sync_rx_mon_entity"`
    Sync_rx_ext_fw_gtp_log_info int `json:"sync_rx_ext_fw_gtp_log_info"`
    Sync_rx_ext_fw_gtp_u_info int `json:"sync_rx_ext_fw_gtp_u_info"`
    Sync_rx_ext_fw_gtp_ext_info int `json:"sync_rx_ext_fw_gtp_ext_info"`
    Sync_rx_ext_fw_gtp_log_ext_info int `json:"sync_rx_ext_fw_gtp_log_ext_info"`
    Sync_rx_ddos_drop_counter int `json:"sync_rx_ddos_drop_counter"`
    Sync_rx_invalid_sync_packet_counter int `json:"sync_rx_invalid_sync_packet_counter"`
    Sync_pkt_empty_buff_counter int `json:"sync_pkt_empty_buff_counter"`
    Sync_pkt_no_sending_vgrp_counter int `json:"sync_pkt_no_sending_vgrp_counter"`
    Sync_pkt_no_receiving_vgrp_counter int `json:"sync_pkt_no_receiving_vgrp_counter"`
    Query_pkt_no_receiving_ip_counter int `json:"query_pkt_no_receiving_ip_counter"`
    Sync_pkt_failed_buff_copy_counter int `json:"sync_pkt_failed_buff_copy_counter"`
    Sync_rx_bad_protocol_counter int `json:"sync_rx_bad_protocol_counter"`
    Sync_rx_no_vgrp_counter int `json:"sync_rx_no_vgrp_counter"`
    Sync_rx_by_inactive_peer_counter int `json:"sync_rx_by_inactive_peer_counter"`
    Sync_rx_table_entry_update_counter int `json:"sync_rx_table_entry_update_counter"`
    Sync_rx_table_entry_create_counter int `json:"sync_rx_table_entry_create_counter"`
    Sync_rx_table_entry_del_counter int `json:"sync_rx_table_entry_del_counter"`
    Sync_rx_aflex_update_counter int `json:"sync_rx_aflex_update_counter"`
    Sync_rx_aflex_create_counter int `json:"sync_rx_aflex_create_counter"`
    Sync_rx_aflex_del_counter int `json:"sync_rx_aflex_del_counter"`
    Sync_rx_aflex_frag_counter int `json:"sync_rx_aflex_frag_counter"`
    Query_rx_invalid_partition_counter int `json:"query_rx_invalid_partition_counter"`
    Query_rx_invalid_ha_group_counter int `json:"query_rx_invalid_ha_group_counter"`
    Query_rx_invalid_sync_version_counter int `json:"query_rx_invalid_sync_version_counter"`
    Query_rx_invalid_msg_dir_counter int `json:"query_rx_invalid_msg_dir_counter"`
    Sync_rx_out_of_order_pkt_counter int `json:"sync_rx_out_of_order_pkt_counter"`
    Sync_rx_unreached_pkt_counter int `json:"sync_rx_unreached_pkt_counter"`
    Sync_rx_ext_fw_gtp_echo_ext_info int `json:"sync_rx_ext_fw_gtp_echo_ext_info"`
    Sync_rx_smp_create_counter int `json:"sync_rx_smp_create_counter"`
    Sync_rx_smp_delete_counter int `json:"sync_rx_smp_delete_counter"`
    Sync_rx_smp_update_counter int `json:"sync_rx_smp_update_counter"`
    Sync_tx_smp_create_counter int `json:"sync_tx_smp_create_counter"`
    Sync_tx_smp_delete_counter int `json:"sync_tx_smp_delete_counter"`
    Sync_tx_smp_update_counter int `json:"sync_tx_smp_update_counter"`
    Sync_rx_smp_clear_counter int `json:"sync_rx_smp_clear_counter"`
    Sync_tx_smp_clear_counter int `json:"sync_tx_smp_clear_counter"`
    Sync_rx_ext_fw_so_shadow_ext_info int `json:"sync_rx_ext_fw_so_shadow_ext_info"`
    Sync_tx_aflex_table_entry_add_counter int `json:"sync_tx_aflex_table_entry_add_counter"`
    Sync_rx_aflex_table_entry_add_counter int `json:"sync_rx_aflex_table_entry_add_counter"`
    Sync_tx_aflex_table_entry_append_counter int `json:"sync_tx_aflex_table_entry_append_counter"`
    Sync_rx_aflex_table_entry_append_counter int `json:"sync_rx_aflex_table_entry_append_counter"`
    Sync_tx_aflex_table_entry_delete_counter int `json:"sync_tx_aflex_table_entry_delete_counter"`
    Sync_rx_aflex_table_entry_delete_counter int `json:"sync_rx_aflex_table_entry_delete_counter"`
    Sync_tx_aflex_table_entry_incr_counter int `json:"sync_tx_aflex_table_entry_incr_counter"`
    Sync_rx_aflex_table_entry_incr_counter int `json:"sync_rx_aflex_table_entry_incr_counter"`
    Sync_tx_aflex_table_entry_lookup_counter int `json:"sync_tx_aflex_table_entry_lookup_counter"`
    Sync_rx_aflex_table_entry_lookup_counter int `json:"sync_rx_aflex_table_entry_lookup_counter"`
    Sync_tx_aflex_table_entry_lifetime_counter int `json:"sync_tx_aflex_table_entry_lifetime_counter"`
    Sync_rx_aflex_table_entry_lifetime_counter int `json:"sync_rx_aflex_table_entry_lifetime_counter"`
    Sync_tx_aflex_table_entry_replace_counter int `json:"sync_tx_aflex_table_entry_replace_counter"`
    Sync_rx_aflex_table_entry_replace_counter int `json:"sync_rx_aflex_table_entry_replace_counter"`
    Sync_tx_aflex_table_entry_set_counter int `json:"sync_tx_aflex_table_entry_set_counter"`
    Sync_rx_aflex_table_entry_set_counter int `json:"sync_rx_aflex_table_entry_set_counter"`
    Sync_tx_aflex_table_entry_timeout_counter int `json:"sync_tx_aflex_table_entry_timeout_counter"`
    Sync_rx_aflex_table_entry_timeout_counter int `json:"sync_rx_aflex_table_entry_timeout_counter"`
    Sync_tx_aflex_table_entry_fastsync_counter int `json:"sync_tx_aflex_table_entry_fastsync_counter"`
    Sync_rx_aflex_table_entry_fastsync_counter int `json:"sync_rx_aflex_table_entry_fastsync_counter"`
    Sync_tx_aflex_table_entry_error_counter int `json:"sync_tx_aflex_table_entry_error_counter"`
    Sync_tx_aflex_table_entry_not_eligible_counter int `json:"sync_tx_aflex_table_entry_not_eligible_counter"`
    Sync_rx_ext_fw_limit_entry int `json:"sync_rx_ext_fw_limit_entry"`
    Sync_tx_fw_set_dscp_counter int `json:"sync_tx_fw_set_dscp_counter"`
    Sync_rx_fw_set_dscp_counter int `json:"sync_rx_fw_set_dscp_counter"`
    Dns_cache_sync_tx_create_counter int `json:"dns_cache_sync_tx_create_counter"`
    Dns_cache_sync_rx_create_counter int `json:"dns_cache_sync_rx_create_counter"`
    Dns_cache_sync_tx_del_counter int `json:"dns_cache_sync_tx_del_counter"`
    Dns_cache_sync_rx_del_counter int `json:"dns_cache_sync_rx_del_counter"`
    Dns_cache_sync_tx_frag_counter int `json:"dns_cache_sync_tx_frag_counter"`
    Dns_cache_sync_rx_frag_counter int `json:"dns_cache_sync_rx_frag_counter"`
    Dns_cache_sync_tx_error_counter int `json:"dns_cache_sync_tx_error_counter"`
    Dns_cache_sync_rx_error_counter int `json:"dns_cache_sync_rx_error_counter"`
}

func (p *VrrpAStateStats) GetId() string{
    return "1"
}

func (p *VrrpAStateStats) getPath() string{
    return "vrrp-a/state/stats"
}

func (p *VrrpAStateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAStateStats,error) {
logger.Println("VrrpAStateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAStateStats
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
