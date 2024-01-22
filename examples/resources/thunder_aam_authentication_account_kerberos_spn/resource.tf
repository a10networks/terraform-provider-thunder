provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_account_kerberos_spn" "thunder_aam_authentication_account_kerberos_spn" {

  account                = "45"
  name                   = "52"
  password               = 0
  realm                  = "61"
  service_principal_name = "8"
  user_tag               = "46"

}