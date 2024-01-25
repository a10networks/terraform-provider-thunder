provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_error_stats" "thunder_vpn_error_stats" {

}
output "get_vpn_error_stats" {
  value = ["${data.thunder_vpn_error_stats.thunder_vpn_error_stats}"]
}