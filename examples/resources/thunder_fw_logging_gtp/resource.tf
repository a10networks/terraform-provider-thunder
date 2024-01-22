provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_logging_gtp" "thunder_fw_logging_gtp" {
  sampling_enable {
    counters1 = "all"
  }
}