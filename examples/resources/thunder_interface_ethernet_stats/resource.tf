provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_ethernet_stats" "thunder_interface_ethernet_stats" {
  ifnum = 2
}
output "get_interface_ethernet_stats" {
  value = ["${data.thunder_interface_ethernet_stats.thunder_interface_ethernet_stats}"]
}