provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_prefix_list" "thunder_ipv6_prefix_list" {
  name = "a10"
  rules {
    seq    = 257495
    action = "deny"
    ipaddr = "21c9:f6a4:d8ca:fa07:fa29:85a6:1b61:2d1/24"
    ge     = 26
    le     = 77
  }
}