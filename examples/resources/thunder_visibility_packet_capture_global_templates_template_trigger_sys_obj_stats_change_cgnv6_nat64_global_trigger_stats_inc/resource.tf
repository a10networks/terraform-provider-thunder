provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_nat64_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_nat64_global_trigger_stats_inc" {

  name                            = "test"
  eif_limit_exceeded              = 1
  fullcone_failure                = 1
  fullcone_self_hairpinning_drop  = 1
  ha_nat_pool_batch_type_mismatch = 1
  ha_nat_pool_unusable            = 1
  nat_pool_unusable               = 1
  nat_port_unavailable_icmp       = 1
  nat_port_unavailable_tcp        = 1
  nat_port_unavailable_udp        = 1
  new_user_resource_unavailable   = 1
  no_class_list_match             = 1
  no_radius_profile_match         = 1
  user_quota_failure              = 1
  user_quota_unusable             = 1
  user_quota_unusable_drop        = 1
}