provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_fpga_drop_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_fpga_drop_trigger_stats_rate" {

  name                  = "test"
  bad_ip_chksum_drop    = 1
  bad_ip_flags_drop     = 1
  bad_ip_frag_offset    = 1
  bad_ip_hdr_len        = 1
  bad_ip_payload_len    = 1
  bad_ip_ttl_drop       = 1
  duration              = 60
  empty_frag_drop       = 1
  fcs_drop              = 1
  hrx_drop              = 1
  icmp_pod_drop         = 1
  ipv4_frag             = 1
  ipv4_opt_drop         = 1
  land_drop             = 1
  mic_frag_drop         = 1
  mrx_drop              = 1
  no_ip_payload_drop    = 1
  oversize_ip_payload   = 1
  runt_ip_hdr           = 1
  runt_tcpudp_hdr       = 1
  siz_drop              = 1
  tcp_bad_chksum        = 1
  tcp_bad_ip_len        = 1
  tcp_bad_urg_offet     = 1
  tcp_fin_sin           = 1
  tcp_frag_hdr          = 1
  tcp_null_flags        = 1
  tcp_null_scan         = 1
  tcp_short_hdr         = 1
  tcp_syn_frag          = 1
  tcp_xmas_flags        = 1
  tcp_xmas_scan         = 1
  threshold_exceeded_by = 5
  tun_mismatch          = 1
  udp_bad_chksum        = 1
  udp_bad_ip_len        = 1
  udp_kb_frags          = 1
  udp_port_lb           = 1
  udp_short_hdr         = 1
}