provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_filtering_policy" "thunder_template_gtp_filtering_policy" {
  name = "test"
  rat_type_filtering {
    rat_type_value       = "wlan"
    rat_type_filt_action = "monitor"
  }
  user_tag = "77"
}