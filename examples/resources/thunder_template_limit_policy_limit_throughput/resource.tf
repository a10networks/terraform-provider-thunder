provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_limit_policy_limit_throughput" "thunder_template_limit_policy_limit_throughput" {

  policy_number = 2
  uplink        = 22
  downlink      = 112
  total         = 222
}