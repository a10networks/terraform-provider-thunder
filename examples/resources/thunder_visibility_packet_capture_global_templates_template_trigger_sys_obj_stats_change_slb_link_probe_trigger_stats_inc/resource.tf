provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_link_probe_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_link_probe_trigger_stats_inc" {

  name                         = "test"
  err_entry_create_failed      = 1
  err_entry_create_oom         = 1
  err_entry_insert_failed      = 1
  err_l4_sess_alloc            = 1
  err_probe_tcp_conn_send      = 1
  err_smart_nat_alloc          = 1
  err_smart_nat_port_alloc     = 1
  err_tmpl_probe_create_failed = 1
  err_tmpl_probe_create_oom    = 1
  total_http_response_bad      = 1
  total_tcp_err                = 1
}