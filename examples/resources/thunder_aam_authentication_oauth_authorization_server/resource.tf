provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_oauth_authorization_server" "thunder_aam_authentication_oauth_authorization_server" {
  authorization_endpoint = "90"
  client_method          = "ignored"
  issuer                 = "8"
  name                   = "30"
  sampling_enable {
    counters1 = "all"
  }
  server_method  = "post"
  token_endpoint = "73"
  user_tag       = "33"
}