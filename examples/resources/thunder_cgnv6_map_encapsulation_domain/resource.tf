provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_encapsulation_domain" "thunder_cgnv6_map_encapsulation_domain" {
  name = "test"
  basic_mapping_rule {
    rule_ipv4_address_port_settings = "prefix-addr"
    ea_length                       = 14
  }
  description = "201"
  format      = "draft-03"
  health_check_gateway {
    address_list {
      ipv4_gateway = "10.10.10.10"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
  tunnel_endpoint_address = "2001:db8:3333:4444:5555:6666:7777:8888"
  user_tag                = "71"
}