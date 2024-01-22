provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_rule_list_domain_ip_stats" "thunder_cgnv6_lsn_rule_list_domain_ip_stats" {

  name = "test"
}
output "get_cgnv6_lsn_rule_list_domain_ip_stats" {
  value = ["${data.thunder_cgnv6_lsn_rule_list_domain_ip_stats.thunder_cgnv6_lsn_rule_list_domain_ip_stats}"]
}