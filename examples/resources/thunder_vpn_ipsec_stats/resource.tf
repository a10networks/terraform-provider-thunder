provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ipsec_stats" "thunder_vpn_ipsec_stats" {
  name = "test"

}
output "get_vpn_ipsec_stats" {
  value = ["${data.thunder_vpn_ipsec_stats.thunder_vpn_ipsec_stats}"]
}