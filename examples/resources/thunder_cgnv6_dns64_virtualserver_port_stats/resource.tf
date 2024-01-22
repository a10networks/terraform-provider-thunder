provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_dns64_virtualserver_port_stats" "thunder_cgnv6_dns64_virtualserver_port_stats" {

  name        = "test"
  port_number = 22866
  protocol    = "dns-udp"

}
output "get_cgnv6_dns64_virtualserver_port_stats" {
  value = ["${data.thunder_cgnv6_dns64_virtualserver_port_stats.thunder_cgnv6_dns64_virtualserver_port_stats}"]
}