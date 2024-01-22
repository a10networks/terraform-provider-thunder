provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ipv6" "thunder_ddos_dst_interface_ipv6" {

  addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  ip_proto_list {
    port_num = 16
    deny     = 1
    user_tag = "116"
  }
  l4_type_list {
    protocol      = "tcp"
    deny          = 1
    drop_frag_pkt = 1
    user_tag      = "7"
  }
  log_enable = 1
  port_list {
    port_num = 56
    protocol = "tcp"
    deny     = 1
    user_tag = "32"
  }
  user_tag = "126"

}