provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_traffic_capture" "thunder_debug_traffic_capture" {
  enable                     = 1
  maximum_file_size          = 52
  maximum_history_recordings = 5
  record_direction_type      = "ingress-only"
}