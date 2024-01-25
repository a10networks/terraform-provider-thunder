provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_acos_events_statistics_stats" "thunder_acos_events_statistics_stats" {
}
output "get_acos_events_statistics_stats" {
  value = ["${data.thunder_acos_events_statistics_stats.thunder_acos_events_statistics_stats}"]
}