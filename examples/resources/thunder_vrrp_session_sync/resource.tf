provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_vrrp_session_sync" "sync" {
  action = "disable"
}