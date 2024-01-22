provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_sip_filter_header" "thunder_ddos_template_sip_filter_header" {
  sip_tmpl_name               = "test"
  sip_filter_header_seq       = 2
  sip_filter_header_regex     = "test"
  sip_filter_header_unmatched = 1
  sip_filter_header_blacklist = 1
  user_tag                    = "est"

}