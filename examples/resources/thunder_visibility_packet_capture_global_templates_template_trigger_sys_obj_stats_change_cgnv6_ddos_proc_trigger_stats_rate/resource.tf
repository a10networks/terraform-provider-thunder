provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ddos_proc_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ddos_proc_trigger_stats_rate" {

  name                             = "test"
  duration                         = 60
  ip_node_alloc_failure            = 1
  ip_other_block_alloc_failure     = 1
  ip_port_block_alloc_failure      = 1
  l3_entry_add_to_bgp_failure      = 1
  l3_entry_add_to_hw_failure       = 1
  l3_entry_drop_max_hw_exceeded    = 1
  l3_entry_match_drop              = 1
  l3_entry_match_drop_hw           = 1
  l3_entry_remove_from_bgp_failure = 1
  l4_entry_drop_max_hw_exceeded    = 1
  l4_entry_list_alloc_failure      = 1
  l4_entry_match_drop              = 1
  l4_entry_match_drop_hw           = 1
  syn_cookie_verification_failed   = 1
  threshold_exceeded_by            = 5
}