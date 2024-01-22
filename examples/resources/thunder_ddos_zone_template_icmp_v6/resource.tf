provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v6" "thunder_ddos_zone_template_icmp_v6" {

  icmp_tmpl_name = "testing"
  user_tag       = "test"
  type_list {
    type_number      = 3
    icmp_type_action = "ignore"
  }
  type_other {
    icmp_type_other_action = "ignore"
  }
  filter_list {
    icmp_filter_name          = "test"
    icmp_filter_seq           = 33
    icmp_filter_regex         = "test"
    icmp_filter_inverse_match = 1
    icmp_filter_action        = "ignore"
    user_tag                  = "test"
  }
}