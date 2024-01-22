provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_traffic_control_active_rule_set" "thunder_traffic_control_active_rule_set" {

  name = "test"
}