provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_automated_captures" "thunder_visibility_packet_capture_automated_captures" {
  slb_port_tmpl_error_code_return_inc  = 0
  slb_port_tmpl_high_error_code_return = 0
}