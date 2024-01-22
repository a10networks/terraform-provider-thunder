provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_rule_list_domain_list_name" "thunder_cgnv6_lsn_rule_list_domain_list_name" {

  name             = "test"
  name_domain_list = "testing"
  sampling_enable {
    counters1 = "all"
  }
}