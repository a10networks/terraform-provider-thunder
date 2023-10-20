provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_status" "test_thunder_fw_status" {
  uuid = "string"
}