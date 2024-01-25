provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_rc_cache_global_stats" "thunder_slb_rc_cache_global_stats" {
}
output "get_slb_rc_cache_global_stats" {
  value = ["${data.thunder_slb_rc_cache_global_stats.thunder_slb_rc_cache_global_stats}"]
}