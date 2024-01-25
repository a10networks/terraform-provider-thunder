provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_packet_capture_oper" "thunder_visibility_packet_capture_oper" {
}
output "get_visibility_packet_capture_oper" {
  value = ["${data.thunder_visibility_packet_capture_oper.thunder_visibility_packet_capture_oper}"]
}