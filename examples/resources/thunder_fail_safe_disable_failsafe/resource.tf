provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fail_safe_disable_failsafe" "thunder_fail_safe_disable_failsafe" {
  action = "session-memory"
}