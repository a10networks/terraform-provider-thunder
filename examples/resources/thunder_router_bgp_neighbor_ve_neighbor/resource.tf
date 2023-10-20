provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_neighbor_ve_neighbor" "VEBGP" {
  ve              = 300
  unnumbered      = 1
  as_number       = "300"
  peer_group_name = "A11"
}