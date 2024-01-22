provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v6_type" "thunder_ddos_zone_template_icmp_v6_type" {
  icmp_tmpl_name   = "testing"
  icmp_type_action = "drop"
  type_number      = 192
  user_tag         = "80"

}