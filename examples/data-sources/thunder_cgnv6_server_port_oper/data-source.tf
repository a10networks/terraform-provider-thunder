provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_server_port_oper" "thunder_cgnv6_server_port_oper" {

  name        = "test12345"
  protocol    = "tcp"
  port_number = 54519
}
output "get_cgnv6_server_port_oper" {
  value = ["${data.thunder_cgnv6_server_port_oper.thunder_cgnv6_server_port_oper}"]
}