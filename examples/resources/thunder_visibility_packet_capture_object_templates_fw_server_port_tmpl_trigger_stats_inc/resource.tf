provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc" {

  name                 = "test"
  es_resp_400          = 1
  es_resp_500          = 1
  es_resp_invalid_http = 1
}