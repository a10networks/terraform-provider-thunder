provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_hw_accelerate" "thunder_fw_hw_accelerate" {
  sampling_enable {
    counters1 = "all"
  }
}