
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_rule_set"   "RuleSet"  {

    name = "test" 
    session_statistic = "enable" 
    user_tag = "testing" 
    sampling_enable {
        counters1 = "all"
      }
    
    rule_list {
        name = "testin123" 
        action = "permit" 
        idle_timeout = 4 
        dscp_list  {
            dscp_value = "default"
          }
        user_tag = "testing"
      }

  }
