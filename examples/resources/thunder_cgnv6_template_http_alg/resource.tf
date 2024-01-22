provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_http_alg" "thunder_cgnv6_template_http_alg" {
  name                     = "test"
  header_name_client_ip    = "X-Forwarded-For"
  header_name_msisdn       = "X-MSISDN"
  include_tunnel_ip        = 1
  method                   = "append"
  request_insert_client_ip = 1
  request_insert_msisdn    = 1
  user_tag                 = "61"
}