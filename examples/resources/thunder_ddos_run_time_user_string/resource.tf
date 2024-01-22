provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_run_time_user_string" "thunder_ddos_run_time_user_string" {
  value = "testing"
}