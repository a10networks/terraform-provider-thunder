provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_other_filter" "thunder_ddos_template_other_filter" {
  name                   = "test"
  other_filter_action    = "blacklist-src"
  other_filter_regex     = "704"
  other_filter_seq       = 2
  other_filter_unmatched = 0
  user_tag               = "11"

}