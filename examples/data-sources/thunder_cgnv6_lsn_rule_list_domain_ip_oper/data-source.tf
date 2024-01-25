provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_rule_list_domain_ip_oper" "thunder_cgnv6_lsn_rule_list_domain_ip_oper" {

  name = "test"
}
output "get_cgnv6_lsn_rule_list_domain_ip_oper" {
  value = ["${data.thunder_cgnv6_lsn_rule_list_domain_ip_oper.thunder_cgnv6_lsn_rule_list_domain_ip_oper}"]
}