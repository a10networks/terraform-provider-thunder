provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ldap_server_host_ipv6" "test_ldap_server" {
  ipv6_addr = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
  domain    = "test_domain"
  ipv6_addr_cfg {
    port    = 3268
    ssl     = 0
    timeout = 20
  }
  cn_value = "cloud_native"
  group    = "testing"
  base     = "test"
  dn_value = "test_domain"
  domain_cfg {
    port    = 389
    ssl     = 0
    timeout = 10
  }
  port_cfg {
    port    = 389
    ssl     = 0
    timeout = 10
  }
}