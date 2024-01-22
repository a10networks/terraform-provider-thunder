provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ike_gateway_stats" "thunder_vpn_ike_gateway_stats" {

  name = "test"

}
output "get_vpn_ike_gateway_stats" {
  value = ["${data.thunder_vpn_ike_gateway_stats.thunder_vpn_ike_gateway_stats}"]
}