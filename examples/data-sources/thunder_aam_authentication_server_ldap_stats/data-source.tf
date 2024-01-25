provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_server_ldap_stats" "thunder_aam_authentication_server_ldap_stats" {
}
output "get_aam_authentication_server_ldap_stats" {
  value = ["${data.thunder_aam_authentication_server_ldap_stats.thunder_aam_authentication_server_ldap_stats}"]
}