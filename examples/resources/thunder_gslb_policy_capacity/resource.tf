provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_capacity" "thunder_gslb_policy_capacity" {
  capacity_enable     = 0
  capacity_fail_break = 0
  threshold           = 90
}