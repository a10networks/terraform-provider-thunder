provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_vrid_stats" "thunder_vrrp_a_vrid_stats" {

  vrid_val = 3
}
output "get_vrrp_a_vrid_stats" {
  value = ["${data.thunder_vrrp_a_vrid_stats.thunder_vrrp_a_vrid_stats}"]
}