provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection_ip_proto_proto_name" "thunder_ddos_dst_zone_web_gui_protection_ip_proto_proto_name" {
  zone_name = "testZone"
  pbe       = "103"
  protocol  = "icmp-v4"
  user_tag  = "126"

}