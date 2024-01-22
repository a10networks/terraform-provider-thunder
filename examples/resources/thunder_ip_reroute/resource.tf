provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_reroute" "thunder_ip_reroute" {

  suppress_protocols {
    ospf      = 0
    ebgp      = 0
    ibgp      = 0
    static    = 0
    isis      = 0
    rip       = 0
    connected = 0
  }
}