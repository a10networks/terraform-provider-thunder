package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAStateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_state_stats`: Statistics for the object state\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAStateStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sync_pkt_tx_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Sent counter",
						},
						"sync_pkt_rcv_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Received counter",
						},
						"sync_rx_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Session Received counter",
						},
						"sync_rx_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Del Session Received counter",
						},
						"sync_rx_update_age_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Age Received counter",
						},
						"sync_tx_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Session Sent counter",
						},
						"sync_tx_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Del Session Sent counter",
						},
						"sync_tx_update_age_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Age Sent counter",
						},
						"sync_rx_persist_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Persist Session Pkts Received counter",
						},
						"sync_rx_persist_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Delete Persist Session Pkts Received counter",
						},
						"sync_rx_persist_update_age_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Persist Age Pkts Received counter",
						},
						"sync_tx_persist_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Persist Session Pkts Sent counter",
						},
						"sync_tx_persist_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Delete Persist Session Pkts Sent counter",
						},
						"sync_tx_persist_update_age_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Persist Age Pkts Sent counter",
						},
						"query_pkt_tx_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query sent counter",
						},
						"query_pkt_rcv_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Received counter",
						},
						"sync_tx_smp_radius_table_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update LSN RADIUS Sent counter",
						},
						"sync_rx_smp_radius_table_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update LSN RADIUS Received counter",
						},
						"query_tx_max_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Query Msg Per Packet",
						},
						"query_tx_min_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Min Query Msg Per Packet",
						},
						"query_pkt_invalid_idx_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Invalid Interface",
						},
						"query_tx_get_buff_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Get Buff Failure",
						},
						"query_rx_zero_info_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Packet Empty",
						},
						"query_rx_full_info_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Packet Full",
						},
						"query_rx_unk_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Query Unknown Type",
						},
						"sync_pkt_invalid_idx_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Invalid Interface",
						},
						"sync_tx_get_buff_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Get Buff Failure",
						},
						"sync_tx_total_info_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Total Info Pkts Sent counter",
						},
						"sync_tx_create_ext_bit_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create with Ext Sent counter",
						},
						"sync_tx_update_seqnos_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Seq Num Sent counter",
						},
						"sync_tx_min_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Sync Msg Per Packet",
						},
						"sync_tx_max_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Min Sync Msg Per Packet",
						},
						"sync_rx_len_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Length Invalid",
						},
						"sync_persist_rx_len_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Length Invalid",
						},
						"sync_persist_rx_proto_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Protocol Invalid",
						},
						"sync_persist_rx_type_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Type Invalid",
						},
						"sync_persist_rx_cannot_process_mandatory": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Process Mandatory Invalid",
						},
						"sync_persist_rx_ext_bit_process_error": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Proc Ext Bit Failure",
						},
						"sync_persist_rx_no_such_vport": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Virt Port Not Found",
						},
						"sync_persist_rx_vporttype_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Virt Port Type Invalid",
						},
						"sync_persist_rx_no_such_rport": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Real Port Not Found",
						},
						"sync_persist_rx_no_such_sg_group": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync No Service Group Found",
						},
						"sync_persist_rx_no_sg_group_info": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync No Service Group Info Found",
						},
						"sync_persist_rx_conn_get_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Persist Conn Sync Get Conn Failure",
						},
						"sync_rx_no_such_vport": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Virt Port Not Found",
						},
						"sync_rx_no_such_rport": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Real Port Not Found",
						},
						"sync_rx_cannot_process_mandatory": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Process Mandatory Invalid",
						},
						"sync_rx_ext_bit_process_error": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Proc Ext Bit Failure",
						},
						"sync_rx_create_ext_bit_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create with Ext Received counter",
						},
						"sync_rx_conn_exists": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Conn Exists",
						},
						"sync_rx_conn_get_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Get Conn Failure",
						},
						"sync_rx_proto_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Protocol Invalid",
						},
						"sync_rx_no_dst_for_vport_inline": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync 'dst' not found for vport inline",
						},
						"sync_rx_no_such_nat_pool": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync NAT Pool Error",
						},
						"sync_rx_no_such_sg_node": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync no SG node found",
						},
						"sync_rx_del_no_such_session": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Del Conn not Found",
						},
						"sync_rx_type_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Type Invalid",
						},
						"sync_rx_zero_info_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Packet Empty",
						},
						"sync_rx_dcmsg_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync forward CPU",
						},
						"sync_rx_total_info_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Total Info Pkts Received counter",
						},
						"sync_rx_update_seqnos_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update Seq Num Received counter",
						},
						"sync_rx_unk_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Unknown Type",
						},
						"sync_rx_apptype_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync App Type Invalid",
						},
						"sync_query_dcmsg_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync query forward CPU",
						},
						"sync_get_buff_failed_rt": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Get Buff Failure No Route",
						},
						"sync_get_buff_failed_port": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Get Buff Failure Wrong Port",
						},
						"sync_rx_lsn_create_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync LSN Create Standby",
						},
						"sync_rx_nat_create_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync NAT Create Standby",
						},
						"sync_rx_nat_alloc_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync NAT Alloc Standby",
						},
						"sync_rx_insert_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Insert Tuple",
						},
						"sync_rx_sfw": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync SFW",
						},
						"sync_rx_create_static_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Create Static Standby",
						},
						"sync_rx_ext_pptp": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Ext PPTP",
						},
						"sync_rx_ext_rtsp": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Ext RTSP",
						},
						"sync_rx_reserve_ha": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Reserve HA Conn",
						},
						"sync_rx_seq_deltas": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Seq Deltas Failure",
						},
						"sync_rx_ftp_control": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync FTP Control Failure",
						},
						"sync_rx_ext_lsn_acl": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync LSN ACL Failure",
						},
						"sync_rx_ext_lsn_ac_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync LSN ACL Idle Timeout Failure",
						},
						"sync_rx_ext_sip_alg": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync SIP TCP ALG Failure",
						},
						"sync_rx_ext_h323_alg": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync H323 TCP ALG Failure",
						},
						"sync_rx_ext_nat_mac": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync NAT MAC Failure",
						},
						"sync_tx_lsn_fullcone": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update LSN Fullcone Sent counter",
						},
						"sync_rx_lsn_fullcone": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Update LSN Fullcone Received counter",
						},
						"sync_err_lsn_fullcone": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync LSN Fullcone Failure",
						},
						"sync_tx_update_sctp_conn_addr": {
							Type: schema.TypeInt, Optional: true, Description: "Update SCTP Addresses Sent",
						},
						"sync_rx_update_sctp_conn_addr": {
							Type: schema.TypeInt, Optional: true, Description: "Update SCTP Addresses Received",
						},
						"sync_rx_ext_nat_alg_tcp_info": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync NAT ALG TCP Information",
						},
						"sync_rx_ext_dcfw_rule_id": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync FIREWALL session rule ID information Failure",
						},
						"sync_rx_ext_dcfw_log": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync FIREWALL session logging information Failure",
						},
						"sync_rx_estab_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync rcv established state",
						},
						"sync_tx_estab_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync send established state",
						},
						"sync_rx_zone_failure_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync Zone Failure",
						},
						"sync_rx_ext_fw_http_logging": {
							Type: schema.TypeInt, Optional: true, Description: "FW HTTP Logging Sync Failures",
						},
						"sync_rx_ext_dcfw_rule_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync FIREWALL session rule idle timeout information Failure",
						},
						"sync_rx_ext_fw_gtp_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP Info Received",
						},
						"sync_rx_not_expect_sync_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "unexpected session sync packets",
						},
						"sync_rx_ext_fw_apps": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync FIREWALL application information Failure",
						},
						"sync_tx_mon_entity": {
							Type: schema.TypeInt, Optional: true, Description: "Acos Monitoring Entities Sync Messages Sent",
						},
						"sync_rx_mon_entity": {
							Type: schema.TypeInt, Optional: true, Description: "Acos monitoring Entities Sync Messages Received",
						},
						"sync_rx_ext_fw_gtp_log_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP Log Info Received",
						},
						"sync_rx_ext_fw_gtp_u_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP U Info Received",
						},
						"sync_rx_ext_fw_gtp_ext_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP Ext Info Received",
						},
						"sync_rx_ext_fw_gtp_log_ext_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP Ext Log Info Received",
						},
						"sync_rx_ddos_drop_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive ddos protect packet",
						},
						"sync_rx_invalid_sync_packet_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive invalid packet",
						},
						"sync_pkt_empty_buff_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync drop sending packet for empty buffer",
						},
						"sync_pkt_no_sending_vgrp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync drop sending packet for invalid sending virtual group",
						},
						"sync_pkt_no_receiving_vgrp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync drop sending packet for invalid receiving virtual group",
						},
						"query_pkt_no_receiving_ip_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync drop sending packet for invalid receiving ip",
						},
						"sync_pkt_failed_buff_copy_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync drop sending packet for failure in sending buffer copy",
						},
						"sync_rx_bad_protocol_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with bad protocol",
						},
						"sync_rx_no_vgrp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with no virtual group",
						},
						"sync_rx_by_inactive_peer_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet by inactive peer",
						},
						"sync_rx_table_entry_update_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with table entry update",
						},
						"sync_rx_table_entry_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with table entry create",
						},
						"sync_rx_table_entry_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with table entry delete",
						},
						"sync_rx_aflex_update_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with aflex update",
						},
						"sync_rx_aflex_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with aflex create",
						},
						"sync_rx_aflex_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with aflex delete",
						},
						"sync_rx_aflex_frag_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive packet with aflex fragment",
						},
						"query_rx_invalid_partition_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive query packet with invalid partition",
						},
						"query_rx_invalid_ha_group_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive query packet with invalid ha group",
						},
						"query_rx_invalid_sync_version_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive query packet with invalid sync version",
						},
						"query_rx_invalid_msg_dir_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive query packet with invalid message dir",
						},
						"sync_rx_out_of_order_pkt_counter": {
							Type: schema.TypeInt, Optional: true, Description: "total number of out of order packets received",
						},
						"sync_rx_unreached_pkt_counter": {
							Type: schema.TypeInt, Optional: true, Description: "total number of unreached packets",
						},
						"sync_rx_ext_fw_gtp_echo_ext_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW GTP Echo Ext Info Received",
						},
						"sync_rx_smp_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Create SMP Session Pkts Received counter",
						},
						"sync_rx_smp_delete_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Delete SMP Session Pkts Received counter",
						},
						"sync_rx_smp_update_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Update SMP Session Pkts Received counter",
						},
						"sync_tx_smp_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Create SMP Session Pkts Sent counter",
						},
						"sync_tx_smp_delete_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Delete SMP Session Pkts Sent counter",
						},
						"sync_tx_smp_update_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Update SMP Session Pkts Sent counter",
						},
						"sync_rx_smp_clear_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Clear SMP Session Pkts Received counter",
						},
						"sync_tx_smp_clear_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Clear SMP Session Pkts Sent counter",
						},
						"sync_rx_ext_fw_so_shadow_ext_info": {
							Type: schema.TypeInt, Optional: true, Description: "FW Scaleout Shadow Ext Info Received",
						},
						"sync_tx_aflex_table_entry_add_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry add",
						},
						"sync_rx_aflex_table_entry_add_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry add",
						},
						"sync_tx_aflex_table_entry_append_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry append",
						},
						"sync_rx_aflex_table_entry_append_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry append",
						},
						"sync_tx_aflex_table_entry_delete_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry delete",
						},
						"sync_rx_aflex_table_entry_delete_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry delete",
						},
						"sync_tx_aflex_table_entry_incr_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry incr",
						},
						"sync_rx_aflex_table_entry_incr_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry incr",
						},
						"sync_tx_aflex_table_entry_lookup_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry lookup",
						},
						"sync_rx_aflex_table_entry_lookup_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry lookup",
						},
						"sync_tx_aflex_table_entry_lifetime_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry lifetime",
						},
						"sync_rx_aflex_table_entry_lifetime_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry lifetime",
						},
						"sync_tx_aflex_table_entry_replace_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry replace",
						},
						"sync_rx_aflex_table_entry_replace_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry replace",
						},
						"sync_tx_aflex_table_entry_set_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry set",
						},
						"sync_rx_aflex_table_entry_set_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry set",
						},
						"sync_tx_aflex_table_entry_timeout_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry timeout",
						},
						"sync_rx_aflex_table_entry_timeout_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry timeout",
						},
						"sync_tx_aflex_table_entry_fastsync_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync send packet with aflex table entry fast sync",
						},
						"sync_rx_aflex_table_entry_fastsync_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Sync receive packet with aflex table entry fast sync",
						},
						"sync_tx_aflex_table_entry_error_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Error on send packet with aflex table entry",
						},
						"sync_tx_aflex_table_entry_not_eligible_counter": {
							Type: schema.TypeInt, Optional: true, Description: "send of aflex table entry not eligible",
						},
						"sync_rx_ext_fw_limit_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Sync FW Limit Entry Info Failure",
						},
						"sync_tx_fw_set_dscp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync send fw set dscp counter",
						},
						"sync_rx_fw_set_dscp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Sync receive fw set dscp counter",
						},
						"dns_cache_sync_tx_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Create Sent counter",
						},
						"dns_cache_sync_rx_create_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Create Received counter",
						},
						"dns_cache_sync_tx_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Del Sent counter",
						},
						"dns_cache_sync_rx_del_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Del Received counter",
						},
						"dns_cache_sync_tx_frag_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Frag Sent counter",
						},
						"dns_cache_sync_rx_frag_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Frag Received counter",
						},
						"dns_cache_sync_tx_error_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Error Sent counter",
						},
						"dns_cache_sync_rx_error_counter": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Cache Sync Error Received counter",
						},
					},
				},
			},
		},
	}
}

func resourceVrrpAStateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAStateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAStateStatsStats := setObjectVrrpAStateStatsStats(res)
		d.Set("stats", VrrpAStateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAStateStatsStats(ret edpt.DataVrrpAStateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sync_pkt_tx_counter":                            ret.DtVrrpAStateStats.Stats.Sync_pkt_tx_counter,
			"sync_pkt_rcv_counter":                           ret.DtVrrpAStateStats.Stats.Sync_pkt_rcv_counter,
			"sync_rx_create_counter":                         ret.DtVrrpAStateStats.Stats.Sync_rx_create_counter,
			"sync_rx_del_counter":                            ret.DtVrrpAStateStats.Stats.Sync_rx_del_counter,
			"sync_rx_update_age_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_update_age_counter,
			"sync_tx_create_counter":                         ret.DtVrrpAStateStats.Stats.Sync_tx_create_counter,
			"sync_tx_del_counter":                            ret.DtVrrpAStateStats.Stats.Sync_tx_del_counter,
			"sync_tx_update_age_counter":                     ret.DtVrrpAStateStats.Stats.Sync_tx_update_age_counter,
			"sync_rx_persist_create_counter":                 ret.DtVrrpAStateStats.Stats.Sync_rx_persist_create_counter,
			"sync_rx_persist_del_counter":                    ret.DtVrrpAStateStats.Stats.Sync_rx_persist_del_counter,
			"sync_rx_persist_update_age_counter":             ret.DtVrrpAStateStats.Stats.Sync_rx_persist_update_age_counter,
			"sync_tx_persist_create_counter":                 ret.DtVrrpAStateStats.Stats.Sync_tx_persist_create_counter,
			"sync_tx_persist_del_counter":                    ret.DtVrrpAStateStats.Stats.Sync_tx_persist_del_counter,
			"sync_tx_persist_update_age_counter":             ret.DtVrrpAStateStats.Stats.Sync_tx_persist_update_age_counter,
			"query_pkt_tx_counter":                           ret.DtVrrpAStateStats.Stats.Query_pkt_tx_counter,
			"query_pkt_rcv_counter":                          ret.DtVrrpAStateStats.Stats.Query_pkt_rcv_counter,
			"sync_tx_smp_radius_table_counter":               ret.DtVrrpAStateStats.Stats.Sync_tx_smp_radius_table_counter,
			"sync_rx_smp_radius_table_counter":               ret.DtVrrpAStateStats.Stats.Sync_rx_smp_radius_table_counter,
			"query_tx_max_packed":                            ret.DtVrrpAStateStats.Stats.Query_tx_max_packed,
			"query_tx_min_packed":                            ret.DtVrrpAStateStats.Stats.Query_tx_min_packed,
			"query_pkt_invalid_idx_counter":                  ret.DtVrrpAStateStats.Stats.Query_pkt_invalid_idx_counter,
			"query_tx_get_buff_failed":                       ret.DtVrrpAStateStats.Stats.Query_tx_get_buff_failed,
			"query_rx_zero_info_counter":                     ret.DtVrrpAStateStats.Stats.Query_rx_zero_info_counter,
			"query_rx_full_info_counter":                     ret.DtVrrpAStateStats.Stats.Query_rx_full_info_counter,
			"query_rx_unk_counter":                           ret.DtVrrpAStateStats.Stats.Query_rx_unk_counter,
			"sync_pkt_invalid_idx_counter":                   ret.DtVrrpAStateStats.Stats.Sync_pkt_invalid_idx_counter,
			"sync_tx_get_buff_failed":                        ret.DtVrrpAStateStats.Stats.Sync_tx_get_buff_failed,
			"sync_tx_total_info_counter":                     ret.DtVrrpAStateStats.Stats.Sync_tx_total_info_counter,
			"sync_tx_create_ext_bit_counter":                 ret.DtVrrpAStateStats.Stats.Sync_tx_create_ext_bit_counter,
			"sync_tx_update_seqnos_counter":                  ret.DtVrrpAStateStats.Stats.Sync_tx_update_seqnos_counter,
			"sync_tx_min_packed":                             ret.DtVrrpAStateStats.Stats.Sync_tx_min_packed,
			"sync_tx_max_packed":                             ret.DtVrrpAStateStats.Stats.Sync_tx_max_packed,
			"sync_rx_len_invalid":                            ret.DtVrrpAStateStats.Stats.Sync_rx_len_invalid,
			"sync_persist_rx_len_invalid":                    ret.DtVrrpAStateStats.Stats.Sync_persist_rx_len_invalid,
			"sync_persist_rx_proto_not_supported":            ret.DtVrrpAStateStats.Stats.Sync_persist_rx_proto_not_supported,
			"sync_persist_rx_type_invalid":                   ret.DtVrrpAStateStats.Stats.Sync_persist_rx_type_invalid,
			"sync_persist_rx_cannot_process_mandatory":       ret.DtVrrpAStateStats.Stats.Sync_persist_rx_cannot_process_mandatory,
			"sync_persist_rx_ext_bit_process_error":          ret.DtVrrpAStateStats.Stats.Sync_persist_rx_ext_bit_process_error,
			"sync_persist_rx_no_such_vport":                  ret.DtVrrpAStateStats.Stats.Sync_persist_rx_no_such_vport,
			"sync_persist_rx_vporttype_not_supported":        ret.DtVrrpAStateStats.Stats.Sync_persist_rx_vporttype_not_supported,
			"sync_persist_rx_no_such_rport":                  ret.DtVrrpAStateStats.Stats.Sync_persist_rx_no_such_rport,
			"sync_persist_rx_no_such_sg_group":               ret.DtVrrpAStateStats.Stats.Sync_persist_rx_no_such_sg_group,
			"sync_persist_rx_no_sg_group_info":               ret.DtVrrpAStateStats.Stats.Sync_persist_rx_no_sg_group_info,
			"sync_persist_rx_conn_get_failed":                ret.DtVrrpAStateStats.Stats.Sync_persist_rx_conn_get_failed,
			"sync_rx_no_such_vport":                          ret.DtVrrpAStateStats.Stats.Sync_rx_no_such_vport,
			"sync_rx_no_such_rport":                          ret.DtVrrpAStateStats.Stats.Sync_rx_no_such_rport,
			"sync_rx_cannot_process_mandatory":               ret.DtVrrpAStateStats.Stats.Sync_rx_cannot_process_mandatory,
			"sync_rx_ext_bit_process_error":                  ret.DtVrrpAStateStats.Stats.Sync_rx_ext_bit_process_error,
			"sync_rx_create_ext_bit_counter":                 ret.DtVrrpAStateStats.Stats.Sync_rx_create_ext_bit_counter,
			"sync_rx_conn_exists":                            ret.DtVrrpAStateStats.Stats.Sync_rx_conn_exists,
			"sync_rx_conn_get_failed":                        ret.DtVrrpAStateStats.Stats.Sync_rx_conn_get_failed,
			"sync_rx_proto_not_supported":                    ret.DtVrrpAStateStats.Stats.Sync_rx_proto_not_supported,
			"sync_rx_no_dst_for_vport_inline":                ret.DtVrrpAStateStats.Stats.Sync_rx_no_dst_for_vport_inline,
			"sync_rx_no_such_nat_pool":                       ret.DtVrrpAStateStats.Stats.Sync_rx_no_such_nat_pool,
			"sync_rx_no_such_sg_node":                        ret.DtVrrpAStateStats.Stats.Sync_rx_no_such_sg_node,
			"sync_rx_del_no_such_session":                    ret.DtVrrpAStateStats.Stats.Sync_rx_del_no_such_session,
			"sync_rx_type_invalid":                           ret.DtVrrpAStateStats.Stats.Sync_rx_type_invalid,
			"sync_rx_zero_info_counter":                      ret.DtVrrpAStateStats.Stats.Sync_rx_zero_info_counter,
			"sync_rx_dcmsg_counter":                          ret.DtVrrpAStateStats.Stats.Sync_rx_dcmsg_counter,
			"sync_rx_total_info_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_total_info_counter,
			"sync_rx_update_seqnos_counter":                  ret.DtVrrpAStateStats.Stats.Sync_rx_update_seqnos_counter,
			"sync_rx_unk_counter":                            ret.DtVrrpAStateStats.Stats.Sync_rx_unk_counter,
			"sync_rx_apptype_not_supported":                  ret.DtVrrpAStateStats.Stats.Sync_rx_apptype_not_supported,
			"sync_query_dcmsg_counter":                       ret.DtVrrpAStateStats.Stats.Sync_query_dcmsg_counter,
			"sync_get_buff_failed_rt":                        ret.DtVrrpAStateStats.Stats.Sync_get_buff_failed_rt,
			"sync_get_buff_failed_port":                      ret.DtVrrpAStateStats.Stats.Sync_get_buff_failed_port,
			"sync_rx_lsn_create_sby":                         ret.DtVrrpAStateStats.Stats.Sync_rx_lsn_create_sby,
			"sync_rx_nat_create_sby":                         ret.DtVrrpAStateStats.Stats.Sync_rx_nat_create_sby,
			"sync_rx_nat_alloc_sby":                          ret.DtVrrpAStateStats.Stats.Sync_rx_nat_alloc_sby,
			"sync_rx_insert_tuple":                           ret.DtVrrpAStateStats.Stats.Sync_rx_insert_tuple,
			"sync_rx_sfw":                                    ret.DtVrrpAStateStats.Stats.Sync_rx_sfw,
			"sync_rx_create_static_sby":                      ret.DtVrrpAStateStats.Stats.Sync_rx_create_static_sby,
			"sync_rx_ext_pptp":                               ret.DtVrrpAStateStats.Stats.Sync_rx_ext_pptp,
			"sync_rx_ext_rtsp":                               ret.DtVrrpAStateStats.Stats.Sync_rx_ext_rtsp,
			"sync_rx_reserve_ha":                             ret.DtVrrpAStateStats.Stats.Sync_rx_reserve_ha,
			"sync_rx_seq_deltas":                             ret.DtVrrpAStateStats.Stats.Sync_rx_seq_deltas,
			"sync_rx_ftp_control":                            ret.DtVrrpAStateStats.Stats.Sync_rx_ftp_control,
			"sync_rx_ext_lsn_acl":                            ret.DtVrrpAStateStats.Stats.Sync_rx_ext_lsn_acl,
			"sync_rx_ext_lsn_ac_idle_timeout":                ret.DtVrrpAStateStats.Stats.Sync_rx_ext_lsn_ac_idle_timeout,
			"sync_rx_ext_sip_alg":                            ret.DtVrrpAStateStats.Stats.Sync_rx_ext_sip_alg,
			"sync_rx_ext_h323_alg":                           ret.DtVrrpAStateStats.Stats.Sync_rx_ext_h323_alg,
			"sync_rx_ext_nat_mac":                            ret.DtVrrpAStateStats.Stats.Sync_rx_ext_nat_mac,
			"sync_tx_lsn_fullcone":                           ret.DtVrrpAStateStats.Stats.Sync_tx_lsn_fullcone,
			"sync_rx_lsn_fullcone":                           ret.DtVrrpAStateStats.Stats.Sync_rx_lsn_fullcone,
			"sync_err_lsn_fullcone":                          ret.DtVrrpAStateStats.Stats.Sync_err_lsn_fullcone,
			"sync_tx_update_sctp_conn_addr":                  ret.DtVrrpAStateStats.Stats.Sync_tx_update_sctp_conn_addr,
			"sync_rx_update_sctp_conn_addr":                  ret.DtVrrpAStateStats.Stats.Sync_rx_update_sctp_conn_addr,
			"sync_rx_ext_nat_alg_tcp_info":                   ret.DtVrrpAStateStats.Stats.Sync_rx_ext_nat_alg_tcp_info,
			"sync_rx_ext_dcfw_rule_id":                       ret.DtVrrpAStateStats.Stats.Sync_rx_ext_dcfw_rule_id,
			"sync_rx_ext_dcfw_log":                           ret.DtVrrpAStateStats.Stats.Sync_rx_ext_dcfw_log,
			"sync_rx_estab_counter":                          ret.DtVrrpAStateStats.Stats.Sync_rx_estab_counter,
			"sync_tx_estab_counter":                          ret.DtVrrpAStateStats.Stats.Sync_tx_estab_counter,
			"sync_rx_zone_failure_counter":                   ret.DtVrrpAStateStats.Stats.Sync_rx_zone_failure_counter,
			"sync_rx_ext_fw_http_logging":                    ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_http_logging,
			"sync_rx_ext_dcfw_rule_idle_timeout":             ret.DtVrrpAStateStats.Stats.Sync_rx_ext_dcfw_rule_idle_timeout,
			"sync_rx_ext_fw_gtp_info":                        ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_info,
			"sync_rx_not_expect_sync_pkt":                    ret.DtVrrpAStateStats.Stats.Sync_rx_not_expect_sync_pkt,
			"sync_rx_ext_fw_apps":                            ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_apps,
			"sync_tx_mon_entity":                             ret.DtVrrpAStateStats.Stats.Sync_tx_mon_entity,
			"sync_rx_mon_entity":                             ret.DtVrrpAStateStats.Stats.Sync_rx_mon_entity,
			"sync_rx_ext_fw_gtp_log_info":                    ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_log_info,
			"sync_rx_ext_fw_gtp_u_info":                      ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_u_info,
			"sync_rx_ext_fw_gtp_ext_info":                    ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_ext_info,
			"sync_rx_ext_fw_gtp_log_ext_info":                ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_log_ext_info,
			"sync_rx_ddos_drop_counter":                      ret.DtVrrpAStateStats.Stats.Sync_rx_ddos_drop_counter,
			"sync_rx_invalid_sync_packet_counter":            ret.DtVrrpAStateStats.Stats.Sync_rx_invalid_sync_packet_counter,
			"sync_pkt_empty_buff_counter":                    ret.DtVrrpAStateStats.Stats.Sync_pkt_empty_buff_counter,
			"sync_pkt_no_sending_vgrp_counter":               ret.DtVrrpAStateStats.Stats.Sync_pkt_no_sending_vgrp_counter,
			"sync_pkt_no_receiving_vgrp_counter":             ret.DtVrrpAStateStats.Stats.Sync_pkt_no_receiving_vgrp_counter,
			"query_pkt_no_receiving_ip_counter":              ret.DtVrrpAStateStats.Stats.Query_pkt_no_receiving_ip_counter,
			"sync_pkt_failed_buff_copy_counter":              ret.DtVrrpAStateStats.Stats.Sync_pkt_failed_buff_copy_counter,
			"sync_rx_bad_protocol_counter":                   ret.DtVrrpAStateStats.Stats.Sync_rx_bad_protocol_counter,
			"sync_rx_no_vgrp_counter":                        ret.DtVrrpAStateStats.Stats.Sync_rx_no_vgrp_counter,
			"sync_rx_by_inactive_peer_counter":               ret.DtVrrpAStateStats.Stats.Sync_rx_by_inactive_peer_counter,
			"sync_rx_table_entry_update_counter":             ret.DtVrrpAStateStats.Stats.Sync_rx_table_entry_update_counter,
			"sync_rx_table_entry_create_counter":             ret.DtVrrpAStateStats.Stats.Sync_rx_table_entry_create_counter,
			"sync_rx_table_entry_del_counter":                ret.DtVrrpAStateStats.Stats.Sync_rx_table_entry_del_counter,
			"sync_rx_aflex_update_counter":                   ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_update_counter,
			"sync_rx_aflex_create_counter":                   ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_create_counter,
			"sync_rx_aflex_del_counter":                      ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_del_counter,
			"sync_rx_aflex_frag_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_frag_counter,
			"query_rx_invalid_partition_counter":             ret.DtVrrpAStateStats.Stats.Query_rx_invalid_partition_counter,
			"query_rx_invalid_ha_group_counter":              ret.DtVrrpAStateStats.Stats.Query_rx_invalid_ha_group_counter,
			"query_rx_invalid_sync_version_counter":          ret.DtVrrpAStateStats.Stats.Query_rx_invalid_sync_version_counter,
			"query_rx_invalid_msg_dir_counter":               ret.DtVrrpAStateStats.Stats.Query_rx_invalid_msg_dir_counter,
			"sync_rx_out_of_order_pkt_counter":               ret.DtVrrpAStateStats.Stats.Sync_rx_out_of_order_pkt_counter,
			"sync_rx_unreached_pkt_counter":                  ret.DtVrrpAStateStats.Stats.Sync_rx_unreached_pkt_counter,
			"sync_rx_ext_fw_gtp_echo_ext_info":               ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_gtp_echo_ext_info,
			"sync_rx_smp_create_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_smp_create_counter,
			"sync_rx_smp_delete_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_smp_delete_counter,
			"sync_rx_smp_update_counter":                     ret.DtVrrpAStateStats.Stats.Sync_rx_smp_update_counter,
			"sync_tx_smp_create_counter":                     ret.DtVrrpAStateStats.Stats.Sync_tx_smp_create_counter,
			"sync_tx_smp_delete_counter":                     ret.DtVrrpAStateStats.Stats.Sync_tx_smp_delete_counter,
			"sync_tx_smp_update_counter":                     ret.DtVrrpAStateStats.Stats.Sync_tx_smp_update_counter,
			"sync_rx_smp_clear_counter":                      ret.DtVrrpAStateStats.Stats.Sync_rx_smp_clear_counter,
			"sync_tx_smp_clear_counter":                      ret.DtVrrpAStateStats.Stats.Sync_tx_smp_clear_counter,
			"sync_rx_ext_fw_so_shadow_ext_info":              ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_so_shadow_ext_info,
			"sync_tx_aflex_table_entry_add_counter":          ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_add_counter,
			"sync_rx_aflex_table_entry_add_counter":          ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_add_counter,
			"sync_tx_aflex_table_entry_append_counter":       ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_append_counter,
			"sync_rx_aflex_table_entry_append_counter":       ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_append_counter,
			"sync_tx_aflex_table_entry_delete_counter":       ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_delete_counter,
			"sync_rx_aflex_table_entry_delete_counter":       ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_delete_counter,
			"sync_tx_aflex_table_entry_incr_counter":         ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_incr_counter,
			"sync_rx_aflex_table_entry_incr_counter":         ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_incr_counter,
			"sync_tx_aflex_table_entry_lookup_counter":       ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_lookup_counter,
			"sync_rx_aflex_table_entry_lookup_counter":       ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_lookup_counter,
			"sync_tx_aflex_table_entry_lifetime_counter":     ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_lifetime_counter,
			"sync_rx_aflex_table_entry_lifetime_counter":     ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_lifetime_counter,
			"sync_tx_aflex_table_entry_replace_counter":      ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_replace_counter,
			"sync_rx_aflex_table_entry_replace_counter":      ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_replace_counter,
			"sync_tx_aflex_table_entry_set_counter":          ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_set_counter,
			"sync_rx_aflex_table_entry_set_counter":          ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_set_counter,
			"sync_tx_aflex_table_entry_timeout_counter":      ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_timeout_counter,
			"sync_rx_aflex_table_entry_timeout_counter":      ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_timeout_counter,
			"sync_tx_aflex_table_entry_fastsync_counter":     ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_fastsync_counter,
			"sync_rx_aflex_table_entry_fastsync_counter":     ret.DtVrrpAStateStats.Stats.Sync_rx_aflex_table_entry_fastsync_counter,
			"sync_tx_aflex_table_entry_error_counter":        ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_error_counter,
			"sync_tx_aflex_table_entry_not_eligible_counter": ret.DtVrrpAStateStats.Stats.Sync_tx_aflex_table_entry_not_eligible_counter,
			"sync_rx_ext_fw_limit_entry":                     ret.DtVrrpAStateStats.Stats.Sync_rx_ext_fw_limit_entry,
			"sync_tx_fw_set_dscp_counter":                    ret.DtVrrpAStateStats.Stats.Sync_tx_fw_set_dscp_counter,
			"sync_rx_fw_set_dscp_counter":                    ret.DtVrrpAStateStats.Stats.Sync_rx_fw_set_dscp_counter,
			"dns_cache_sync_tx_create_counter":               ret.DtVrrpAStateStats.Stats.Dns_cache_sync_tx_create_counter,
			"dns_cache_sync_rx_create_counter":               ret.DtVrrpAStateStats.Stats.Dns_cache_sync_rx_create_counter,
			"dns_cache_sync_tx_del_counter":                  ret.DtVrrpAStateStats.Stats.Dns_cache_sync_tx_del_counter,
			"dns_cache_sync_rx_del_counter":                  ret.DtVrrpAStateStats.Stats.Dns_cache_sync_rx_del_counter,
			"dns_cache_sync_tx_frag_counter":                 ret.DtVrrpAStateStats.Stats.Dns_cache_sync_tx_frag_counter,
			"dns_cache_sync_rx_frag_counter":                 ret.DtVrrpAStateStats.Stats.Dns_cache_sync_rx_frag_counter,
			"dns_cache_sync_tx_error_counter":                ret.DtVrrpAStateStats.Stats.Dns_cache_sync_tx_error_counter,
			"dns_cache_sync_rx_error_counter":                ret.DtVrrpAStateStats.Stats.Dns_cache_sync_rx_error_counter,
		},
	}
}

func getObjectVrrpAStateStatsStats(d []interface{}) edpt.VrrpAStateStatsStats {

	count1 := len(d)
	var ret edpt.VrrpAStateStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sync_pkt_tx_counter = in["sync_pkt_tx_counter"].(int)
		ret.Sync_pkt_rcv_counter = in["sync_pkt_rcv_counter"].(int)
		ret.Sync_rx_create_counter = in["sync_rx_create_counter"].(int)
		ret.Sync_rx_del_counter = in["sync_rx_del_counter"].(int)
		ret.Sync_rx_update_age_counter = in["sync_rx_update_age_counter"].(int)
		ret.Sync_tx_create_counter = in["sync_tx_create_counter"].(int)
		ret.Sync_tx_del_counter = in["sync_tx_del_counter"].(int)
		ret.Sync_tx_update_age_counter = in["sync_tx_update_age_counter"].(int)
		ret.Sync_rx_persist_create_counter = in["sync_rx_persist_create_counter"].(int)
		ret.Sync_rx_persist_del_counter = in["sync_rx_persist_del_counter"].(int)
		ret.Sync_rx_persist_update_age_counter = in["sync_rx_persist_update_age_counter"].(int)
		ret.Sync_tx_persist_create_counter = in["sync_tx_persist_create_counter"].(int)
		ret.Sync_tx_persist_del_counter = in["sync_tx_persist_del_counter"].(int)
		ret.Sync_tx_persist_update_age_counter = in["sync_tx_persist_update_age_counter"].(int)
		ret.Query_pkt_tx_counter = in["query_pkt_tx_counter"].(int)
		ret.Query_pkt_rcv_counter = in["query_pkt_rcv_counter"].(int)
		ret.Sync_tx_smp_radius_table_counter = in["sync_tx_smp_radius_table_counter"].(int)
		ret.Sync_rx_smp_radius_table_counter = in["sync_rx_smp_radius_table_counter"].(int)
		ret.Query_tx_max_packed = in["query_tx_max_packed"].(int)
		ret.Query_tx_min_packed = in["query_tx_min_packed"].(int)
		ret.Query_pkt_invalid_idx_counter = in["query_pkt_invalid_idx_counter"].(int)
		ret.Query_tx_get_buff_failed = in["query_tx_get_buff_failed"].(int)
		ret.Query_rx_zero_info_counter = in["query_rx_zero_info_counter"].(int)
		ret.Query_rx_full_info_counter = in["query_rx_full_info_counter"].(int)
		ret.Query_rx_unk_counter = in["query_rx_unk_counter"].(int)
		ret.Sync_pkt_invalid_idx_counter = in["sync_pkt_invalid_idx_counter"].(int)
		ret.Sync_tx_get_buff_failed = in["sync_tx_get_buff_failed"].(int)
		ret.Sync_tx_total_info_counter = in["sync_tx_total_info_counter"].(int)
		ret.Sync_tx_create_ext_bit_counter = in["sync_tx_create_ext_bit_counter"].(int)
		ret.Sync_tx_update_seqnos_counter = in["sync_tx_update_seqnos_counter"].(int)
		ret.Sync_tx_min_packed = in["sync_tx_min_packed"].(int)
		ret.Sync_tx_max_packed = in["sync_tx_max_packed"].(int)
		ret.Sync_rx_len_invalid = in["sync_rx_len_invalid"].(int)
		ret.Sync_persist_rx_len_invalid = in["sync_persist_rx_len_invalid"].(int)
		ret.Sync_persist_rx_proto_not_supported = in["sync_persist_rx_proto_not_supported"].(int)
		ret.Sync_persist_rx_type_invalid = in["sync_persist_rx_type_invalid"].(int)
		ret.Sync_persist_rx_cannot_process_mandatory = in["sync_persist_rx_cannot_process_mandatory"].(int)
		ret.Sync_persist_rx_ext_bit_process_error = in["sync_persist_rx_ext_bit_process_error"].(int)
		ret.Sync_persist_rx_no_such_vport = in["sync_persist_rx_no_such_vport"].(int)
		ret.Sync_persist_rx_vporttype_not_supported = in["sync_persist_rx_vporttype_not_supported"].(int)
		ret.Sync_persist_rx_no_such_rport = in["sync_persist_rx_no_such_rport"].(int)
		ret.Sync_persist_rx_no_such_sg_group = in["sync_persist_rx_no_such_sg_group"].(int)
		ret.Sync_persist_rx_no_sg_group_info = in["sync_persist_rx_no_sg_group_info"].(int)
		ret.Sync_persist_rx_conn_get_failed = in["sync_persist_rx_conn_get_failed"].(int)
		ret.Sync_rx_no_such_vport = in["sync_rx_no_such_vport"].(int)
		ret.Sync_rx_no_such_rport = in["sync_rx_no_such_rport"].(int)
		ret.Sync_rx_cannot_process_mandatory = in["sync_rx_cannot_process_mandatory"].(int)
		ret.Sync_rx_ext_bit_process_error = in["sync_rx_ext_bit_process_error"].(int)
		ret.Sync_rx_create_ext_bit_counter = in["sync_rx_create_ext_bit_counter"].(int)
		ret.Sync_rx_conn_exists = in["sync_rx_conn_exists"].(int)
		ret.Sync_rx_conn_get_failed = in["sync_rx_conn_get_failed"].(int)
		ret.Sync_rx_proto_not_supported = in["sync_rx_proto_not_supported"].(int)
		ret.Sync_rx_no_dst_for_vport_inline = in["sync_rx_no_dst_for_vport_inline"].(int)
		ret.Sync_rx_no_such_nat_pool = in["sync_rx_no_such_nat_pool"].(int)
		ret.Sync_rx_no_such_sg_node = in["sync_rx_no_such_sg_node"].(int)
		ret.Sync_rx_del_no_such_session = in["sync_rx_del_no_such_session"].(int)
		ret.Sync_rx_type_invalid = in["sync_rx_type_invalid"].(int)
		ret.Sync_rx_zero_info_counter = in["sync_rx_zero_info_counter"].(int)
		ret.Sync_rx_dcmsg_counter = in["sync_rx_dcmsg_counter"].(int)
		ret.Sync_rx_total_info_counter = in["sync_rx_total_info_counter"].(int)
		ret.Sync_rx_update_seqnos_counter = in["sync_rx_update_seqnos_counter"].(int)
		ret.Sync_rx_unk_counter = in["sync_rx_unk_counter"].(int)
		ret.Sync_rx_apptype_not_supported = in["sync_rx_apptype_not_supported"].(int)
		ret.Sync_query_dcmsg_counter = in["sync_query_dcmsg_counter"].(int)
		ret.Sync_get_buff_failed_rt = in["sync_get_buff_failed_rt"].(int)
		ret.Sync_get_buff_failed_port = in["sync_get_buff_failed_port"].(int)
		ret.Sync_rx_lsn_create_sby = in["sync_rx_lsn_create_sby"].(int)
		ret.Sync_rx_nat_create_sby = in["sync_rx_nat_create_sby"].(int)
		ret.Sync_rx_nat_alloc_sby = in["sync_rx_nat_alloc_sby"].(int)
		ret.Sync_rx_insert_tuple = in["sync_rx_insert_tuple"].(int)
		ret.Sync_rx_sfw = in["sync_rx_sfw"].(int)
		ret.Sync_rx_create_static_sby = in["sync_rx_create_static_sby"].(int)
		ret.Sync_rx_ext_pptp = in["sync_rx_ext_pptp"].(int)
		ret.Sync_rx_ext_rtsp = in["sync_rx_ext_rtsp"].(int)
		ret.Sync_rx_reserve_ha = in["sync_rx_reserve_ha"].(int)
		ret.Sync_rx_seq_deltas = in["sync_rx_seq_deltas"].(int)
		ret.Sync_rx_ftp_control = in["sync_rx_ftp_control"].(int)
		ret.Sync_rx_ext_lsn_acl = in["sync_rx_ext_lsn_acl"].(int)
		ret.Sync_rx_ext_lsn_ac_idle_timeout = in["sync_rx_ext_lsn_ac_idle_timeout"].(int)
		ret.Sync_rx_ext_sip_alg = in["sync_rx_ext_sip_alg"].(int)
		ret.Sync_rx_ext_h323_alg = in["sync_rx_ext_h323_alg"].(int)
		ret.Sync_rx_ext_nat_mac = in["sync_rx_ext_nat_mac"].(int)
		ret.Sync_tx_lsn_fullcone = in["sync_tx_lsn_fullcone"].(int)
		ret.Sync_rx_lsn_fullcone = in["sync_rx_lsn_fullcone"].(int)
		ret.Sync_err_lsn_fullcone = in["sync_err_lsn_fullcone"].(int)
		ret.Sync_tx_update_sctp_conn_addr = in["sync_tx_update_sctp_conn_addr"].(int)
		ret.Sync_rx_update_sctp_conn_addr = in["sync_rx_update_sctp_conn_addr"].(int)
		ret.Sync_rx_ext_nat_alg_tcp_info = in["sync_rx_ext_nat_alg_tcp_info"].(int)
		ret.Sync_rx_ext_dcfw_rule_id = in["sync_rx_ext_dcfw_rule_id"].(int)
		ret.Sync_rx_ext_dcfw_log = in["sync_rx_ext_dcfw_log"].(int)
		ret.Sync_rx_estab_counter = in["sync_rx_estab_counter"].(int)
		ret.Sync_tx_estab_counter = in["sync_tx_estab_counter"].(int)
		ret.Sync_rx_zone_failure_counter = in["sync_rx_zone_failure_counter"].(int)
		ret.Sync_rx_ext_fw_http_logging = in["sync_rx_ext_fw_http_logging"].(int)
		ret.Sync_rx_ext_dcfw_rule_idle_timeout = in["sync_rx_ext_dcfw_rule_idle_timeout"].(int)
		ret.Sync_rx_ext_fw_gtp_info = in["sync_rx_ext_fw_gtp_info"].(int)
		ret.Sync_rx_not_expect_sync_pkt = in["sync_rx_not_expect_sync_pkt"].(int)
		ret.Sync_rx_ext_fw_apps = in["sync_rx_ext_fw_apps"].(int)
		ret.Sync_tx_mon_entity = in["sync_tx_mon_entity"].(int)
		ret.Sync_rx_mon_entity = in["sync_rx_mon_entity"].(int)
		ret.Sync_rx_ext_fw_gtp_log_info = in["sync_rx_ext_fw_gtp_log_info"].(int)
		ret.Sync_rx_ext_fw_gtp_u_info = in["sync_rx_ext_fw_gtp_u_info"].(int)
		ret.Sync_rx_ext_fw_gtp_ext_info = in["sync_rx_ext_fw_gtp_ext_info"].(int)
		ret.Sync_rx_ext_fw_gtp_log_ext_info = in["sync_rx_ext_fw_gtp_log_ext_info"].(int)
		ret.Sync_rx_ddos_drop_counter = in["sync_rx_ddos_drop_counter"].(int)
		ret.Sync_rx_invalid_sync_packet_counter = in["sync_rx_invalid_sync_packet_counter"].(int)
		ret.Sync_pkt_empty_buff_counter = in["sync_pkt_empty_buff_counter"].(int)
		ret.Sync_pkt_no_sending_vgrp_counter = in["sync_pkt_no_sending_vgrp_counter"].(int)
		ret.Sync_pkt_no_receiving_vgrp_counter = in["sync_pkt_no_receiving_vgrp_counter"].(int)
		ret.Query_pkt_no_receiving_ip_counter = in["query_pkt_no_receiving_ip_counter"].(int)
		ret.Sync_pkt_failed_buff_copy_counter = in["sync_pkt_failed_buff_copy_counter"].(int)
		ret.Sync_rx_bad_protocol_counter = in["sync_rx_bad_protocol_counter"].(int)
		ret.Sync_rx_no_vgrp_counter = in["sync_rx_no_vgrp_counter"].(int)
		ret.Sync_rx_by_inactive_peer_counter = in["sync_rx_by_inactive_peer_counter"].(int)
		ret.Sync_rx_table_entry_update_counter = in["sync_rx_table_entry_update_counter"].(int)
		ret.Sync_rx_table_entry_create_counter = in["sync_rx_table_entry_create_counter"].(int)
		ret.Sync_rx_table_entry_del_counter = in["sync_rx_table_entry_del_counter"].(int)
		ret.Sync_rx_aflex_update_counter = in["sync_rx_aflex_update_counter"].(int)
		ret.Sync_rx_aflex_create_counter = in["sync_rx_aflex_create_counter"].(int)
		ret.Sync_rx_aflex_del_counter = in["sync_rx_aflex_del_counter"].(int)
		ret.Sync_rx_aflex_frag_counter = in["sync_rx_aflex_frag_counter"].(int)
		ret.Query_rx_invalid_partition_counter = in["query_rx_invalid_partition_counter"].(int)
		ret.Query_rx_invalid_ha_group_counter = in["query_rx_invalid_ha_group_counter"].(int)
		ret.Query_rx_invalid_sync_version_counter = in["query_rx_invalid_sync_version_counter"].(int)
		ret.Query_rx_invalid_msg_dir_counter = in["query_rx_invalid_msg_dir_counter"].(int)
		ret.Sync_rx_out_of_order_pkt_counter = in["sync_rx_out_of_order_pkt_counter"].(int)
		ret.Sync_rx_unreached_pkt_counter = in["sync_rx_unreached_pkt_counter"].(int)
		ret.Sync_rx_ext_fw_gtp_echo_ext_info = in["sync_rx_ext_fw_gtp_echo_ext_info"].(int)
		ret.Sync_rx_smp_create_counter = in["sync_rx_smp_create_counter"].(int)
		ret.Sync_rx_smp_delete_counter = in["sync_rx_smp_delete_counter"].(int)
		ret.Sync_rx_smp_update_counter = in["sync_rx_smp_update_counter"].(int)
		ret.Sync_tx_smp_create_counter = in["sync_tx_smp_create_counter"].(int)
		ret.Sync_tx_smp_delete_counter = in["sync_tx_smp_delete_counter"].(int)
		ret.Sync_tx_smp_update_counter = in["sync_tx_smp_update_counter"].(int)
		ret.Sync_rx_smp_clear_counter = in["sync_rx_smp_clear_counter"].(int)
		ret.Sync_tx_smp_clear_counter = in["sync_tx_smp_clear_counter"].(int)
		ret.Sync_rx_ext_fw_so_shadow_ext_info = in["sync_rx_ext_fw_so_shadow_ext_info"].(int)
		ret.Sync_tx_aflex_table_entry_add_counter = in["sync_tx_aflex_table_entry_add_counter"].(int)
		ret.Sync_rx_aflex_table_entry_add_counter = in["sync_rx_aflex_table_entry_add_counter"].(int)
		ret.Sync_tx_aflex_table_entry_append_counter = in["sync_tx_aflex_table_entry_append_counter"].(int)
		ret.Sync_rx_aflex_table_entry_append_counter = in["sync_rx_aflex_table_entry_append_counter"].(int)
		ret.Sync_tx_aflex_table_entry_delete_counter = in["sync_tx_aflex_table_entry_delete_counter"].(int)
		ret.Sync_rx_aflex_table_entry_delete_counter = in["sync_rx_aflex_table_entry_delete_counter"].(int)
		ret.Sync_tx_aflex_table_entry_incr_counter = in["sync_tx_aflex_table_entry_incr_counter"].(int)
		ret.Sync_rx_aflex_table_entry_incr_counter = in["sync_rx_aflex_table_entry_incr_counter"].(int)
		ret.Sync_tx_aflex_table_entry_lookup_counter = in["sync_tx_aflex_table_entry_lookup_counter"].(int)
		ret.Sync_rx_aflex_table_entry_lookup_counter = in["sync_rx_aflex_table_entry_lookup_counter"].(int)
		ret.Sync_tx_aflex_table_entry_lifetime_counter = in["sync_tx_aflex_table_entry_lifetime_counter"].(int)
		ret.Sync_rx_aflex_table_entry_lifetime_counter = in["sync_rx_aflex_table_entry_lifetime_counter"].(int)
		ret.Sync_tx_aflex_table_entry_replace_counter = in["sync_tx_aflex_table_entry_replace_counter"].(int)
		ret.Sync_rx_aflex_table_entry_replace_counter = in["sync_rx_aflex_table_entry_replace_counter"].(int)
		ret.Sync_tx_aflex_table_entry_set_counter = in["sync_tx_aflex_table_entry_set_counter"].(int)
		ret.Sync_rx_aflex_table_entry_set_counter = in["sync_rx_aflex_table_entry_set_counter"].(int)
		ret.Sync_tx_aflex_table_entry_timeout_counter = in["sync_tx_aflex_table_entry_timeout_counter"].(int)
		ret.Sync_rx_aflex_table_entry_timeout_counter = in["sync_rx_aflex_table_entry_timeout_counter"].(int)
		ret.Sync_tx_aflex_table_entry_fastsync_counter = in["sync_tx_aflex_table_entry_fastsync_counter"].(int)
		ret.Sync_rx_aflex_table_entry_fastsync_counter = in["sync_rx_aflex_table_entry_fastsync_counter"].(int)
		ret.Sync_tx_aflex_table_entry_error_counter = in["sync_tx_aflex_table_entry_error_counter"].(int)
		ret.Sync_tx_aflex_table_entry_not_eligible_counter = in["sync_tx_aflex_table_entry_not_eligible_counter"].(int)
		ret.Sync_rx_ext_fw_limit_entry = in["sync_rx_ext_fw_limit_entry"].(int)
		ret.Sync_tx_fw_set_dscp_counter = in["sync_tx_fw_set_dscp_counter"].(int)
		ret.Sync_rx_fw_set_dscp_counter = in["sync_rx_fw_set_dscp_counter"].(int)
		ret.Dns_cache_sync_tx_create_counter = in["dns_cache_sync_tx_create_counter"].(int)
		ret.Dns_cache_sync_rx_create_counter = in["dns_cache_sync_rx_create_counter"].(int)
		ret.Dns_cache_sync_tx_del_counter = in["dns_cache_sync_tx_del_counter"].(int)
		ret.Dns_cache_sync_rx_del_counter = in["dns_cache_sync_rx_del_counter"].(int)
		ret.Dns_cache_sync_tx_frag_counter = in["dns_cache_sync_tx_frag_counter"].(int)
		ret.Dns_cache_sync_rx_frag_counter = in["dns_cache_sync_rx_frag_counter"].(int)
		ret.Dns_cache_sync_tx_error_counter = in["dns_cache_sync_tx_error_counter"].(int)
		ret.Dns_cache_sync_rx_error_counter = in["dns_cache_sync_rx_error_counter"].(int)
	}
	return ret
}

func dataToEndpointVrrpAStateStats(d *schema.ResourceData) edpt.VrrpAStateStats {
	var ret edpt.VrrpAStateStats

	ret.Stats = getObjectVrrpAStateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
