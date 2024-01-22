provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_active_rdt" "thunder_gslb_policy_active_rdt" {
  controller       = 0
  difference       = 0
  enable           = 0
  fail_break       = 0
  ignore_id        = 1
  keep_tracking    = 0
  limit            = 16383
  proto_rdt_enable = 0
  samples          = 5
  single_shot      = 0
  skip             = 3
  timeout          = 3
  tolerance        = 10
}