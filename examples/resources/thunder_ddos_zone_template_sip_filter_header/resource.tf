provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_sip_filter_header" "thunder_ddos_zone_template_sip_filter_header" {
  sip_tmpl_name         = "test"
  sip_filter_action     = "drop"
  sip_filter_header_seq = 33
  sip_filter_name       = "test"
  sip_header_cfg {
    sip_filter_header_regex         = "1011"
    sip_filter_header_inverse_match = 0
  }
  user_tag = "58"
}