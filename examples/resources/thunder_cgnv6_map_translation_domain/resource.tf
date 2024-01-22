provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_domain" "thunder_cgnv6_map_translation_domain" {
  name        = "test"
  description = "156"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "84"
}