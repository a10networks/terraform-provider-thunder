provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ip" "thunder_ddos_dst_interface_ip" {
  addr = "10.10.10.11"
  ip_proto_list {
    port_num = 114
    deny     = 1
    user_tag = "82"
  }
  l4_type_list {
    protocol      = "tcp"
    deny          = 1
    drop_frag_pkt = 1
    user_tag      = "15"
  }
  log_enable   = 1
  log_periodic = 1
  port_list {
    port_num = 15
    protocol = "tcp"
    deny     = 1
    user_tag = "73"
  }
  user_tag = "104"
}