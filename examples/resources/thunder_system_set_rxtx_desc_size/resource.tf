provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_set_rxtx_desc_size" "thunder_system_set_rxtx_desc_size" {
  port_index = 18
  rxd_size   = 52094
  txd_size   = 61332
}