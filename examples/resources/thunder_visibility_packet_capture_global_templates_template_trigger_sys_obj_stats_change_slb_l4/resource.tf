provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l4" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l4" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by   = 5
    duration                = 60
    syncookiessentfailed    = 1
    svrselfail              = 1
    snat_fail               = 1
    snat_no_fwd_route       = 1
    snat_no_rev_route       = 1
    snat_icmp_error_process = 1
    snat_icmp_no_match      = 1
    smart_nat_id_mismatch   = 1
    syncookiescheckfailed   = 1
    connlimit_drop          = 1
    conn_rate_limit_drop    = 1
    conn_rate_limit_reset   = 1
    dns_policy_drop         = 1
    no_resourse_drop        = 1
    bw_rate_limit_exceed    = 1
    l4_cps_exceed           = 1
    nat_cps_exceed          = 1
    l7_cps_exceed           = 1
    ssl_cps_exceed          = 1
    ssl_tpt_exceed          = 1
    concurrent_conn_exceed  = 1
    svr_syn_handshake_fail  = 1
    synattack               = 1
  }
}