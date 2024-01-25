provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_rules_by_zone_oper" "thunder_rule_set_rules_by_zone_oper" {
  oper {
    group_list {
      rule_list {
        source_list {
        }
        dest_list {
        }
        service_list {
        }
        dscp_list {
        }
      }
    }
  }

}
output "get_rule_set_rules_by_zone_oper" {
  value = ["${data.thunder_rule_set_rules_by_zone_oper.thunder_rule_set_rules_by_zone_oper}"]
}