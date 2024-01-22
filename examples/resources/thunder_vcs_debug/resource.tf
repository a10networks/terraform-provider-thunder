provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_debug" "thunder_vcs_debug" {
  daemon       = 0
  daemon_msg   = 0
  election     = 0
  election_pdu = 0
  encoder      = 0
  handler      = 0
  info         = 0
  lib          = 0
  net          = 0
  ssl          = 0
  util         = 0
  vblade       = 0
  vblade_msg   = 0
  vmaster      = 0
  vmaster_msg  = 0
}