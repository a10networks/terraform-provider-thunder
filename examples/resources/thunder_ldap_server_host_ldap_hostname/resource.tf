provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ldap_server_host_ldap_hostname" "thunder_ldap_server_host_ldap_hostname" {
  hostname = "a11"
  cn_value = "5"
  dn_value = "97"
  port_cfg {
    port    = 7221
    ssl     = 0
    timeout = 38
  }
}