provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_set_rxtx_queue" "thunder_system_set_rxtx_queue" {
  port_index = 17
  rxq_size   = 226
  txq_size   = 788
}