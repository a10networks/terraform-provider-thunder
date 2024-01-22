provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_ve_stats" "thunder_network_ve_stats" {
}
output "get_network_ve_stats" {
  value = ["${data.thunder_network_ve_stats.thunder_network_ve_stats}"]
}