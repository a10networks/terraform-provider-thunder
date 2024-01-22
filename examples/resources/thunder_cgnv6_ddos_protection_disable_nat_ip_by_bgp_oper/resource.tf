provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper" "thunder_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper" {

}
output "get_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper" {
  value = ["${data.thunder_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper.thunder_cgnv6_ddos_protection_disable_nat_ip_by_bgp_oper}"]
}