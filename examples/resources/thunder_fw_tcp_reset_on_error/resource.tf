provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_tcp_reset_on_error" "test_thunder_fw_tcp_reset_on_error" {
  enable = 1
}