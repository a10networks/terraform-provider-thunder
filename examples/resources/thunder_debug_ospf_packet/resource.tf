provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_packet" "thunder_debug_ospf_packet" {
  dd         = 0
  detail     = 0
  hello      = 0
  ls_ack     = 0
  ls_request = 0
  ls_update  = 0
  recv       = 0
  send       = 0
}