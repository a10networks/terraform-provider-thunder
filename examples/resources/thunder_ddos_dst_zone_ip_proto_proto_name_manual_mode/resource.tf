provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_name_manual_mode" "thunder_ddos_dst_zone_ip_proto_proto_name_manual_mode" {
  zone_name = "testZone"
  protocol  = "icmp-v4"
  config    = "configuration"
  user_tag  = "94"
}