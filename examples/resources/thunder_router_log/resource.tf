provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_log" "thunder_router_log" {

  file {
    per_protocol = 0
    rotate       = 95
    size         = 298733
  }
  log_buffer = 1
}