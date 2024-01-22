provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_saml_service_provider_stats" "thunder_aam_authentication_saml_service_provider_stats" {

  name = "test"
}
output "get_aam_authentication_saml_service_provider_stats" {
  value = ["${data.thunder_aam_authentication_saml_service_provider_stats.thunder_aam_authentication_saml_service_provider_stats}"]
}