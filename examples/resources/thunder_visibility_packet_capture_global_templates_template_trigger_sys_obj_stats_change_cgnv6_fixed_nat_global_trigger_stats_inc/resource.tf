provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global_trigger_stats_inc" {

  name                                    = "test"
  config_not_found                        = 1
  dest_rlist_drop                         = 1
  dest_rlist_pass_through                 = 1
  dest_rlist_snat_drop                    = 1
  dslite_eif_limit_exceeded               = 1
  dslite_inbound_filtered                 = 1
  fixed_nat_fullcone_self_hairpinning_dro = 1
  fullcone_failure                        = 1
  ha_session_user_quota_exceeded          = 1
  nat_port_unavailable_icmp               = 1
  nat_port_unavailable_tcp                = 1
  nat_port_unavailable_udp                = 1
  nat44_eif_limit_exceeded                = 1
  nat44_inbound_filtered                  = 1
  nat64_eif_limit_exceeded                = 1
  nat64_inbound_filtered                  = 1
  port_overload_failed                    = 1
  session_user_quota_exceeded             = 1
  sixrd_drop                              = 1
  standby_drop                            = 1
}