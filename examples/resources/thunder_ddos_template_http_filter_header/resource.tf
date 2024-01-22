provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_http_filter_header" "thunder_ddos_template_http_filter_header" {
  http_tmpl_name               = "test"
  http_filter_header_seq       = 2
  http_filter_header_regex     = "test"
  http_filter_header_unmatched = 1
  http_filter_header_blacklist = 1
  user_tag                     = "test"
}