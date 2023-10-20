provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_vrid" "test_thunder_fw_vrid" {
  vrid = 3
}