provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_sip" "thunder_debug_sip" {
  ack       = 0
  bye       = 0
  cancel    = 0
  info      = 0
  invite    = 0
  message   = 0
  method    = 0
  notify    = 0
  options   = 0
  prack     = 0
  publish   = 0
  refer     = 0
  register  = 0
  subscribe = 0
  update    = 0
}