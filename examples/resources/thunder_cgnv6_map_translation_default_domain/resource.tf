provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_default_domain" "thunder_cgnv6_map_translation_default_domain" {

  name = "test"
}