provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_jwt" "thunder_aam_authentication_jwt" {
  action           = "redirect"
  issuer           = "101"
  jwt_relay_uri    = "79"
  name             = "27"
  secret_string    = "84"
  signature_secret = 1
  token_lifetime   = 300
  user_tag         = "75"

}