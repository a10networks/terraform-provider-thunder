provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_ldap_instance" "thunder_aam_authentication_server_ldap_instance" {
  admin_dn       = "63"
  admin_secret   = 1
  auth_type      = "ad"
  base           = "83"
  bind_with_dn   = 1
  default_domain = "test"
  derive_bind_dn {
    username_attr = "test"
  }
  dn_attribute = "cn"
  host {
    hostip = "10.10.10.10"
  }
  name = "test"
  sampling_enable {
    counters1 = "all"
  }
  secret_string = "test"
  timeout       = 10
}