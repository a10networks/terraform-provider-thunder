provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_system_status" "test_thunder_fw_system_status" {
  uuid = "string"
}