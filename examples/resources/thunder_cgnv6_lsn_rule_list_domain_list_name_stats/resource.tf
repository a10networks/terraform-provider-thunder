provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_rule_list_domain_list_name_stats" "thunder_cgnv6_lsn_rule_list_domain_list_name_stats" {

  name             = "test"
  name_domain_list = "testing"

}
output "get_cgnv6_lsn_rule_list_domain_list_name_stats" {
  value = ["${data.thunder_cgnv6_lsn_rule_list_domain_list_name_stats.thunder_cgnv6_lsn_rule_list_domain_list_name_stats}"]
}