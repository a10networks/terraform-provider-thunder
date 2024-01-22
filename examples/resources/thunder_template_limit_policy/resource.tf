provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_limit_policy" "thunder_template_limit_policy" {
  policy_number             = 2
  max_min_fair              = 1
  limit_concurrent_sessions = 20
  log                       = 1
  limit_scope               = "subscriber-prefix"
  prefix_length             = 2
  limit_throughput {
    uplink   = 22
    downlink = 112
    total    = 222
  }
  limit_cps {
    value = 22
  }
}