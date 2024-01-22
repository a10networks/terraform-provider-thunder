provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v4" "thunder_ddos_template_icmp_v4" {
  icmp_tmpl_name = "20"
  type_list {
    type_number = 196
    type_rate   = 4400054
    code {
      code_number = 122
      code_rate   = 14668888
    }
    code_other {
      code_other_rate = 15963103
    }
    user_tag = "25"
  }
  type_other {
    type_other_deny = 1
  }
  user_tag = "71"
}