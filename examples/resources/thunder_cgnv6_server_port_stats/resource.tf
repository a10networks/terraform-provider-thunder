provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_server_port_stats" "thunder_cgnv6_server_port_stats" {

  name        = "test12345"
  protocol    = "tcp"
  port_number = 54519

}
output "get_cgnv6_server_port_stats" {
  value = ["${data.thunder_cgnv6_server_port_stats.thunder_cgnv6_server_port_stats}"]
}