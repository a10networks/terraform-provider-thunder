provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_sip_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_sip_trigger_stats_rate" {

  name                             = "test"
  duration                         = 60
  msg_proxy_client_fail            = 1
  msg_proxy_fail_start_server_conn = 1
  msg_proxy_server_fail            = 1
  threshold_exceeded_by            = 5
}