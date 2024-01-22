provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_violation_actions" "thunder_ddos_violation_actions" {
  name                   = "test"
  blackhole              = 23
  blacklist_src          = 30
  send_notification_only = 1
  user_tag               = "34"

}