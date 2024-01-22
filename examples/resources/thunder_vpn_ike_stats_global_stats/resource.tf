provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ike_stats_global_stats" "thunder_vpn_ike_stats_global_stats" {

}
output "get_vpn_ike_stats_global_stats" {
  value = ["${data.thunder_vpn_ike_stats_global_stats.thunder_vpn_ike_stats_global_stats}"]
}