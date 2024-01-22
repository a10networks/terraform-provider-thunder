provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_http_http_protocol_check" "thunder_slb_template_http_http_protocol_check" {

  name                        = "slb-http"
  get_and_payload             = "drop"
  h2up_content_length_alias   = "drop"
  h2up_with_host_and_auth     = "drop"
  h2up_with_transfer_encoding = "drop"
  header_filter_rule_list {
    seq_num            = 1
    header_name_value  = "19"
    header_value_value = "57"
    action_value       = "drop"
    user_tag           = "47"
  }
  malformed_h2up_header_value          = "drop"
  malformed_h2up_scheme_value          = "drop"
  multiple_content_length              = "drop"
  multiple_transfer_encoding           = "drop"
  transfer_encoding_and_content_length = "drop"
}