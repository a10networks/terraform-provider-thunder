provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_account" "thunder_aam_authentication_account" {

  kerberos_spn_list {
    name                   = "41"
    realm                  = "36"
    account                = "4"
    service_principal_name = "34"
    password               = 0
    user_tag               = "95"
  }
  sampling_enable {
    counters1 = "all"
  }

}