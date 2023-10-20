provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_app" "test_thunder_fw_app" {
  sampling_enable {
    counters1 = "all"
  }
}