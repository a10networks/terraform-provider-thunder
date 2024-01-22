provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_limit_policy_limit_pps" "thunder_template_limit_policy_limit_pps" {

  policy_number = 2

  ddos_protection_factor = 32
  downlink               = 2132013510
  downlink_relaxed       = 1
  total                  = 1327262937
  total_relaxed          = 1
  uplink                 = 597387216
  uplink_relaxed         = 1
}