provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ip_l4_type" "thunder_ddos_dst_interface_ip_l4_type" {
  addr          = "10.10.10.10"
  deny          = 1
  drop_frag_pkt = 1
  protocol      = "tcp"
  user_tag      = "21"

}