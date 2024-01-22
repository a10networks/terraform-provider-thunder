provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_packet_capture" "thunder_debug_packet_capture" {
  capture_config = "test"
}