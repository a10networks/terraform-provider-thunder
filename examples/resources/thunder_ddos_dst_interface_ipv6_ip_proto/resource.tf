provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ipv6_ip_proto" "thunder_ddos_dst_interface_ipv6_ip_proto" {

  deny     = 1
  port_num = 31
  user_tag = "8"
  addr     = "2001:db8:3333:4444:5555:6666:7777:8888"
}