provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_icmp_rate" "thunder_system_icmp_rate" {
  sampling_enable {
    counters1 = "all"
  }
}