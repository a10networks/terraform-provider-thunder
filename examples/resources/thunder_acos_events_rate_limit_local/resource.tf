provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_rate_limit_local" "thunder_acos_events_rate_limit_local" {
  limit = 32
}