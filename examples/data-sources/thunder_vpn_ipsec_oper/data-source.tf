provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ipsec_oper" "thunder_vpn_ipsec_oper" {
  name = "test"

}
output "get_vpn_ipsec_oper" {
  value = ["${data.thunder_vpn_ipsec_oper.thunder_vpn_ipsec_oper}"]
}