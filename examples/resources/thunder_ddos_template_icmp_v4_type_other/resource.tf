provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v4_type_other" "thunder_ddos_template_icmp_v4_type_other" {
  icmp_tmpl_name  = "20"
  type_other_deny = 1

}