provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_traffic_control_rule_set" "thunder_traffic_control_rule_set" {
  name     = "testing_rule"
  user_tag = "test"
  sampling_enable {
    counters1 = "all"
  }
  rule_list {
    name     = "test_rule"
    user_tag = "test"
    sampling_enable {
      counters1 = "all"
    }
  }
}