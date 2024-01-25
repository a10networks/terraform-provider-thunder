provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_so_counters_stats" "thunder_so_counters_stats" {
}
output "get_so_counters_stats" {
  value = ["${data.thunder_so_counters_stats.thunder_so_counters_stats}"]
}