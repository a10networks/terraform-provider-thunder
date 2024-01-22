provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_dns_cache" "thunder_system_dns_cache" {
  sampling_enable {
    counters1 = "all"
  }
}