provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_protocol_limit" "thunder_gslb_protocol_limit" {
  ardt_query    = 220
  ardt_response = 1011
  ardt_session  = 32770
  conn_response = 10
  message       = 11000
  response      = 3610
}