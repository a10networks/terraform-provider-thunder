provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_saml_global" "thunder_aam_authentication_saml_global" {
  sampling_enable {
    counters1 = "all"
  }
}