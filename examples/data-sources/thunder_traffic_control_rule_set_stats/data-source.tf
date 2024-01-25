provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_traffic_control_rule_set_stats" "thunder_traffic_control_rule_set_stats" {

  name = "test"

}
output "get_traffic_control_rule_set_stats" {
  value = ["${data.thunder_traffic_control_rule_set_stats.thunder_traffic_control_rule_set_stats}"]
}