provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v6_filter" "thunder_ddos_zone_template_icmp_v6_filter" {
  icmp_tmpl_name            = "testing"
  icmp_filter_action        = "drop"
  icmp_filter_inverse_match = 1
  icmp_filter_name          = "34"
  icmp_filter_regex         = "529"
  icmp_filter_seq           = 27
  user_tag                  = "28"

}