provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_tcp_rst_close_immediate" "test_thunder_fw_tcp_rst_close_immediate" {
  status = "enable"
}