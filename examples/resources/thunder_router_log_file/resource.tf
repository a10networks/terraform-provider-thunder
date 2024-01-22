provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_log_file" "thunder_router_log_file" {
  per_protocol = 0
  rotate       = 52
  size         = 647800
}