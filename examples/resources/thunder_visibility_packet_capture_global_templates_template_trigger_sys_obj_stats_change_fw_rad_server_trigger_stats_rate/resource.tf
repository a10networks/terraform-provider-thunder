provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_rate" {

  name                        = "test"
  duration                    = 60
  ha_standby_dropped          = 1
  invalid_key                 = 1
  ipv6_prefix_length_mismatch = 1
  radius_request_dropped      = 1
  radius_table_full           = 1
  request_bad_secret_dropped  = 1
  request_ignored             = 1
  request_malformed_dropped   = 1
  request_no_key_vap_dropped  = 1
  threshold_exceeded_by       = 5
}