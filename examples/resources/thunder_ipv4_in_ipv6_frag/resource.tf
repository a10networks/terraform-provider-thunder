provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv4_in_ipv6_frag" "thunder_ipv4_in_ipv6_frag" {
  sampling_enable {
    counters1 = "all"
  }
}