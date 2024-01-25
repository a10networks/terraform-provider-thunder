provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_packet_capture_capture_config_stats" "thunder_visibility_packet_capture_capture_config_stats" {
  name = "test"


}
output "get_visibility_packet_capture_capture_config_stats" {
  value = ["${data.thunder_visibility_packet_capture_capture_config_stats.thunder_visibility_packet_capture_capture_config_stats}"]
}