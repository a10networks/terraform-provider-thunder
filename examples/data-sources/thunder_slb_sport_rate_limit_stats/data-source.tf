provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_sport_rate_limit_stats" "thunder_slb_sport_rate_limit_stats" {
}
output "get_slb_sport_rate_limit_stats" {
  value = ["${data.thunder_slb_sport_rate_limit_stats.thunder_slb_sport_rate_limit_stats}"]
}