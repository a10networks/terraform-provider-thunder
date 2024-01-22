provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v6" "thunder_ddos_template_icmp_v6" {
  icmp_tmpl_name = "test"
  type_list {
    type_number = 112
    type_rate   = 195164
    code {
      code_number = 216
      code_rate   = 3573500
    }
    code_other {
      code_other_rate = 8046606
    }
    user_tag = "125"
  }
  type_other {
    type_other_deny = 1
  }
  user_tag = "86"
}