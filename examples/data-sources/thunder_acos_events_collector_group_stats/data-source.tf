provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_acos_events_collector_group_stats" "thunder_acos_events_collector_group_stats" {

  name = "test"

}
output "get_acos_events_collector_group_stats" {
  value = ["${data.thunder_acos_events_collector_group_stats.thunder_acos_events_collector_group_stats}"]
}