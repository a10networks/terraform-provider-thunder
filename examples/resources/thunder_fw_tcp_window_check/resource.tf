provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_tcp_window_check" "test_thunder_fw_tcp_window_check" {
  status = "enable"
  sampling_enable {
    counters1 = "all"
  }
}