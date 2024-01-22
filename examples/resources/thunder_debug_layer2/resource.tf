provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_layer2" "thunder_debug_layer2" {
  all       = 0
  arp       = 0
  interface = 0
  misc      = 0
  trace     = 0
  vlan      = 0
}