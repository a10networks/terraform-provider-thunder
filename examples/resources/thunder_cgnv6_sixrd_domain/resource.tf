provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sixrd_domain" "thunder_cgnv6_sixrd_domain" {
  name = "test"
  mtu  = 1429
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "120"
}