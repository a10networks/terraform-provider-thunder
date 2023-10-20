provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_gtp_in_gtp_filtering" "test_thunder_fw_gtp_in_gtp_filtering" {
  gtp_in_gtp_value = "string_val"
}