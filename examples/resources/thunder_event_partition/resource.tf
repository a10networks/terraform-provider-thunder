provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_event_partition" "thunder_event_partition" {
  email      = "on"
  logging    = "on"
  user_tag   = "87"
  vnp_events = "part-create"
}