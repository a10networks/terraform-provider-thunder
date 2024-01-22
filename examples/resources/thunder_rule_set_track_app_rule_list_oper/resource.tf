provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_track_app_rule_list_oper" "thunder_rule_set_track_app_rule_list_oper" {
  oper {
    rule_list {
    }
  }

}
output "get_rule_set_track_app_rule_list_oper" {
  value = ["${data.thunder_rule_set_track_app_rule_list_oper.thunder_rule_set_track_app_rule_list_oper}"]
}