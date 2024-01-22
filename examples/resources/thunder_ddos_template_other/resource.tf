provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_other" "thunder_ddos_template_other" {
  name = "test"
  filter_list {
    other_filter_seq       = 1
    other_filter_regex     = "1266"
    other_filter_unmatched = 1
    other_filter_action    = "blacklist-src"
    user_tag               = "29"
  }
  user_tag = "7"

}