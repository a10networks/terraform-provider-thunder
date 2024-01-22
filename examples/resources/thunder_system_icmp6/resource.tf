provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_icmp6" "thunder_system_icmp6" {
  sampling_enable {
    counters1 = "all"
  }
}