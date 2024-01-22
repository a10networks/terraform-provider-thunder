provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_interface_tunnel_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_interface_tunnel_tmpl_trigger_stats_inc" {

  name            = "test"
  num_rx_err_pkts = 1
  num_tx_err_pkts = 1
}