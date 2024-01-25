provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_traffic_control_rule_set_oper" "thunder_traffic_control_rule_set_oper" {

  name = "testing_rule"
}
output "get_traffic_control_rule_set_oper" {
  value = ["${data.thunder_traffic_control_rule_set_oper.thunder_traffic_control_rule_set_oper}"]
}