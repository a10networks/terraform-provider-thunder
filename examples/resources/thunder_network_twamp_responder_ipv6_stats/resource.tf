provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_twamp_responder_ipv6_stats" "thunder_network_twamp_responder_ipv6_stats" {
}
output "get_network_twamp_responder_ipv6_stats" {
  value = ["${data.thunder_network_twamp_responder_ipv6_stats.thunder_network_twamp_responder_ipv6_stats}"]
}