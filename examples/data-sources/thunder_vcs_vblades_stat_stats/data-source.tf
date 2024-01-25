provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vcs_vblades_stat_stats" "thunder_vcs_vblades_stat_stats" {
  vblade_id = 1
}
output "get_vcs_vblades_stat_stats" {
  value = ["${data.thunder_vcs_vblades_stat_stats.thunder_vcs_vblades_stat_stats}"]
}