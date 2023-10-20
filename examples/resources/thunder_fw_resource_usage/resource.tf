provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_resource_usage" "test_thunder_fw_resource_usage" {
  uuid = "string"
}