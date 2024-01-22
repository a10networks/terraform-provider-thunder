provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_udp_filter" "thunder_ddos_src_port_template_udp_filter" {
  name                 = "testing"
  udp_filter_action    = "blacklist-src"
  udp_filter_regex     = "919"
  udp_filter_seq       = 5
  udp_filter_unmatched = 0
  user_tag             = "41"
}