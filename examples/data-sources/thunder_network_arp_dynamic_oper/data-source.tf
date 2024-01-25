provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_arp_dynamic_oper" "thunder_network_arp_dynamic_oper" {
}
output "get_network_arp_dynamic_oper" {
  value = ["${data.thunder_network_arp_dynamic_oper.thunder_network_arp_dynamic_oper}"]
}