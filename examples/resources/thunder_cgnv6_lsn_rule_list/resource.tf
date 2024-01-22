provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_rule_list" "thunder_cgnv6_lsn_rule_list" {
  name = "test"
  domain_ip {
    sampling_enable {
      counters1 = "all"
    }
  }
  domain_name_list {
    name_domain = "test"
    user_tag    = "69"
    sampling_enable {
      counters1 = "all"
    }
  }
  http_match_domain_name = 1
  user_tag               = "test"
}