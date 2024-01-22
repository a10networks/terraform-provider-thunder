provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_alg_sip" "thunder_cgnv6_stateful_firewall_alg_sip" {
  sampling_enable {
    counters1 = "all"
  }
  sip_value = "disable"
}