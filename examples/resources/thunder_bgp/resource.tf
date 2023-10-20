provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_bgp" "BGP" {
  extended_asn_cap = 1
  nexthop_trigger {
    enable = 1
    delay  = 10
  }
}

