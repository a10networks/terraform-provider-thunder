provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_local_log" "thunder_debug_local_log" {
  group = "query"
  level = 1
}