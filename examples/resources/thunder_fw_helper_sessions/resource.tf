provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_helper_sessions" "test_thunder_fw_helper_sessions" {
  mode         = "disable"
  idle_timeout = 30
  limit        = 33
}