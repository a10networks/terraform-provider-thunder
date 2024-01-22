provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate" {

  name                  = "test"
  collisions            = 1
  crc                   = 1
  duration              = 60
  giants                = 1
  giants_output         = 1
  input_errors          = 1
  output_errors         = 1
  runts                 = 1
  threshold_exceeded_by = 5
}