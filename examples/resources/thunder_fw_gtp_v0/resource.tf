provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_gtp_v0" "test_thunder_fw_gtp_v0" {
  gtpv0_value = "string_val"
}