provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_port_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_slb_port_tmpl_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  es_resp_300           = 1
  es_resp_400           = 1
  es_resp_500           = 1
  resp_3xx              = 1
  resp_4xx              = 1
  resp_5xx              = 1
  threshold_exceeded_by = 5
}