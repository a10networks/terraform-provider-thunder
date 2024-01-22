provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_ip_proto" "thunder_ddos_zone_template_ip_proto" {

  filter_list {
    other_filter_name          = "36"
    other_filter_seq           = 78
    other_filter_regex         = "742"
    other_filter_inverse_match = 1
    other_filter_action        = "drop"
    user_tag                   = "118"
  }
  user_tag = "26"
  name     = "test"
}