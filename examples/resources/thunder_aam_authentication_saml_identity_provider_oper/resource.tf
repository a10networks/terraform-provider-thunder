provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_saml_identity_provider_oper" "thunder_aam_authentication_saml_identity_provider_oper" {

  name = "test"

}
output "get_aam_authentication_saml_identity_provider_oper" {
  value = ["${data.thunder_aam_authentication_saml_identity_provider_oper.thunder_aam_authentication_saml_identity_provider_oper}"]
}