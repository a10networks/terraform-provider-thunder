provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ipv6_l4_type" "thunder_ddos_dst_interface_ipv6_l4_type" {
  addr          = "2001:db8:3333:4444:5555:6666:7777:8888"
  deny          = 1
  drop_frag_pkt = 1
  protocol      = "tcp"
  user_tag      = "63"

}