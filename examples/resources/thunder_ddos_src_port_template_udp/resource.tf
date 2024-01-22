provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_udp" "thunder_ddos_src_port_template_udp" {
  drop_ntp_monlist = 1
  filter_list {
    udp_filter_seq       = 4
    udp_filter_regex     = "609"
    udp_filter_unmatched = 1
    udp_filter_action    = "blacklist-src"
    user_tag             = "64"
  }
  max_payload_size = 620
  min_payload_size = 82
  user_tag         = "114"
  name             = "testing"
}