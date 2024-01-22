provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_aaa_policy_aaa_rule" "thunder_aam_aaa_policy_aaa_rule" {

  index               = 123
  action              = "allow"
  auth_failure_bypass = 1
  name                = "test"
  domain_name         = "test"
  host {
    host_match_type = "contains"
    host_str        = "59"
  }
  sampling_enable {
    counters1 = "all"
  }
  uri {
    match_type = "contains"
    uri_str    = "61"
  }
  user_agent {
    user_agent_match_type = "contains"
    user_agent_str        = "439"
  }
  user_tag = "127"
}