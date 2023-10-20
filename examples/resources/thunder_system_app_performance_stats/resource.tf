provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_system_app_performance_stats" "app_performance" {}

output "get_app_performance" {
  value = ["${data.thunder_system_app_performance_stats.app_performance}"]
}