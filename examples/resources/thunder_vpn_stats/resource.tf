provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_stats" "thunder_vpn_stats" {

}
output "get_vpn_stats" {
  value = ["${data.thunder_vpn_stats.thunder_vpn_stats}"]
}