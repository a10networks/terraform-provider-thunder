provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_force_self_standby" "thunder_vrrp_a_force_self_standby" {
  action         = "enable"
  all_partitions = 0
  skip_check     = 0
  vrid           = 0
}