provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_category_reputation_scope" "thunderWebCategoryReputationScopeTest" {
  name = "testingWebCategory"
  greater_than {
    greater_trustworthy   = 1
    greater_low_risk      = 0
    greater_moderate_risk = 0
    greater_suspicious    = 0
    greater_malicious     = 0
    greater_threshold     = 0
  }
  less_than {
    less_trustworthy   = 0
    less_low_risk      = 1
    less_moderate_risk = 0
    less_suspicious    = 0
    less_malicious     = 0
    less_threshold     = 0
  }
  uuid     = "null"
  user_tag = "a"
  sampling_enable {
    counters1 = "trustworthy"
  }
}