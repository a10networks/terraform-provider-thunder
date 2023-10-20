provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_gtp" "test_thunder_fw_gtp" {
  sampling_enable {
    counters1 = "all"
  }
  gtp_value = "enable"
}