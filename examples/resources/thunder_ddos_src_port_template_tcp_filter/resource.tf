provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_tcp_filter" "thunder_ddos_src_port_template_tcp_filter" {
  tcp_filter_action    = "blacklist-src"
  tcp_filter_regex     = "48"
  tcp_filter_seq       = 4
  tcp_filter_unmatched = 0
  user_tag             = "28"
  name                 = "testing"
}