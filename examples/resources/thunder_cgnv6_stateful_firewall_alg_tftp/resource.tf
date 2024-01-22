provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_alg_tftp" "thunder_cgnv6_stateful_firewall_alg_tftp" {
  sampling_enable {
    counters1 = "all"
  }
  tftp_value = "disable"
}