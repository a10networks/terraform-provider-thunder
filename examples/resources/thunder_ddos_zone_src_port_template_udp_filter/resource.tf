provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_udp_filter" "thunder_ddos_zone_src_port_template_udp_filter" {
  name                     = "test"
  udp_filter_action        = "drop"
  udp_filter_inverse_match = 1
  udp_filter_name          = "46"
  udp_filter_regex         = "513"
  udp_filter_seq           = 68
  user_tag                 = "48"
}