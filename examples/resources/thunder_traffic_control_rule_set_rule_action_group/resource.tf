provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_traffic_control_rule_set_rule_action_group" "thunder_traffic_control_rule_set_rule_action_group" {

  name         = "test"
  limit_policy = 3
}