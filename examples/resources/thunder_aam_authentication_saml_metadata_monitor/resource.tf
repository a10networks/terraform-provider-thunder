provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_saml_metadata_monitor" "thunder_aam_authentication_saml_metadata_monitor" {
  acs_continuous_fail_threshold = 32
  acs_missing_period            = 144
  acs_missing_threshold         = 216
  status                        = "enable"
}