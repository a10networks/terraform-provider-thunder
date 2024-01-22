provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_bgp" "thunder_debug_bgp" {
  all        = 0
  bfd        = 0
  dampening  = 0
  events     = 0
  filters    = 0
  fsm        = 0
  in         = 0
  keepalives = 0
  nht        = 0
  nsm        = 0
  out        = 0
  updates    = 0
}