provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_policy" "thunder_template_gtp_policy" {

  name                   = "test"
  filtering_policy_name  = "test_filtering"
  general_policy_name    = "test_general"
  logging_policy_name    = "test_log"
  rate_limit_policy_name = "testrate"
  sampling_enable {
    counters1 = "all"
    counters2 = "drop-flt-message-filtering"
  }
  user_tag = "40"
}