provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_number_dynamic_entry_overflow_policy" "thunder_ddos_dst_zone_ip_proto_proto_number_dynamic_entry_overflow_policy" {
  zone_name    = "testZone"
  action       = "bypass"
  dummy_name   = "configuration"
  user_tag     = "116"
  protocol_num = 223
}