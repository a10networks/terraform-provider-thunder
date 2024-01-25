provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_view_memory_view_stats" "thunder_system_view_memory_view_stats" {
}
output "get_system_view_memory_view_stats" {
  value = ["${data.thunder_system_view_memory_view_stats.thunder_system_view_memory_view_stats}"]
}