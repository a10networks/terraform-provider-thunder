provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_rule_list_domain_ip" "thunder_cgnv6_lsn_rule_list_domain_ip" {

  name = "test"
  sampling_enable {
    counters1 = "all"
  }
}