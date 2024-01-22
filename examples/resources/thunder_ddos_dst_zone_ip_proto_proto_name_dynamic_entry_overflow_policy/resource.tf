provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_name_dynamic_entry_overflow_policy" "thunder_ddos_dst_zone_ip_proto_proto_name_dynamic_entry_overflow_policy" {
  zone_name  = "testZone"
  action     = "bypass"
  dummy_name = "configuration"
  protocol   = "icmp-v4"
  user_tag   = "34"
}