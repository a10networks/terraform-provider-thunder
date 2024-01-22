provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_icmp_rate_limit" "thunder_network_icmp_rate_limit" {
  icmp_lockup            = 20383
  icmp_lockup_period     = 8706
  icmp_normal_rate_limit = 3000
}