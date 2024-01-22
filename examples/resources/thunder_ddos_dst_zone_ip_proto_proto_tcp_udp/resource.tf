provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_tcp_udp" "thunder_ddos_dst_zone_ip_proto_proto_tcp_udp" {
  zone_name     = "testZone"
  deny          = 1
  drop_frag_pkt = 1
  protocol      = "tcp"

}