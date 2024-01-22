provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ldap_server_host_ipv4" "thunder_ldap_server_host_ipv4" {
  cn_value  = "21"
  dn_value  = "24"
  ipv4_addr = "10.10.10.10"
}