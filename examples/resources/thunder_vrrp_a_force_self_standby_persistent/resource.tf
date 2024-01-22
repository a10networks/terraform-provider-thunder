provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_force_self_standby_persistent" "thunder_vrrp_a_force_self_standby_persistent" {

  vrid = 3
}