provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_pbslb" "thunder_system_pbslb" {
  sampling_enable {
    counters1 = "all"
  }
}