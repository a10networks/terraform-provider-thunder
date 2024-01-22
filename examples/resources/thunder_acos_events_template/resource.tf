provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_template" "thunder_acos_events_template" {

  message_selector_list {
    name     = "test"
    user_tag = "62"
  }
  name     = "test"
  user_tag = "25"
}