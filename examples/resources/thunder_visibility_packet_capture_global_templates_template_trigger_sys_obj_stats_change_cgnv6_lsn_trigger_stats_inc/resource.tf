provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_trigger_stats_inc" {

  name                                    = "test"
  adc_port_allocation_failed              = 1
  data_sesn_user_quota_exceeded           = 1
  fullcone_ext_mem_alloc_failure          = 1
  fullcone_ext_mem_alloc_init_faulure     = 1
  fullcone_failure                        = 1
  fullcone_self_hairpinning_drop          = 1
  h323_alg_alloc_single_port_failure      = 1
  h323_alg_create_rtcp_fullcone_failure   = 1
  h323_alg_create_rtp_fullcone_failure    = 1
  h323_alg_create_single_fullcone_failure = 1
  ha_nat_pool_batch_type_mismatch         = 1
  ha_nat_pool_unusable                    = 1
  mgcp_alg_create_rtcp_fullcone_failure   = 1
  mgcp_alg_create_rtp_fullcone_failure    = 1
  mgcp_alg_port_pair_alloc_from_quota_par = 1
  nat_pool_unusable                       = 1
  port_overloading_inc_overflow           = 1
  port_overloading_out_of_memory          = 1
  sip_alg_alloc_rtp_rtcp_port_failure     = 1
  sip_alg_alloc_single_port_failure       = 1
  sip_alg_create_rtcp_fullcone_failure    = 1
  sip_alg_create_rtp_fullcone_failure     = 1
  sip_alg_create_single_fullcone_failure  = 1
  sip_alg_quota_inc_failure               = 1
  user_quota_failure                      = 1
  user_quota_unusable                     = 1
  user_quota_unusable_drop                = 1
}