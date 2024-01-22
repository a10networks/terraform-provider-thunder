provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_in_ipv4_frag" "thunder_ipv6_in_ipv4_frag" {
  sampling_enable {
    counters1 = "all"
  }
}