provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_app_performance" "thunder_system_app_performance" {
  sampling_enable {
    counters1 = "all"
  }
}