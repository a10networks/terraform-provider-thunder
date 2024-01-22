provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ipsec_sa_clients_oper" "thunder_vpn_ipsec_sa_clients_oper" {

}
output "get_vpn_ipsec_sa_clients_oper" {
  value = ["${data.thunder_vpn_ipsec_sa_clients_oper.thunder_vpn_ipsec_sa_clients_oper}"]
}