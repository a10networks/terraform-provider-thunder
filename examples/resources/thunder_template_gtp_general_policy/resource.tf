provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_general_policy" "thunder_template_gtp_general_policy" {
  name                   = "test"
  handover_timeout       = 2
  max_mesg_length_action = "drop"
  maximum_message_length = 1500
  tunnel_timeout         = 1440
  user_tag               = "100"
  v0_action              = "drop"
}