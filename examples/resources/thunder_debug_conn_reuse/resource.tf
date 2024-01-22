provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_conn_reuse" "thunder_debug_conn_reuse" {
  level = 2
}