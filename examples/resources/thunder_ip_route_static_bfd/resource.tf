provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_route_static_bfd" "ipStaticBFD" {
  local_ip   = "20.20.20.21"
  nexthop_ip = "21.21.21.22"
}