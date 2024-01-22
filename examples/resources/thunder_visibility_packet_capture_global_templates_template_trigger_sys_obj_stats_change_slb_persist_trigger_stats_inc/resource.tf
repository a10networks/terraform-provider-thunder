provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist_trigger_stats_inc" {

  name                       = "test"
  cookie_invalid             = 1
  cookie_not_found           = 1
  cookie_persist_fail        = 1
  cssl_sid_not_found         = 1
  cssl_sid_not_match         = 1
  dst_ip_fail                = 1
  dst_ip_hash_fail           = 1
  dst_ip_new_sess_cache_fail = 1
  dst_ip_new_sess_sel_fail   = 1
  hash_tbl_create_fail       = 1
  hash_tbl_rst_adddel        = 1
  hash_tbl_rst_updown        = 1
  hash_tbl_trylock_fail      = 1
  header_hash_fail           = 1
  src_ip_fail                = 1
  src_ip_hash_fail           = 1
  src_ip_new_sess_cache_fail = 1
  src_ip_new_sess_sel_fail   = 1
  ssl_sid_persist_fail       = 1
  ssl_sid_session_fail       = 1
  sssl_sid_not_found         = 1
  sssl_sid_not_match         = 1
  url_hash_fail              = 1
}