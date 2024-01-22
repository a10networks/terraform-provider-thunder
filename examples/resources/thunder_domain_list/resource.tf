provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_domain_list" "thunder_domain_list" {
  name = "testing"
  match_type_axfr {
    axfr_domain        = "test_zone"
    axfr_ip_address    = "10.10.10.10"
    axfr_ipv6_address  = "2001:db8:3333:4444:5555:6666:7777:8888"
    ip_axfr_port_num   = 20
    ipv6_axfr_port_num = 22
    ip_refresh_intvl   = 32
    ipv6_refresh_intvl = 34
  }
  match_type_equals {
    equals = "test"
  }
  match_type_suffix {
    suffix = "testing"
  }
  user_tag = "test"
}