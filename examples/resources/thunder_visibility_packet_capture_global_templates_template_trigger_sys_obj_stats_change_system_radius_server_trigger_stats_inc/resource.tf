provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_radius_server_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_radius_server_trigger_stats_inc" {

  name                          = "test"
  ha_standby_dropped            = 1
  invalid_key                   = 1
  ipv6_prefix_length_mismatch   = 1
  radius_request_dropped        = 1
  radius_table_full             = 1
  request_bad_secret_dropped    = 1
  request_malformed_dropped     = 1
  request_no_key_vap_dropped    = 1
  secret_not_configured_dropped = 1
}