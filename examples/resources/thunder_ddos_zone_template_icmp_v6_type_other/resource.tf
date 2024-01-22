provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_icmp_v6_type_other" "thunder_ddos_zone_template_icmp_v6_type_other" {
  icmp_tmpl_name = "testing"

  icmp_type_other_action = "ignore"

}