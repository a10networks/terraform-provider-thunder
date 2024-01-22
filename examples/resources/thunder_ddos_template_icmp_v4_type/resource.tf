provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v4_type" "thunder_ddos_template_icmp_v4_type" {
  icmp_tmpl_name = "20"
  code {
    code_number = 167
    code_rate   = 1248190
  }
  code_other {
    code_other_rate = 14325190
  }
  type_number = 230
  type_rate   = 14374233
  user_tag    = "66"
}