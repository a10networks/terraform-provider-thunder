provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_filter" "thunder_ddos_template_tcp_filter" {
  name                 = "test"
  tcp_filter_action    = "blacklist-src"
  tcp_filter_regex     = "278"
  tcp_filter_seq       = 2
  tcp_filter_unmatched = 1
  user_tag             = "49"

}