provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_reporting_stats" "thunder_visibility_reporting_stats" {
}
output "get_visibility_reporting_stats" {
  value = ["${data.thunder_visibility_reporting_stats.thunder_visibility_reporting_stats}"]
}