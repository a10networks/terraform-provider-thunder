provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_bandwidth" "thunder_system_bandwidth" {
  sampling_enable {
    counters1 = "all"
  }
}