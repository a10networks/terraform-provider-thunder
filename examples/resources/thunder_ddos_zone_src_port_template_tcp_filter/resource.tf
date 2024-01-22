provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_tcp_filter" "thunder_ddos_zone_src_port_template_tcp_filter" {

  tcp_filter_action        = "drop"
  tcp_filter_inverse_match = 1
  tcp_filter_name          = "10"
  tcp_filter_regex         = "174"
  tcp_filter_seq           = 116
  user_tag                 = "4"
  name                     = "test"
}