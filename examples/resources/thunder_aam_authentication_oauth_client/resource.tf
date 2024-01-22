provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_oauth_client" "thunder_aam_authentication_oauth_client" {
  client_id              = "56"
  client_secret          = "79"
  grant_type             = "implicit"
  infinity               = 0
  name                   = "62"
  no_reply               = 0
  parameter_nonce_enable = 0
  redirection_endpoint   = "58"
  scope                  = "86"
  session_init_ttl       = 41
  type                   = "openid-connect"
  user_tag               = "54"
}