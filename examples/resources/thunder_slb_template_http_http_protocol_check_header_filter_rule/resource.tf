provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_http_http_protocol_check_header_filter_rule" "thunder_slb_template_http_http_protocol_check_header_filter_rule" {

  name               = "slb-http"
  action_value       = "drop"
  header_name_value  = "4"
  header_value_value = "12"
  seq_num            = 2
  user_tag           = "39"
}