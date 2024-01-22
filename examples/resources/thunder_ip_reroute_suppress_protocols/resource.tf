provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_reroute_suppress_protocols" "thunder_ip_reroute_suppress_protocols" {
  connected = 0
  ebgp      = 0
  ibgp      = 0
  isis      = 0
  ospf      = 0
  rip       = 0
  static    = 0
}