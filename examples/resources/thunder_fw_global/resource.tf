provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_global" "test_thunder_fw_global" {
  sampling_enable {
    counters1 = "all"
  }
}