provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nptv6_domain" "thunder_cgnv6_nptv6_domain" {
  name           = "test"
  inside_prefix  = "2001:db8:3333:4444:5555:6666:7777:8888/64"
  outside_prefix = "2001:db8:3333:4444:5555:6666:7777:8888/64"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "57"
}