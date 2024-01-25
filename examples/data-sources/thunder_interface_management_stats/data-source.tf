provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_management_stats" "thunder_interface_management_stats" {
}
output "get_interface_management_stats" {
  value = ["${data.thunder_interface_management_stats.thunder_interface_management_stats}"]
}