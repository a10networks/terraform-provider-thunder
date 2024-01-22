provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  es_resp_400           = 1
  es_resp_500           = 1
  es_resp_invalid_http  = 1
  threshold_exceeded_by = 5
}