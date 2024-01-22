provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_application_oper" "thunder_rule_set_application_oper" {
  oper {
    protocol      = protocol
    rule_set_only = rule_set_only
    rule_list {
      stat_list {
        conns = conns
        bytes = bytes
      }
    }
  }

}
output "get_rule_set_application_oper" {
  value = ["${data.thunder_rule_set_application_oper.thunder_rule_set_application_oper}"]
}