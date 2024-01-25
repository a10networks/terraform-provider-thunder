provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vcs_stat_stats" "thunder_vcs_stat_stats" {
}
output "get_vcs_stat_stats" {
  value = ["${data.thunder_vcs_stat_stats.thunder_vcs_stat_stats}"]
}