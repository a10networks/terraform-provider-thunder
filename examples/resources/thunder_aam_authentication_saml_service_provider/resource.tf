provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_saml_service_provider" "thunder_aam_authentication_saml_service_provider" {
  name = "test"
  require_assertion_signed {
    require_assertion_signed_enable = 1
  }
  saml_request_signed {
    saml_request_signed_disable = 1
  }
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "test"
}