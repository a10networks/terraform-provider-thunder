provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_tcp" "thunder_ddos_src_port_template_tcp" {

  filter_list {
    tcp_filter_seq       = 1
    tcp_filter_regex     = "1049"
    tcp_filter_unmatched = 1
    tcp_filter_action    = "blacklist-src"
    user_tag             = "89"
  }
  user_tag = "24"
  name     = "testing"
}