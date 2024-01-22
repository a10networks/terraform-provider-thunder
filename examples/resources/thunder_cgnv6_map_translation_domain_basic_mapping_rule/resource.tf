provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_domain_basic_mapping_rule" "thunder_cgnv6_map_translation_domain_basic_mapping_rule" {

  name      = "test"
  ea_length = 3

  rule_ipv4_address_port_settings = "prefix-addr"

}