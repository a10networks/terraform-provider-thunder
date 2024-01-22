provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_message_selector" "thunder_acos_events_message_selector" {
  name     = "test"
  user_tag = "98"
}