provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_saml" "thunder_aam_authentication_relay_saml" {
  idp_auth_uri  = "11"
  method        = "get-from-backend"
  name          = "56"
  relay_acs_uri = "11"
  retry_number  = 0
  sampling_enable {
    counters1 = "all"
  }
  server_cookie_name = "79"
  user_tag           = "34"
}