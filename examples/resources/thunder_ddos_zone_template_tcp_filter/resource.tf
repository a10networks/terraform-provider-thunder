provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp_filter" "thunder_ddos_zone_template_tcp_filter" {
  name                     = "test"
  tcp_filter_action        = "drop"
  tcp_filter_inverse_match = 1
  tcp_filter_name          = "5"
  tcp_filter_regex         = "592"
  tcp_filter_seq           = 163
  user_tag                 = "45"
}