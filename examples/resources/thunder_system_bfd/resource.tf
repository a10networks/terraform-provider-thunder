provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_bfd" "thunder_system_bfd" {
  sampling_enable {
    counters1 = "all"
  }
}