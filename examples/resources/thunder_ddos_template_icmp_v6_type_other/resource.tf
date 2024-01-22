provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v6_type_other" "thunder_ddos_template_icmp_v6_type_other" {
  icmp_tmpl_name  = "test"
  type_other_deny = 1

}