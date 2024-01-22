provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_udp" "thunder_ddos_zone_src_port_template_udp" {

  name = "test"
  ntp_monlist_cfg {
    ntp_monlist        = 1
    ntp_monlist_action = "ignore"
  }
  max_payload_size_cfg {
    max_payload_size        = 3
    max_payload_size_action = "drop"
  }
  user_tag = "test"
  filter_list {
    udp_filter_name   = "test"
    udp_filter_seq    = 3
    udp_filter_regex  = "test"
    udp_filter_action = "ignore"
    user_tag          = "test"
  }
}