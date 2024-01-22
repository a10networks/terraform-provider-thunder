provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_rules_by_zone_stats" "thunder_rule_set_rules_by_zone_stats" {
  stats {
    dummy = dummy
  }

}
output "get_rule_set_rules_by_zone_stats" {
  value = ["${data.thunder_rule_set_rules_by_zone_stats.thunder_rule_set_rules_by_zone_stats}"]
}