provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_virtual_wire_global_stats" "thunder_network_virtual_wire_global_stats" {
}
output "get_network_virtual_wire_global_stats" {
  value = ["${data.thunder_network_virtual_wire_global_stats.thunder_network_virtual_wire_global_stats}"]
}