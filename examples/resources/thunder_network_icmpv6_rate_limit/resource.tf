provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_icmpv6_rate_limit" "thunder_network_icmpv6_rate_limit" {
  icmpv6_lockup            = 65510
  icmpv6_lockup_period     = 4036
  icmpv6_normal_rate_limit = 3000
}