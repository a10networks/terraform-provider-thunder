provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_dns64_virtualserver" "thunder_cgnv6_dns64_virtualserver" {
  enable_disable_action = "enable"
  ip_address            = "10.10.10.10"
  name                  = "test"
  policy                = 0
  port_list {
    port_number = 15592
    protocol    = "dns-udp"
    action      = "enable"
    auto        = 0
    precedence  = 0
    user_tag    = "114"
    sampling_enable {
      counters1 = "all"
    }

  }
}