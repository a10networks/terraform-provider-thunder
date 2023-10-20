provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_top_k_rules" "test_thunder_fw_top_k_rules" {
  uuid = "string"
}