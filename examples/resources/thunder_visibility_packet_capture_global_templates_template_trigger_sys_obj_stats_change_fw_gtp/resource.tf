provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_gtp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_gtp" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by                   = 5
    duration                                = 60
    out_of_session_memory                   = 1
    gtp_smp_path_check_failed               = 1
    gtp_smp_check_failed                    = 1
    gtp_smp_session_count_check_failed      = 1
    gtp_c_ref_count_smp_exceeded            = 1
    gtp_u_smp_in_rml_with_sess              = 1
    gtp_tunnel_rate_limit_entry_create_fail = 1
    gtp_rate_limit_smp_create_failure       = 1
    gtp_rate_limit_t3_ctr_create_failure    = 1
    gtp_rate_limit_entry_create_failure     = 1
    gtp_smp_dec_sess_count_check_failed     = 1
  }
}