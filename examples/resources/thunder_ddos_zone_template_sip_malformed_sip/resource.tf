provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_sip_malformed_sip" "thunder_ddos_zone_template_sip_malformed_sip" {
  sip_tmpl_name                         = "test"
  malformed_sip_check                   = "enable-check"
  malformed_sip_max_line_size           = 2
  malformed_sip_max_uri_length          = 3
  malformed_sip_max_header_name_length  = 3
  malformed_sip_max_header_value_length = 4
  malformed_sip_call_id_max_length      = 2
  malformed_sip_sdp_max_length          = 3
  malformed_sip_action                  = "reset"

}