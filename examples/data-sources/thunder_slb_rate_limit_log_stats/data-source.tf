provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_rate_limit_log_stats" "thunder_slb_rate_limit_log_stats" {
}
output "get_slb_rate_limit_log_stats" {
  value = ["${data.thunder_slb_rate_limit_log_stats.thunder_slb_rate_limit_log_stats}"]
}