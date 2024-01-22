provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_traffic_control_rule_set_rule_oper" "thunder_traffic_control_rule_set_rule_oper" {

  name = "test"
}
output "get_traffic_control_rule_set_rule_oper" {
  value = ["${data.thunder_traffic_control_rule_set_rule_oper.thunder_traffic_control_rule_set_rule_oper}"]
}