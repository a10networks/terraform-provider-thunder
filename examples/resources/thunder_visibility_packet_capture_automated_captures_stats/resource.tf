provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_packet_capture_automated_captures_stats" "thunder_visibility_packet_capture_automated_captures_stats" {
}
output "get_visibility_packet_capture_automated_captures_stats" {
  value = ["${data.thunder_visibility_packet_capture_automated_captures_stats.thunder_visibility_packet_capture_automated_captures_stats}"]
}