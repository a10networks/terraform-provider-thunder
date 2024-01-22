provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_dns_persistent_cache" "thunder_slb_dns_persistent_cache" {
  sampling_enable {
    counters1 = "all"
  }
}