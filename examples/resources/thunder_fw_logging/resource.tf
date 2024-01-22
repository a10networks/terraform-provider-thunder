provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_logging" "thunder_fw_logging" {
  gtp {
    sampling_enable {
      counters1 = "all"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}