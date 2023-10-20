provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_reroute" "reRoute" {
  suppress_protocols {
    isis      = 1
    ebgp      = 1
    rip       = 1
    connected = 1
    static    = 1
    ospf      = 1
    ibgp      = 1
  }
}