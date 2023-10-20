provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_limit_entry" "test_thunder_fw_limit_entry" {
  uuid = "string"
}