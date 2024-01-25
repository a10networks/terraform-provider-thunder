provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dns_tcp_zone_port_stats" "thunder_ddos_dns_tcp_zone_port_stats" {

}
output "get_ddos_dns_tcp_zone_port_stats" {
  value = ["${data.thunder_ddos_dns_tcp_zone_port_stats.thunder_ddos_dns_tcp_zone_port_stats}"]
}