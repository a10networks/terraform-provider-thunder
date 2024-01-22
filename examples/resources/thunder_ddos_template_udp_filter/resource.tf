provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_udp_filter" "thunder_ddos_template_udp_filter" {
  name                 = "test"
  udp_filter_action    = "blacklist-src"
  udp_filter_regex     = "1252"
  udp_filter_seq       = 3
  udp_filter_unmatched = 1
  user_tag             = "94"

}