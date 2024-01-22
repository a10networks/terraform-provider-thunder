provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_session_sync" "thunder_vrrp_a_session_sync" {
  action = "disable"
}