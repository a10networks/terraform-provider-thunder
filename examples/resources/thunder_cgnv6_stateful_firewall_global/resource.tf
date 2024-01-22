provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_global" "thunder_cgnv6_stateful_firewall_global" {
  respond_to_user_mac = 0
  sampling_enable {
    counters1 = "all"
  }
  stateful_firewall_value = "enable"
}