provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_lif_stats" "thunder_network_lif_stats" {
}
output "get_network_lif_stats" {
  value = ["${data.thunder_network_lif_stats.thunder_network_lif_stats}"]
}