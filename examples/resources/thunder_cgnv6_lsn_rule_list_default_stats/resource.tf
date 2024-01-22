provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_rule_list_default_stats" "thunder_cgnv6_lsn_rule_list_default_stats" {

  name = "test"
}
output "get_cgnv6_lsn_rule_list_default_stats" {
  value = ["${data.thunder_cgnv6_lsn_rule_list_default_stats.thunder_cgnv6_lsn_rule_list_default_stats}"]
}