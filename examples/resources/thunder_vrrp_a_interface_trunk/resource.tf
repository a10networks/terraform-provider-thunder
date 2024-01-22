provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_interface_trunk" "thunder_vrrp_a_interface_trunk" {

  no_heartbeat = 1
  trunk_val    = 3
  user_tag     = "test_vrrp_a"
}