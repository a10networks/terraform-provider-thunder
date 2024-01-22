provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_form_based_instance_request_uri" "thunder_aam_authentication_relay_form_based_instance_request_uri" {

  action_uri              = "www.test.com"
  name                    = "test"
  domain_variable         = "12"
  match_type              = "equals"
  max_packet_collect_size = 1048576
  other_variables         = "122"
  password_variable       = "48"
  uri                     = "www.a10.com"
  user_tag                = "13"
  user_variable           = "63"
}