provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_ip_proto_filter" "thunder_ddos_zone_template_ip_proto_filter" {
  name                       = "test"
  other_filter_action        = "drop"
  other_filter_inverse_match = 0
  other_filter_name          = "47"
  other_filter_regex         = "1047"
  other_filter_seq           = 34
  user_tag                   = "95"

}