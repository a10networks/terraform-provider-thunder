provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v6_type" "thunder_ddos_template_icmp_v6_type" {
  icmp_tmpl_name = "test"
  code {
    code_number = 188
    code_rate   = 13608811
  }
  code_other {
    code_other_rate = 13697557
  }
  type_number = 165
  type_rate   = 7089273
  user_tag    = "32"

}