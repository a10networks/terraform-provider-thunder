provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_so_counters_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_so_counters_trigger_stats_rate" {

  name                                    = "test"
  duration                                = 60
  so_pkts_l2redirect_dest_mac_zero_drop   = 1
  so_pkts_l2redirect_interface_not_up     = 1
  so_pkts_l2redirect_invalid_redirect_inf = 1
  so_pkts_l2redirect_port_retrieval_error = 1
  so_pkts_l2redirect_vlan_retrieval_error = 1
  so_pkts_l3_redirect_chassis_dest_mac_er = 1
  so_pkts_l3_redirect_encap_error_drop    = 1
  so_pkts_l3_redirect_fragmentation_error = 1
  so_pkts_l3_redirect_inner_mac_zero_drop = 1
  so_pkts_l3_redirect_invalid_dev_dir     = 1
  so_pkts_l3_redirect_table_error         = 1
  so_pkts_l3_redirect_table_no_entry_foun = 1
  so_pkts_slb_nat_release_fail            = 1
  so_pkts_slb_nat_reserve_fail            = 1
  threshold_exceeded_by                   = 5
}