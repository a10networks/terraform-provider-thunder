provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_netflow_monitor_stats" "thunder_netflow_monitor_stats" {

  name = "a11"
}

output "get_netflow_monitor_stats" {
  value = ["${data.thunder_netflow_monitor_stats.thunder_netflow_monitor_stats}"]
}