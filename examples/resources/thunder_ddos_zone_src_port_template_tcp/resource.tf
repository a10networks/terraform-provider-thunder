provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_tcp" "thunder_ddos_zone_src_port_template_tcp" {

  filter_list {
    tcp_filter_name          = "36"
    tcp_filter_seq           = 41
    tcp_filter_regex         = "1100"
    tcp_filter_inverse_match = 1
    tcp_filter_action        = "drop"
    user_tag                 = "75"
  }
  user_tag = "84"
  name     = "test"
}