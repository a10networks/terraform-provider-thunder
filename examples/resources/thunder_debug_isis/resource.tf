provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_isis" "thunder_debug_isis" {
  all    = 0
  bfd    = 0
  events = 0
  ifsm   = 0
  lsp    = 0
  nfsm   = 0
  nsm    = 0
  pdu    = 0
  spf    = 0
}