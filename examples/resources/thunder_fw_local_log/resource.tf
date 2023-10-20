provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_local_log" "test_thunder_fw_local_log" {
  local_logging = 1
}