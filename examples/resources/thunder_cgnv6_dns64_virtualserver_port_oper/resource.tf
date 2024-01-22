provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_dns64_virtualserver_port_oper" "thunder_cgnv6_dns64_virtualserver_port_oper" {

  name        = "test"
  port_number = 22866
  protocol    = "dns-udp"

}
output "get_cgnv6_dns64_virtualserver_port_oper" {
  value = ["${data.thunder_cgnv6_dns64_virtualserver_port_oper.thunder_cgnv6_dns64_virtualserver_port_oper}"]
}