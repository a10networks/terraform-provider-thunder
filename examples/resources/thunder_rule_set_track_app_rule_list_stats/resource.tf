provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_track_app_rule_list_stats" "thunder_rule_set_track_app_rule_list_stats" {

  name = "test"
}
output "get_rule_set_track_app_rule_list_stats" {
  value = ["${data.thunder_rule_set_track_app_rule_list_stats.thunder_rule_set_track_app_rule_list_stats}"]
}