provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v4_type" "thunder_ddos_zone_template_icmp_v4_type" {
  icmp_tmpl_name   = "testing"
  icmp_type_action = "drop"
  type_number      = 97
  user_tag         = "60"

}