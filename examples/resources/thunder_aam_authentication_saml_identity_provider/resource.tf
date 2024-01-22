provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_saml_identity_provider" "thunder_aam_authentication_saml_identity_provider" {
  name            = "35"
  reload_interval = 28800
  reload_metadata = 0
  user_tag        = "11"
}