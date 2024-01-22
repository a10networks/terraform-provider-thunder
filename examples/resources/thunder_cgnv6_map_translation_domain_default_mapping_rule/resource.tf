provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_domain_default_mapping_rule" "thunder_cgnv6_map_translation_domain_default_mapping_rule" {

  name             = "test"
  rule_ipv6_prefix = "2001:db8:3333:4444:5555:6666:7777:8888/32"

}