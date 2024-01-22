provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v4" "thunder_ddos_zone_template_icmp_v4" {
  icmp_tmpl_name = "testing"
  user_tag       = "test"
  type_list {
    type_number      = 2
    icmp_type_action = "ignore"
    user_tag         = "test"
  }
  type_other {
    icmp_type_other_action = "ignore"
  }
  filter_list {
    icmp_filter_name          = "test"
    icmp_filter_seq           = 3
    icmp_filter_regex         = "test"
    icmp_filter_inverse_match = 1
    icmp_filter_action        = "ignore"
    user_tag                  = "test"
  }
}