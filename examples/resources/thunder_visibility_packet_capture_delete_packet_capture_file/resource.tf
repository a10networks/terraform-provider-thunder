provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_delete_packet_capture_file" "thunder_visibility_packet_capture_delete_packet_capture_file" {
  all       = 1
  file_name = "test"
}